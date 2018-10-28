package handler

import (
	"net/http"
	"github.com/xxg3053/goblog/common"
)

type healthCheckResponse struct {
	Status string `json:"status"`
	ServBy string `json:"serv_by"`
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	//dbUp := DBClient.Check()
	servBy := common.GetIP()
	if true {
		common.WriteJsonResponse(w, http.StatusOK, healthCheckResponse{Status: "UP", ServBy: servBy})
	} else {
		common.WriteJsonResponse(w, http.StatusServiceUnavailable, healthCheckResponse{Status: "Database unaccessible", ServBy: servBy})
	}
}
