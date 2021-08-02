package usecases

import (
	"TransferSaldo/packages"
	"TransferSaldo/src/models"
	"TransferSaldo/src/repositories"
	"fmt"
	"net/http"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func init() {
	env := map[string]string{
		"SERVER_PORT": "8080",
		"MONGO_DB":    "LinkAja",
		"MONGO_URI":   "mongodb+srv://teredict:25285282@teredict.jkdyu.mongodb.net/LinkAja?retryWrites=true&w=majority",
		"CORS_HEADER": "x-api-key,Authorization,Content-Type,Origin,Accept,Access-Control-Allow-Headers,Access-Control-Request-Method,Access-Control-Request-Headers,Access-Control-Allow-Origin",
		"CORS_METHOD": "OPTION,GET,PUT,POST,DELETE",
		"CORS_ORIGIN": "*",
	}

	for key, value := range env {
		os.Setenv(key, value)
	}
	packages.LoadMongo()
}

func TestTransferBalance(t *testing.T) {
	db := packages.LoadDatabase()
	repository := repositories.NewRepo(db)
	usecase := NewUC(repository)
	type args struct {
		sender int
		req    *models.TransferBalanceModel
	}
	tests := []struct {
		name    string
		args    args
		want1   string
		want2   int
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				sender: 555002,
				req: &models.TransferBalanceModel{
					ToAccountNumber: 555001,
					Amount:          500,
				},
			},
			want1:   "Success Transfer Balance",
			want2:   http.StatusOK,
			wantErr: false,
		},
		{
			name: "Not Found",
			args: args{
				sender: 123456,
				req: &models.TransferBalanceModel{
					ToAccountNumber: 555001,
					Amount:          500,
				},
			},
			want1:   "Cannot fetch sender ID from database",
			want2:   http.StatusBadRequest,
			wantErr: true,
		},
		{
			name: "Insufficient Balance",
			args: args{
				sender: 555002,
				req: &models.TransferBalanceModel{
					ToAccountNumber: 555001,
					Amount:          5000000,
				},
			},
			want1:   "Sender's balance is insuffucient",
			want2:   http.StatusBadRequest,
			wantErr: true,
		},
	}

	for _, testitem := range tests {
		t.Run(testitem.name, func(t *testing.T) {
			got1, got2, err := usecase.TransferBalance(testitem.args.sender, testitem.args.req)
			if (err != nil) != testitem.wantErr {
				t.Errorf("GetAccountID() Method error got %v, wantErr %v", err, testitem.wantErr)
				return
			}
			assert.Equal(t, testitem.want1, got1, "Wrong Message")
			assert.Equal(t, testitem.want2, got2, "Wrong Code")
		})
	}

}

func TestUpdateBalance(t *testing.T) {
	db := packages.LoadDatabase()
	repository := repositories.NewRepo(db)
	type args struct {
		uc      uc
		account int
		amount  int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				uc: uc{
					repo: repository,
				},
				account: 555002,
				amount:  500,
			},
			wantErr: false,
		},
		{
			name: "Not Found",
			args: args{
				uc: uc{
					repo: repository,
				},
				account: 123456,
				amount:  500,
			},
			wantErr: false,
		},
	}
	for _, testitem := range tests {
		t.Run(testitem.name, func(t *testing.T) {
			err := updateBalance(&testitem.args.uc, testitem.args.account, testitem.args.amount)
			fmt.Println("ERROR=>", err)
			if (err != nil) != testitem.wantErr {
				t.Errorf("updateBalance() Method error got %v, wantErr %v", err, testitem.wantErr)
				return
			}
		})
	}
}
