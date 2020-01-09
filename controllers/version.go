package controllers

import (
	"encoding/json"
	"github.com/SkYNewZ/aloesia/config"
	"net/http"
)

// VersionHandler version endpoint: print bin version and built date
func VersionHandler(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(map[string]string{ //nolint
		"build_time": config.BuildTime,
		"verssion":   config.CommitHash,
	})
}
