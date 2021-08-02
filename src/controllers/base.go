package controllers

import (
	"TransferSaldo/src/usecases"
	"net/http"
)

type ctrl struct {
	uc usecases.UC
}

type Controllers interface {
	GetBalance(w http.ResponseWriter, r *http.Request)
	TransferBalance(w http.ResponseWriter, r *http.Request)
}

func NewCtrl(uc usecases.UC) Controllers {
	return &ctrl{uc: uc}
}
