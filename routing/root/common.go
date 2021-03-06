package root

import (
    "net/http"
    "github.com/alexkarpovich/quiqstee-user/lib"
)

func (h *RootHandler) HealthCheck(w http.ResponseWriter, r *http.Request) {
    lib.SendJson(w, "Still alive!", http.StatusOK)
}

func (h *RootHandler) Handler404(w http.ResponseWriter, r *http.Request) {
    lib.SendJsonError(w, "404 - Page is not found.", http.StatusNotFound)
}
