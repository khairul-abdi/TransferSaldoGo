package controllers

import (
	"TransferSaldo/packages"
	"net/http"
)

func (c *ctrl) GetBalance(w http.ResponseWriter, r *http.Request) {
	account_number, err := c.uc.GetAccountID(r.URL.Path)
	if err != nil {
		packages.BasicResponse(w, "Error Fetch Account ID from URL", http.StatusBadRequest, err.Error())
		return
	}

	res, err := c.uc.GetBalanceInfo(account_number)
	if err != nil {
		packages.BasicResponse(w, "Failed to Get Balance", http.StatusNotFound, err.Error())
		return
	}

	packages.BasicResponse(w, "Success", http.StatusOK, res)
}
