package packages

import (
	"TransferSaldo/src/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func BasicResponse(w http.ResponseWriter, message string, code int, data interface{}) {
	res := models.BasicResponse{}

	res.Code = code
	res.Message = message
	res.Data = data

	response, err := json.Marshal(res)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "[SERVER] Failed Getting Response Data. Error Details : %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	fmt.Fprint(w, string(response))
}
