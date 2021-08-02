package usecases

import (
	"TransferSaldo/src/models"
	"errors"
	"strconv"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

func (u *uc) GetAccountID(urlstring string) (int, error) {
	urlPart := strings.Split(urlstring, "/")
	account_id, err := strconv.Atoi(urlPart[2])
	if err != nil {
		account_id, err = strconv.Atoi(urlPart[4])
		if err != nil {
			return 0, errors.New("failed to Get Params")
		}
	}

	return account_id, nil
}

func (u *uc) GetBalanceInfo(account_number int) (*models.GetBalance, error) {
	var (
		res models.GetBalance
		acc models.AccountModel
		cus models.CustomerModel
	)
	err := u.repo.FindOne(&acc, bson.M{"account_number": account_number}, "account")
	if err != nil {
		return nil, err
	}

	res.AccountNumber = acc.AccountNumber
	res.Balance = acc.Balance

	customer_number := acc.CustomerNumber

	err = u.repo.FindOne(&cus, bson.M{"customer_number": customer_number}, "customer")
	if err != nil {
		return nil, err
	}
	res.CustomerName = cus.CustomerName
	return &res, nil
}
