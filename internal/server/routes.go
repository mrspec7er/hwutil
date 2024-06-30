package server

import (
	"net/http"

	"github.com/mrspec7er/hwutil/internal/app"
)

func routerConfig() {
	var appHandler app.Handler
	http.HandleFunc("/subscribe", appHandler.Subscribe)
}
