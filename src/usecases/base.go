package usecases

import (
	"TransferSaldo/src/models"
	"TransferSaldo/src/repositories"
	"net/http"
)

type uc struct {
	repo repositories.Repo
}

type UC interface {
	GetAccountID(urlstring string) (int, error)
	GetBalanceInfo(account_number int) (*models.GetBalance, error)
	TransferBalance(sender int, req *models.TransferBalanceModel) (string, int, error)
	ProcessTransferRequest(r *http.Request) (*models.TransferBalanceModel, string, int, error)
}

func NewUC(repo repositories.Repo) UC {
	return &uc{repo: repo}
}
