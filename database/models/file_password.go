// Code generated by SQLBoiler 4.14.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// FilePassword is an object representing the database table.
type FilePassword struct {
	ID           string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	FileID       string    `boil:"file_id" json:"file_id" toml:"file_id" yaml:"file_id"`
	PasswordHash string    `boil:"password_hash" json:"password_hash" toml:"password_hash" yaml:"password_hash"`
	CreatedAt    time.Time `boil:"created_at" json:"created_at" toml:"created_at" yaml:"created_at"`
	UpdatedAt    time.Time `boil:"updated_at" json:"updated_at" toml:"updated_at" yaml:"updated_at"`

	R *filePasswordR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L filePasswordL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var FilePasswordColumns = struct {
	ID           string
	FileID       string
	PasswordHash string
	CreatedAt    string
	UpdatedAt    string
}{
	ID:           "id",
	FileID:       "file_id",
	PasswordHash: "password_hash",
	CreatedAt:    "created_at",
	UpdatedAt:    "updated_at",
}

var FilePasswordTableColumns = struct {
	ID           string
	FileID       string
	PasswordHash string
	CreatedAt    string
	UpdatedAt    string
}{
	ID:           "file_password.id",
	FileID:       "file_password.file_id",
	PasswordHash: "file_password.password_hash",
	CreatedAt:    "file_password.created_at",
	UpdatedAt:    "file_password.updated_at",
}

// Generated where

var FilePasswordWhere = struct {
	ID           whereHelperstring
	FileID       whereHelperstring
	PasswordHash whereHelperstring
	CreatedAt    whereHelpertime_Time
	UpdatedAt    whereHelpertime_Time
}{
	ID:           whereHelperstring{field: "\"file_password\".\"id\""},
	FileID:       whereHelperstring{field: "\"file_password\".\"file_id\""},
	PasswordHash: whereHelperstring{field: "\"file_password\".\"password_hash\""},
	CreatedAt:    whereHelpertime_Time{field: "\"file_password\".\"created_at\""},
	UpdatedAt:    whereHelpertime_Time{field: "\"file_password\".\"updated_at\""},
}

// FilePasswordRels is where relationship names are stored.
var FilePasswordRels = struct {
	File string
}{
	File: "File",
}

// filePasswordR is where relationships are stored.
type filePasswordR struct {
	File *File `boil:"File" json:"File" toml:"File" yaml:"File"`
}

// NewStruct creates a new relationship struct
func (*filePasswordR) NewStruct() *filePasswordR {
	return &filePasswordR{}
}

func (r *filePasswordR) GetFile() *File {
	if r == nil {
		return nil
	}
	return r.File
}

// filePasswordL is where Load methods for each relationship are stored.
type filePasswordL struct{}

var (
	filePasswordAllColumns            = []string{"id", "file_id", "password_hash", "created_at", "updated_at"}
	filePasswordColumnsWithoutDefault = []string{"file_id", "password_hash"}
	filePasswordColumnsWithDefault    = []string{"id", "created_at", "updated_at"}
	filePasswordPrimaryKeyColumns     = []string{"id"}
	filePasswordGeneratedColumns      = []string{}
)

type (
	// FilePasswordSlice is an alias for a slice of pointers to FilePassword.
	// This should almost always be used instead of []FilePassword.
	FilePasswordSlice []*FilePassword
	// FilePasswordHook is the signature for custom FilePassword hook methods
	FilePasswordHook func(context.Context, boil.ContextExecutor, *FilePassword) error

	filePasswordQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	filePasswordType                 = reflect.TypeOf(&FilePassword{})
	filePasswordMapping              = queries.MakeStructMapping(filePasswordType)
	filePasswordPrimaryKeyMapping, _ = queries.BindMapping(filePasswordType, filePasswordMapping, filePasswordPrimaryKeyColumns)
	filePasswordInsertCacheMut       sync.RWMutex
	filePasswordInsertCache          = make(map[string]insertCache)
	filePasswordUpdateCacheMut       sync.RWMutex
	filePasswordUpdateCache          = make(map[string]updateCache)
	filePasswordUpsertCacheMut       sync.RWMutex
	filePasswordUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var filePasswordAfterSelectHooks []FilePasswordHook

var filePasswordBeforeInsertHooks []FilePasswordHook
var filePasswordAfterInsertHooks []FilePasswordHook

var filePasswordBeforeUpdateHooks []FilePasswordHook
var filePasswordAfterUpdateHooks []FilePasswordHook

var filePasswordBeforeDeleteHooks []FilePasswordHook
var filePasswordAfterDeleteHooks []FilePasswordHook

var filePasswordBeforeUpsertHooks []FilePasswordHook
var filePasswordAfterUpsertHooks []FilePasswordHook

// doAfterSelectHooks executes all "after Select" hooks.
func (o *FilePassword) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range filePasswordAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *FilePassword) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range filePasswordBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *FilePassword) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range filePasswordAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *FilePassword) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range filePasswordBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *FilePassword) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range filePasswordAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *FilePassword) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range filePasswordBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *FilePassword) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range filePasswordAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *FilePassword) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range filePasswordBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *FilePassword) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range filePasswordAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFilePasswordHook registers your hook function for all future operations.
func AddFilePasswordHook(hookPoint boil.HookPoint, filePasswordHook FilePasswordHook) {
	switch hookPoint {
	case boil.AfterSelectHook:
		filePasswordAfterSelectHooks = append(filePasswordAfterSelectHooks, filePasswordHook)
	case boil.BeforeInsertHook:
		filePasswordBeforeInsertHooks = append(filePasswordBeforeInsertHooks, filePasswordHook)
	case boil.AfterInsertHook:
		filePasswordAfterInsertHooks = append(filePasswordAfterInsertHooks, filePasswordHook)
	case boil.BeforeUpdateHook:
		filePasswordBeforeUpdateHooks = append(filePasswordBeforeUpdateHooks, filePasswordHook)
	case boil.AfterUpdateHook:
		filePasswordAfterUpdateHooks = append(filePasswordAfterUpdateHooks, filePasswordHook)
	case boil.BeforeDeleteHook:
		filePasswordBeforeDeleteHooks = append(filePasswordBeforeDeleteHooks, filePasswordHook)
	case boil.AfterDeleteHook:
		filePasswordAfterDeleteHooks = append(filePasswordAfterDeleteHooks, filePasswordHook)
	case boil.BeforeUpsertHook:
		filePasswordBeforeUpsertHooks = append(filePasswordBeforeUpsertHooks, filePasswordHook)
	case boil.AfterUpsertHook:
		filePasswordAfterUpsertHooks = append(filePasswordAfterUpsertHooks, filePasswordHook)
	}
}

