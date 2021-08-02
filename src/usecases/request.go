package usecases

import (
	"TransferSaldo/src/models"
	"encoding/json"
	"net/http"
	"reflect"
	"strconv"
)

func (u *uc) ProcessTransferRequest(r *http.Request) (*models.TransferBalanceModel, string, int, error) {
	var req models.TransferBalanceModel

	receiver := r.FormValue("to_account_number")
	amount := r.FormValue("amount")

	if reflect.ValueOf(receiver).IsZero() {
		err := json.Unmarshal([]byte(receiver), &req)
		return nil, "Cannot Fetch Receiver ID", http.StatusBadRequest, err
	}

	if reflect.ValueOf(amount).IsZero() {
		err := json.Unmarshal([]byte(amount), &req)
		return nil, "Cannot Fetch Transfer Amount", http.StatusBadRequest, err
	}

	convert_amount, err := strconv.Atoi(amount)
	if err != nil {
		return nil, "Failed Convert Amount to Int", http.StatusInternalServerError, err
	}

	req.Amount = convert_amount

	conver_receiver, err := strconv.Atoi(receiver)
	if err != nil {
		return nil, "Failed Convert Receiver ID to Int", http.StatusInternalServerError, err
	}

	req.ToAccountNumber = conver_receiver

	return &req, "Success", http.StatusOK, nil
}
