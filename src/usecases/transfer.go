package usecases

import (
	"TransferSaldo/src/models"
	"errors"
	"net/http"

	"gopkg.in/mgo.v2/bson"
)

func (u *uc) TransferBalance(sender int, req *models.TransferBalanceModel) (string, int, error) {
	var acc models.AccountModel

	err := u.repo.FindOne(&acc, bson.M{"account_number": sender}, "account")
	if err != nil {
		return "Cannot fetch sender ID from database", http.StatusBadRequest, err
	}

	if acc.Balance < req.Amount {
		return "Sender's balance is insuffucient", http.StatusBadRequest, errors.New("insufficient Balance")
	}

	err = updateBalance(u, sender, -req.Amount)
	if err != nil {
		return "Failed to transfer balance", http.StatusInternalServerError, err
	}

	err = updateBalance(u, req.ToAccountNumber, req.Amount)
	if err != nil {
		return "Failed to transfer balance", http.StatusInternalServerError, err
	}

	return "Success Transfer Balance", http.StatusOK, nil
}

func updateBalance(u *uc, account int, amount int) error {
	balance := amount
	filter := bson.M{"account_number": account}
	data := bson.M{
		"$inc": bson.M{
			"balance": balance,
		},
	}

	err := u.repo.UpdateOne(data, filter, "account")
	if err != nil {
		return err
	}
	return nil
}
