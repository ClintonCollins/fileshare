package main

import (
	"fmt"
	"html/template"
	"io/fs"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/dustin/go-humanize"
)

type PaginationData struct {
	CurrentPage     int
	TotalPages      int
	TotalItems      int64
	ItemsPerPage    int
	HasNextPage     bool
	NextPage        int
	HasPreviousPage bool
	PreviousPage    int
	PageRangeStart  int
	PageRangeEnd    int
	URL             string
	URLParams       string
}

func (inst *httpInstance) getTemplateFuncMap() template.FuncMap {
	return template.FuncMap{
		"humanizeBytes": func(bytes int64) string {
			return humanize.Bytes(uint64(bytes))
		},
		"inviteURL": func(inviteToken string) string {
			return fmt.Sprintf("%s/login?invite_token=%s", inst.publicURL, inviteToken)
		},
		"unescapeURLParams": func(urlParams string) string {
			s, err := url.QueryUnescape(urlParams)
			if err != nil {
				return urlParams
			}
			return s
		},
		"paginate": func(paginationData PaginationData, itemName string) template.HTML {
			html := strings.Builder{}
			params := ""
			if paginationData.URLParams != "" {
				params = "&" + paginationData.URLParams
			}
			if itemName == "" {
				itemName = "items"
			}
			if paginationData.HasPreviousPage {
				html.WriteString(fmt.Sprintf(`<a role=button href="%s?page=%d%s" class="button-sm">Prev</a>`,
					paginationData.URL, paginationData.PreviousPage, params))
			} else {
				html.WriteString(fmt.Sprintf(`<a role=button href="%s?page=%d%s" class="button-sm" disabled>Prev</a>`,
					paginationData.URL, paginationData.PreviousPage, params))
			}
			html.WriteString(fmt.Sprintf(`<span class="center-items">Showing %d - %d out of %s total %s</span>`,
				paginationData.PageRangeStart, paginationData.PageRangeEnd, humanize.Comma(paginationData.TotalItems),
				itemName))
			if paginationData.HasNextPage {
				html.WriteString(fmt.Sprintf(`<a role=button href="%s?page=%d%s" class="button-sm">Next</a>`,
					paginationData.URL, paginationData.NextPage, params))
			} else {
				html.WriteString(fmt.Sprintf(`<a role=button href="%s?page=%d%s" class="button-sm" disabled>Next</a>`,
					paginationData.URL, paginationData.NextPage, params))
			}
			return template.HTML(html.String())
		},
	}
}

func (inst *httpInstance) renderTemplate(w http.ResponseWriter, data interface{}, templatePaths ...string) {
	tmpl, exists := inst.mainTemplates[strings.Join(templatePaths, ",")]
	if !exists {
		inst.logger.Debug().Msgf("Template not found in cache, parsing... %s", strings.Join(templatePaths, ","))
		t, err := inst.parseTemplate(templatePaths...)
		if err != nil {
			inst.logger.Fatal().Err(err).Msg("Failed to parse template.")
			return
		}
		tmpl = t
	}
	if tmpl == nil {
		inst.logger.Error().Msg("Template is nil.")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	errExecute := tmpl.Execute(w, data)
	if errExecute != nil {
		inst.logger.Error().Err(errExecute).Msg("Failed to execute template.")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (inst *httpInstance) parseTemplate(templatePaths ...string) (*template.Template, error) {
	var templatesFS fs.FS = templates
	if inst.devMode {
		templatesFS = os.DirFS(filepath.Join(getCurrentFileDirectory(), "templates"))
		for i, t := range templatePaths {
			templatePaths[i] = strings.TrimPrefix(t, "templates/")
		}
	}
	tmpl, err := template.New(path.Base(templatePaths[0])).Funcs(inst.getTemplateFuncMap()).ParseFS(templatesFS,
		templatePaths...)
	if err != nil {
		return nil, err
	}
	return tmpl, nil
}

func (inst *httpInstance) buildTemplatesMap() (map[string]*template.Template, error) {
	temps := make(map[string]*template.Template)
	mainTemplateFiles, errRead := fs.ReadDir(templates, "templates")
	if errRead != nil {
		return nil, errRead
	}
	adminTemplateFiles, errRead := fs.ReadDir(templates, "templates/admin")
	if errRead != nil {
		return nil, errRead
	}
	for _, t := range mainTemplateFiles {
		if !t.IsDir() {
			tmpl, err := template.New(t.Name()).Funcs(inst.getTemplateFuncMap()).ParseFS(templates,
				"templates/"+t.Name(),
				"templates/layouts/main.gohtml")
			if err != nil {
				return nil, err
			}
			temps["templates/"+t.Name()+","+"templates/layouts/main.gohtml"] = tmpl
		}
	}
	for _, t := range adminTemplateFiles {
		if !t.IsDir() {
			tmpl, err := template.New(t.Name()).Funcs(inst.getTemplateFuncMap()).ParseFS(templates,
				"templates/admin/"+t.Name(),
				"templates/layouts/main.gohtml")
			if err != nil {
				return nil, err
			}
			temps["templates/admin/"+t.Name()+","+"templates/layouts/main.gohtml"] = tmpl
		}
	}
	return temps, nil
}
