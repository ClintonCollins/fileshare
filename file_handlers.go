package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/gabriel-vasile/mimetype"
	"github.com/go-chi/chi/v5"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"fileshare/database/models"
)

func (inst *httpInstance) postHandleFileUpload(w http.ResponseWriter, r *http.Request) {
	currentAccount := getAccountFromContext(r)
	if currentAccount == nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		inst.logger.Error().Err(err).Msg("Failed to parse multipart form.")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		inst.logger.Error().Err(err).Msg("Failed to get file from form.")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer func() {
		errClose := file.Close()
		if errClose != nil {
			inst.logger.Error().Err(errClose).Msg("Failed to close file.")
		}
	}()
	fileBuf := &bytes.Buffer{}
	teeReader := io.TeeReader(file, fileBuf)
	mType, errMimeType := mimetype.DetectReader(teeReader)
	if errMimeType != nil {
		inst.logger.Error().Err(errMimeType).Msg("Failed to detect mimetype.")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	_, errDrain := io.Copy(io.Discard, teeReader)
	if errDrain != nil {
		inst.logger.Error().Err(errDrain).Msg("Failed to drain tee reader.")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	shortID, errShortID := inst.shortIDGenerator.Generate()
	if errShortID != nil {
		inst.logger.Error().Err(errShortID).Msg("Failed to generate short ID.")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	newFile := models.File{
		AccountID:         currentAccount.ID,
		Name:              fileHeader.Filename,
		Slug:              shortID,
		ShortID:           shortID,
		Size:              fileHeader.Size,
		MimeType:          mType.String(),
		PasswordProtected: false,
		FileGroupID:       null.NewString("", false),
	}

	errInsert := newFile.Insert(r.Context(), inst.db, boil.Infer())
	if errInsert != nil {
		inst.logger.Error().Err(errInsert).Msg("Failed to insert file into database.")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	errMkdir := os.MkdirAll(path.Join(inst.fileStoragePath, currentAccount.ID), 0755)
	if errMkdir != nil {
		inst.logger.Error().Err(errMkdir).Msg("Failed to create directory.")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	storedFile, errFileCreate := os.Create(path.Join(inst.fileStoragePath, currentAccount.ID, newFile.ID))
	if errFileCreate != nil {
		inst.logger.Error().Err(errFileCreate).Msg("Failed to create file.")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer func() {
		errClose := storedFile.Close()
		if errClose != nil {
			inst.logger.Error().Err(errClose).Msg("Failed to close file.")
		}
	}()

	_, errCopy := io.Copy(storedFile, fileBuf)
	if errCopy != nil {
		inst.logger.Error().Err(errCopy).Msg("Failed to copy file.")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	errEncode := json.NewEncoder(w).Encode(&FileUploadResponse{
		Error:   false,
		Message: "File uploaded successfully.",
		URL:     inst.getShortIDURLForFile(newFile.Slug),
	})
	if errEncode != nil {
		inst.logger.Error().Err(errEncode).Msg("Failed to encode JSON response.")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (inst *httpInstance) getFileDownload(w http.ResponseWriter, r *http.Request) {
	shortID := chi.URLParam(r, "shortID")
	if shortID == "" {
		inst.get404(w, r)
		return
	}
	file, err := models.Files(
		models.FileWhere.Slug.EQ(shortID),
		qm.Load(models.FileRels.FileGroup),
		qm.Load(models.FileRels.FilePasswords),
	).One(r.Context(), inst.db)
	if err != nil {
		inst.logger.Error().Err(err).Msg("Failed to get file from database.")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	if file == nil {
		inst.get404(w, r)
		return
	}
	storedFile, errFileOpen := os.Open(path.Join(inst.fileStoragePath, file.AccountID, file.ID))
	if errFileOpen != nil {
		inst.logger.Error().Err(errFileOpen).Msg("Failed to open file.")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	defer func() {
		errClose := storedFile.Close()
		if errClose != nil {
			inst.logger.Error().Err(errClose).Msg("Failed to close file.")
		}
	}()
	storedFileInfo, errFileInfo := storedFile.Stat()
	if errFileInfo != nil {
		inst.logger.Error().Err(errFileInfo).Msg("Failed to get file info.")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	// Make the file download, instead of displaying in browser.
	// w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", file.Name))
	// Make the file show in browser.
	w.Header().Set("Content-Disposition", fmt.Sprintf("inline; filename=%s", file.Name))
	w.Header().Set("Content-Type", file.MimeType)
	w.Header().Set("Content-Length", strconv.FormatInt(storedFileInfo.Size(), 10))
	_, errCopy := io.Copy(w, storedFile)
	if errCopy != nil {
		inst.logger.Error().Err(errCopy).Msg("Failed to copy file.")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
}

func (inst *httpInstance) getFileList(w http.ResponseWriter, r *http.Request) {
	currentAccount := getAccountFromContext(r)
	data := map[string]interface{}{
		"currentAccount": currentAccount,
		"title":          "FileShare File List",
	}

	queryModsForCount := []qm.QueryMod{
		models.FileWhere.AccountID.EQ(currentAccount.ID),
	}
	queryMods := []qm.QueryMod{
		models.FileWhere.AccountID.EQ(currentAccount.ID),
		qm.OrderBy("file.created_at desc"),
		qm.GroupBy("file.id"),
	}

	queryM, queryMC, cParams := commonListQuery(r, 25, "name")
	queryMods = append(queryMods, queryM...)
	queryModsForCount = append(queryModsForCount, queryMC...)

	files, err := models.Files(queryMods...).All(r.Context(), inst.db)
	if err != nil {
		inst.logger.Error().Err(err).Msg("Failed to get files from database.")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	count, err := models.Files(queryModsForCount...).Count(r.Context(), inst.db)
	if err != nil {
		inst.logger.Error().Err(err).Msg("Failed to get files from database.")
		inst.showErrorPage(w, r, http.StatusInternalServerError, "Failed to get files from database.")
		return
	}

	var userFiles []FileTemplateObject
	for _, f := range files {
		isImage := false
		isVideo := false
		isAudio := false
		if strings.HasPrefix(f.MimeType, "image/") {
			isImage = true
		}
		if strings.HasPrefix(f.MimeType, "video/") {
			isVideo = true
		}
		if strings.HasPrefix(f.MimeType, "audio/") {
			isAudio = true
		}
		userFiles = append(userFiles, FileTemplateObject{
			ID:                f.ID,
			Name:              f.Name,
			Slug:              f.Slug,
			ShortID:           f.ShortID,
			Size:              f.Size,
			IsImage:           isImage,
			IsVideo:           isVideo,
			IsAudio:           isAudio,
			MimeType:          f.MimeType,
			URL:               inst.getShortIDURLForFile(f.ShortID),
			PasswordProtected: f.PasswordProtected,
			FileGroupID:       f.FileGroupID,
			CreatedAt:         f.CreatedAt,
		})
	}

	params := url.Values{}
	search := cParams.Search
	if search != "" {
		params.Set("search", search)
	}
	params.Set("limit", strconv.Itoa(cParams.Limit))

	paginationData := genPaginationData(cParams, count, len(userFiles), "/files", params)

	data["files"] = userFiles
	data["search"] = search
	data["paginationData"] = paginationData

	flashMessage, flashErr := inst.getFlashMessage(w, r)
	if flashErr == nil {
		data["FlashMessage"] = flashMessage
	}

	inst.renderTemplate(w, data, "templates/files.gohtml", "templates/layouts/main.gohtml")
}

func (inst *httpInstance) postHandleFileDelete(w http.ResponseWriter, r *http.Request) {
	currentAccount := getAccountFromContext(r)
	err := r.ParseForm()
	if err != nil {
		inst.logger.Error().Err(err).Msg("Failed to parse form.")
		inst.showErrorPage(w, r, http.StatusInternalServerError, "Failed to delete files.")
		return
	}
	fileIDsFormValue, exist := r.PostForm["ids"]
	if !exist || len(fileIDsFormValue) == 0 {
		http.Redirect(w, r, "/files", http.StatusSeeOther)
		return
	}
	var fileIDs []interface{}
	for _, id := range fileIDsFormValue {
		fileIDs = append(fileIDs, id)
	}
	modelFiles, err := models.Files(
		qm.WhereIn("id IN ?", fileIDs...),
		qm.And("account_id = ?", currentAccount.ID),
	).All(r.Context(), inst.db)
	if err != nil {
		inst.logger.Error().Err(err).Msg("Failed to get files from database.")
		inst.showErrorPage(w, r, http.StatusInternalServerError, "Failed to delete files.")
		return
	}
	for _, f := range modelFiles {
		errRemove := os.Remove(path.Join(inst.fileStoragePath, f.AccountID, f.ID))
		if errRemove != nil {
			inst.logger.Error().Err(errRemove).Interface("file", f).Msg("Failed to remove file.")
		}
	}

	filesDeleted, errDeleteAll := modelFiles.DeleteAll(r.Context(), inst.db)
	if errDeleteAll != nil {
		inst.logger.Error().Err(errDeleteAll).Msg("Failed to delete files from database.")
		inst.showErrorPage(w, r, http.StatusInternalServerError, "Failed to delete files.")
		return
	}

	plural := "files"
	if filesDeleted == 1 {
		plural = "file"
	}
	_ = inst.setFlashMessage(w, SuccessFlashAlert, fmt.Sprintf("Successfully deleted %d %s!", filesDeleted, plural))

	http.Redirect(w, r, "/files", http.StatusSeeOther)
}