// One returns a single filePassword record from the query.
func (q filePasswordQuery) One(ctx context.Context, exec boil.ContextExecutor) (*FilePassword, error) {
	o := &FilePassword{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for file_password")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all FilePassword records from the query.
func (q filePasswordQuery) All(ctx context.Context, exec boil.ContextExecutor) (FilePasswordSlice, error) {
	var o []*FilePassword

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to FilePassword slice")
	}

	if len(filePasswordAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all FilePassword records in the query.
func (q filePasswordQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count file_password rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q filePasswordQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if file_password exists")
	}

	return count > 0, nil
}

// File pointed to by the foreign key.
func (o *FilePassword) File(mods ...qm.QueryMod) fileQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.FileID),
	}

	queryMods = append(queryMods, mods...)

	return Files(queryMods...)
}

// LoadFile allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (filePasswordL) LoadFile(ctx context.Context, e boil.ContextExecutor, singular bool, maybeFilePassword interface{}, mods queries.Applicator) error {
	var slice []*FilePassword
	var object *FilePassword

	if singular {
		var ok bool
		object, ok = maybeFilePassword.(*FilePassword)
		if !ok {
			object = new(FilePassword)
			ok = queries.SetFromEmbeddedStruct(&object, &maybeFilePassword)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", object, maybeFilePassword))
			}
		}
	} else {
		s, ok := maybeFilePassword.(*[]*FilePassword)
		if ok {
			slice = *s
		} else {
			ok = queries.SetFromEmbeddedStruct(&slice, maybeFilePassword)
			if !ok {
				return errors.New(fmt.Sprintf("failed to set %T from embedded struct %T", slice, maybeFilePassword))
			}
		}
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &filePasswordR{}
		}
		args = append(args, object.FileID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &filePasswordR{}
			}

			for _, a := range args {
				if a == obj.FileID {
					continue Outer
				}
			}

			args = append(args, obj.FileID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`file`),
		qm.WhereIn(`file.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load File")
	}

	var resultSlice []*File
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice File")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for file")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for file")
	}

	if len(fileAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.File = foreign
		if foreign.R == nil {
			foreign.R = &fileR{}
		}
		foreign.R.FilePasswords = append(foreign.R.FilePasswords, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.FileID == foreign.ID {
				local.R.File = foreign
				if foreign.R == nil {
					foreign.R = &fileR{}
				}
				foreign.R.FilePasswords = append(foreign.R.FilePasswords, local)
				break
			}
		}
	}

	return nil
}

// SetFile of the filePassword to the related item.
// Sets o.R.File to related.
// Adds o to related.R.FilePasswords.
func (o *FilePassword) SetFile(ctx context.Context, exec boil.ContextExecutor, insert bool, related *File) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"file_password\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"file_id"}),
		strmangle.WhereClause("\"", "\"", 2, filePasswordPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.FileID = related.ID
	if o.R == nil {
		o.R = &filePasswordR{
			File: related,
		}
	} else {
		o.R.File = related
	}

	if related.R == nil {
		related.R = &fileR{
			FilePasswords: FilePasswordSlice{o},
		}
	} else {
		related.R.FilePasswords = append(related.R.FilePasswords, o)
	}

	return nil
}

// FilePasswords retrieves all the records using an executor.
func FilePasswords(mods ...qm.QueryMod) filePasswordQuery {
	mods = append(mods, qm.From("\"file_password\""))
	q := NewQuery(mods...)
	if len(queries.GetSelect(q)) == 0 {
		queries.SetSelect(q, []string{"\"file_password\".*"})
	}

	return filePasswordQuery{q}
}

// FindFilePassword retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFilePassword(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*FilePassword, error) {
	filePasswordObj := &FilePassword{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"file_password\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, filePasswordObj)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from file_password")
	}

	if err = filePasswordObj.doAfterSelectHooks(ctx, exec); err != nil {
		return filePasswordObj, err
	}

	return filePasswordObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *FilePassword) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no file_password provided for insertion")
	}

	var err error
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		if o.UpdatedAt.IsZero() {
			o.UpdatedAt = currTime
		}
	}

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(filePasswordColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	filePasswordInsertCacheMut.RLock()
	cache, cached := filePasswordInsertCache[key]
	filePasswordInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			filePasswordAllColumns,
			filePasswordColumnsWithDefault,
			filePasswordColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(filePasswordType, filePasswordMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(filePasswordType, filePasswordMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"file_password\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"file_password\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into file_password")
	}

	if !cached {
		filePasswordInsertCacheMut.Lock()
		filePasswordInsertCache[key] = cache
		filePasswordInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the FilePassword.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *FilePassword) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		o.UpdatedAt = currTime
	}

	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	filePasswordUpdateCacheMut.RLock()
	cache, cached := filePasswordUpdateCache[key]
	filePasswordUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			filePasswordAllColumns,
			filePasswordPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update file_password, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"file_password\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, filePasswordPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(filePasswordType, filePasswordMapping, append(wl, filePasswordPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update file_password row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for file_password")
	}

	if !cached {
		filePasswordUpdateCacheMut.Lock()
		filePasswordUpdateCache[key] = cache
		filePasswordUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q filePasswordQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for file_password")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for file_password")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FilePasswordSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), filePasswordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"file_password\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, filePasswordPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in filePassword slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all filePassword")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *FilePassword) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no file_password provided for upsert")
	}
	if !boil.TimestampsAreSkipped(ctx) {
		currTime := time.Now().In(boil.GetLocation())

		if o.CreatedAt.IsZero() {
			o.CreatedAt = currTime
		}
		o.UpdatedAt = currTime
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(filePasswordColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	filePasswordUpsertCacheMut.RLock()
	cache, cached := filePasswordUpsertCache[key]
	filePasswordUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			filePasswordAllColumns,
			filePasswordColumnsWithDefault,
			filePasswordColumnsWithoutDefault,
			nzDefaults,
		)

		update := updateColumns.UpdateColumnSet(
			filePasswordAllColumns,
			filePasswordPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert file_password, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(filePasswordPrimaryKeyColumns))
			copy(conflict, filePasswordPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"file_password\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(filePasswordType, filePasswordMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(filePasswordType, filePasswordMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if errors.Is(err, sql.ErrNoRows) {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert file_password")
	}

	if !cached {
		filePasswordUpsertCacheMut.Lock()
		filePasswordUpsertCache[key] = cache
		filePasswordUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single FilePassword record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *FilePassword) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no FilePassword provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), filePasswordPrimaryKeyMapping)
	sql := "DELETE FROM \"file_password\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from file_password")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for file_password")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q filePasswordQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no filePasswordQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from file_password")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for file_password")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FilePasswordSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(filePasswordBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), filePasswordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"file_password\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, filePasswordPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from filePassword slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for file_password")
	}

	if len(filePasswordAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *FilePassword) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindFilePassword(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FilePasswordSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := FilePasswordSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), filePasswordPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"file_password\".* FROM \"file_password\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, filePasswordPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in FilePasswordSlice")
	}

	*o = slice

	return nil
}

// FilePasswordExists checks if the FilePassword row exists.
func FilePasswordExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"file_password\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if file_password exists")
	}

	return exists, nil
}

// Exists checks if the FilePassword row exists.
func (o *FilePassword) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	return FilePasswordExists(ctx, exec, o.ID)
}