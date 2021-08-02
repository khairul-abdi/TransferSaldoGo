package usecases

import (
	"TransferSaldo/packages"
	"TransferSaldo/src/models"
	"TransferSaldo/src/repositories"
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

func TestAccountID(t *testing.T) {
	db := packages.LoadDatabase()
	repository := repositories.NewRepo(db)
	usecase := NewUC(repository)
	type args struct {
		urlstring string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Sucess",
			args: args{
				urlstring: "http://localhost:8080/account/555001",
			},
			want:    555001,
			wantErr: false,
		},
		{
			name: "Success",
			args: args{
				urlstring: "localhost:8080/account/555002",
			},
			want:    555002,
			wantErr: false,
		},
		{
			name: "Failed",
			args: args{
				urlstring: "http://localhost:8080/account/555001u",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, testitem := range tests {
		t.Run(
			testitem.name, func(t *testing.T) {
				got1, err := usecase.GetAccountID(testitem.args.urlstring)
				if (err != nil) != testitem.wantErr {
					t.Errorf("GetAccountID() Method error got %v, wantErr %v", err, testitem.wantErr)
					return
				}
				assert.Equal(t, testitem.want, got1, "Cannot Get Correct Account ID")
			})
	}
}

func TestGetBalanceInfo(t *testing.T) {
	db := packages.LoadDatabase()
	repository := repositories.NewRepo(db)
	usecase := NewUC(repository)

	type args struct {
		account_number int
	}
	tests := []struct {
		name    string
		args    args
		want    *models.GetBalance
		wantErr bool
	}{
		{
			name: "Success",
			args: args{
				account_number: 555001,
			},
			want: &models.GetBalance{
				AccountNumber: 555001,
				CustomerName:  "Bob Martin",
				Balance:       9000,
			},
			wantErr: false,
		},
		{
			name: "Not Found",
			args: args{
				account_number: 123456,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, testitem := range tests {
		t.Run(testitem.name, func(t *testing.T) {
			got1, err := usecase.GetBalanceInfo(testitem.args.account_number)
			if (err != nil) != testitem.wantErr {
				t.Errorf("GetBalanceInfo() Method error got %v, wantErr %v", err, testitem.wantErr)
				return
			}
			assert.Equal(t, testitem.want, got1, "Cannot Get Desired Balance")
		})
	}
}
