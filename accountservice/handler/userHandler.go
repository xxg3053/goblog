package handler

import (
	"net/http"
	"github.com/xxg3053/goblog/accountservice/model"
	"github.com/xxg3053/goblog/common"
)

func GetAccount(w http.ResponseWriter, r *http.Request) {

	common.WriteJsonResponse(w, http.StatusOK, model.Account{Id:1,Name:"kenfo"})
}
