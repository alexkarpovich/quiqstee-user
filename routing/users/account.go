package users

import (
  "log"
  "net/http"
  "github.com/alexkarpovich/quiqstee-user/lib"
  "github.com/alexkarpovich/quiqstee-user/database/models"
)

func (h *UserHandler) View(w http.ResponseWriter, r *http.Request) {
  var user models.User
  s, _ := r.Context().Value("user").(*models.User)
  log.Printf("LOGGED_IN_USER: %s", s)

  lib.SendJson(w, user, http.StatusOK)
}
