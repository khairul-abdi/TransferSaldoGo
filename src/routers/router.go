package routers

import (
	"net/http"

	"github.com/go-chi/chi"
)

func (routes *route) RouterGroup() http.Handler {
	router := chi.NewRouter()
	router.Group(func(r chi.Router) {
		r.Get("/account/{account_number}", routes.ctrl.GetBalance)
		r.Post("/account/{from_account_number}/transfer", routes.ctrl.TransferBalance)
	})

	return router
}
