package main

import (
	"context"
	"database/sql"
	"fmt"
	"math"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/rs/zerolog"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"fileshare/database/models"
)

type CommonSearchParams struct {
	Search string
	Page   int
	Limit  int
	Offset int
	Sort   string
	Filter string
}

func genPaginationData(cParams CommonSearchParams, count int64, itemLength int, urlPath string,
	urlParams url.Values) PaginationData {
	return PaginationData{
		CurrentPage:     cParams.Page,
		TotalPages:      int(math.Ceil(float64(count) / float64(cParams.Limit))),
		TotalItems:      count,
		ItemsPerPage:    cParams.Limit,
		HasNextPage:     cParams.Page < int(math.Ceil(float64(count)/float64(cParams.Limit))),
		NextPage:        cParams.Page + 1,
		HasPreviousPage: cParams.Page > 1,
		PreviousPage:    cParams.Page - 1,
		PageRangeStart:  cParams.Offset + 1,
		PageRangeEnd:    cParams.Offset + itemLength,
		URL:             urlPath,
		URLParams:       urlParams.Encode(),
	}
}

func commonListQuery(r *http.Request, limit int, searchColumnNames ...string) ([]qm.QueryMod, []qm.QueryMod,
	CommonSearchParams) {
	page := 1
	if r.FormValue("page") != "" {
		p, err := strconv.Atoi(r.FormValue("page"))
		if err == nil {
			if p > 0 {
				page = p
			}
		}
	}

	if r.FormValue("limit") != "" {
		l, err := strconv.Atoi(r.FormValue("limit"))
		if err == nil {
			if l > 0 {
				limit = l
			}
		}
	}

	offset := (page - 1) * limit

	var queryMods []qm.QueryMod
	var queryModsForCount []qm.QueryMod

	search := r.FormValue("search")
	queryMods = append(queryMods, qm.Limit(limit), qm.Offset(offset))
	if search != "" {
		var searchQM []qm.QueryMod
		var searchQMC []qm.QueryMod
		for i, s := range searchColumnNames {
			searchString := fmt.Sprintf("lower(%s) like ?", s)
			if i == 0 {
				searchQM = append(searchQM, qm.Where(searchString, "%"+search+"%"))
				searchQMC = append(searchQMC, qm.Where(searchString, "%"+search+"%"))
			} else {
				searchQM = append(searchQM, qm.Or(searchString, "%"+search+"%"))
				searchQMC = append(searchQMC, qm.Or(searchString, "%"+search+"%"))
			}
		}
		if len(searchQM) > 0 {
			queryMods = append(queryMods, qm.Expr(searchQM...))
		}
		if len(searchQMC) > 0 {
			queryModsForCount = append(queryModsForCount, qm.Expr(searchQMC...))
		}
	}

	commonParams := CommonSearchParams{
		Search: search,
		Page:   page,
		Limit:  limit,
		Offset: offset,
		Sort:   r.FormValue("sort"),
		Filter: r.FormValue("filter"),
	}
	return queryMods, queryModsForCount, commonParams
}

func cleanUpOldSession(db *sql.DB, logger zerolog.Logger) {
	_, err := models.Sessions(qm.Where("expires_at < now()")).DeleteAll(context.Background(), db)
	if err != nil {
		logger.Error().Err(err).Msg("Error deleting old sessions")
	}
}

func automatedProcessLooper(ctx context.Context, db *sql.DB, logger zerolog.Logger) {
	throttle := time.NewTicker(time.Minute * 1)
	for {
		select {
		case <-ctx.Done():
			return
		case <-throttle.C:
			cleanUpOldSession(db, logger)
		}
	}
}
