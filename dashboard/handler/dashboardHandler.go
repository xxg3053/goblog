package handler

import (
	"net/http"
	"github.com/xxg3053/goblog/common"
	"io/ioutil"
	"encoding/json"
)

func GetAccount(w http.ResponseWriter, r *http.Request) {

	//发起http请求
	// resp,err := http.Get("http://134.175.136.140:32410/accounts/1")
	resp,err := http.Get("http://accountservice.kenfo-test:6767/accounts/1")
	defer resp.Body.Close()
	if err != nil{
		common.WriteJsonResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil{
		common.WriteJsonResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	var result common.Result
	if err := json.Unmarshal(body, &result); err == nil {
		data := result.Data
		common.WriteJsonResponse(w, http.StatusOK, data)
	}else {
		common.WriteJsonResponse(w, http.StatusBadRequest, err.Error())
	}

}
