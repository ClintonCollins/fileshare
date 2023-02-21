package filesharecli

import (
	"database/sql"
	"fmt"
	"net/mail"
	"os"
	"strings"
	"time"

	"github.com/cqroot/prompt"
	"github.com/cqroot/prompt/input"
	"github.com/lib/pq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"

	"fileshare/database/models"
	"fileshare/pkg/configuration"
	"fileshare/pkg/sharelogger"

	"fileshare/database"
)

func Start() {
	app := &cli.App{
		Name:        "fileshare",
		Version:     "0.0.1",
		Description: "Fileshare is a simple file sharing service.",
		Commands: []*cli.Command{
			{
				Name: "invite",
				Subcommands: []*cli.Command{
					{
						Name:   "create",
						Action: inviteCreate,
					},
				},
			},
			{
				Name: "user",
				Subcommands: []*cli.Command{
					{
						Name:   "promote",
						Action: userSuperuser,
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}
}

func setupAll() (*configuration.Configuration, *sql.DB, zerolog.Logger, error) {
	c, errConfig := configuration.Get()
	if errConfig != nil {
		return nil, nil, log.Logger, errConfig
	}
	l := sharelogger.GetLogger(c.DevMode)
	d, errDB := database.GetDB(&database.Config{
		PostgresHost:     c.PostgresHost,
		PostgresPort:     c.PostgresPort,
		PostgresUser:     c.PostgresUser,
		PostgresPassword: c.PostgresPassword,
		PostgresDatabase: c.PostgresDatabase,
	})
	if errDB != nil {
		l.Error().Err(errDB).Msg("")
		return nil, nil, log.Logger, errDB
	}
	return c, d, l, nil
}

func inviteCreate(context *cli.Context) error {
	config, db, logger, err := setupAll()
	if err != nil {
		return err
	}
	email, err := prompt.New().Ask("What is the email of the user you'd like to invite?").Input("",
		input.WithValidateFunc(func(s string) error {
			_, errParse := mail.ParseAddress(s)
			return errParse
		}))
	if err != nil {
		logger.Error().Err(err).Msg("")
		os.Exit(1)
	}
	newInvitation := &models.Invitation{
		Email:     email,
		Active:    true,
		ExpiresAt: time.Now().AddDate(0, 0, 14),
	}
	errInsert := newInvitation.Insert(context.Context, db, boil.Infer())
	if errInsert != nil {
		logger.Error().Err(errInsert).Msg("Unable to create new invitation.")
		return errInsert
	}
	tString := newInvitation.ID
	if config.PublicURL != "" {
		tString = fmt.Sprintf("%s/login?invite_token=%s", config.PublicURL, newInvitation.ID)
	}
	fmt.Printf("Invitation token: %s\nExpires at: %s\n", tString,
		newInvitation.ExpiresAt.Format(time.DateTime))
	return nil
}

func userSuperuser(ctx *cli.Context) error {
	_, db, _, err := setupAll()
	if err != nil {
		return err
	}
	oauthMembers, errAcc := models.AuthProviders().All(ctx.Context, db)
	if errAcc != nil {
		return errAcc
	}
	var choices []string
	for _, o := range oauthMembers {
		choices = append(choices, fmt.Sprintf("(%s) | NAME = %s | EMAIL = %s |", o.AccountID, o.Name, o.Email))
	}
	u, errP := prompt.New().Ask("Select accounts to make Superuser").MultiChoose(choices)
	if errP != nil {
		return errP
	}
	var selectedIds []interface{}
	for _, s := range u {
		selected := strings.Split(s, ")")[0]
		selected = strings.TrimPrefix(selected, "(")
		selectedIds = append(selectedIds, selected)
	}
	_, errExec := queries.Raw(`update account set is_superuser = true where id = any($1)`,
		pq.Array(selectedIds)).ExecContext(ctx.Context, db)
	if errExec != nil {
		return errExec
	}
	fmt.Printf("Successfully made %d accounts superusers.\n", len(selectedIds))
	return nil
}
