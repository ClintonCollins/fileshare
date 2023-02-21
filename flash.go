package main

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"net/http"
)

type FlashAlert string

const (
	SuccessFlashAlert FlashAlert = "success"
	InfoFlashAlert    FlashAlert = "info"
	WarningFlashAlert FlashAlert = "warning"
	ErrorFlashAlert   FlashAlert = "error"

	FlashCookieName = "flash"
)

type FlashMessage struct {
	Alert   FlashAlert
	Message string
}

func (inst *httpInstance) setFlashMessage(w http.ResponseWriter, alert FlashAlert, message string) error {
	msg := FlashMessage{
		Alert:   alert,
		Message: message,
	}

	var b bytes.Buffer
	errEncode := gob.NewEncoder(&b).Encode(msg)
	if errEncode != nil {
		inst.logger.Error().Err(errEncode).Msg("Failed to encode flash message")
		return errEncode
	}
	sessionCookie := &http.Cookie{
		Name:     FlashCookieName,
		Value:    base64.StdEncoding.EncodeToString(b.Bytes()),
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   inst.useSSL || inst.useAutoSSL,
	}
	http.SetCookie(w, sessionCookie)
	return nil
}

func (inst *httpInstance) getFlashMessage(w http.ResponseWriter, r *http.Request) (*FlashMessage, error) {
	cookie, err := r.Cookie(FlashCookieName)
	if err != nil {
		return nil, err
	}
	cookieValue, errDecode := base64.StdEncoding.DecodeString(cookie.Value)
	if errDecode != nil {
		inst.logger.Error().Err(errDecode).Msg("Failed to decode flash message")
		return nil, errDecode
	}
	var msg FlashMessage
	errGob := gob.NewDecoder(bytes.NewReader(cookieValue)).Decode(&msg)
	if errGob != nil {
		inst.logger.Error().Err(errGob).Msg("Failed to decode flash message")
		return nil, errGob
	}

	// Clear the cookie
	sessionCookie := &http.Cookie{
		Name:     FlashCookieName,
		Path:     "/",
		MaxAge:   -1,
		SameSite: http.SameSiteLaxMode,
	}
	http.SetCookie(w, sessionCookie)
	return &msg, nil
}
