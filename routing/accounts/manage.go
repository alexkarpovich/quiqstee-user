package accounts

import (
  "log"
  "net/http"
  "github.com/alexkarpovich/quiqstee-user/lib"
  "github.com/alexkarpovich/quiqstee-user/database/models"
)

func (h *AccountHandler) View(w http.ResponseWriter, r *http.Request) {
  user, _ := r.Context().Value("user").(*models.User)
  log.Printf("LOGGED_IN_USER: %s", user)

  lib.SendJson(w, user, http.StatusOK)
}
