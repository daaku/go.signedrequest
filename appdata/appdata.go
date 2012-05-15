package appdata

import (
	"encoding/base64"
	"github.com/nshah/go.signedrequest/fbsr"
	"log"
	"net/http"
	"net/url"
)

// A handler to allow app_data based request transformation.
type Handler struct {
	Handler http.Handler
	Secret  []byte
}

// Unpacks the URL from app_data if possible.
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	rawSr := r.FormValue("signed_request")
	if rawSr != "" {
		sr, err := fbsr.Unmarshal([]byte(rawSr), h.Secret)
		if err != nil {
			log.Printf("Ignoring error in parsing signed request: %s", err)
		} else if sr.AppData != "" {
			u, err := Decode(sr.AppData)
			if err != nil {
				log.Printf("Ignoring error in parsing app data: %s", err)
			} else {
				r.URL.Path = u.Path
				r.URL.RawQuery = u.RawQuery
				r.Method = "get"
			}
		}
	}
	h.Handler.ServeHTTP(w, r)
}

// Decode a URL from app_data
func Decode(appData string) (*url.URL, error) {
	b, err := base64.URLEncoding.DecodeString(appData)
	if err != nil {
		return nil, err
	}
	return url.ParseRequestURI(string(b))
}

// Encodes a URL for app_data
func Encode(u *url.URL) string {
	return base64.URLEncoding.EncodeToString([]byte(u.RequestURI()))
}
