package info

import (
	"net/http"
	"os"
)

type Handler struct {
}

func (h *Handler) HealthCheck() int {
	resp, err := http.Get("http://127.0.0.1:8080/info")
	if err != nil {
		panic(err)
	}
	return resp.StatusCode
}

func (h *Handler) Version() string {

	AppV := os.Getenv("APP_VERSION")
	return AppV
}
