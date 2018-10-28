package common

import (
	"net"
	"net/http"
	"strconv"
	"encoding/json"
)

func GetIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "error"
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	panic("Unable to determine local IP address (non loopback). Exiting.")
}

type result struct {
	Code int `json:"code"`
	Msg string `json:"msg"`
	Data interface{} `json:"data"`
} 
func WriteJsonResponse(w http.ResponseWriter, status int, data interface{}) {
	rt,_ := json.Marshal(&result{Code:status, Msg:"response success", Data:data})

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Content-Length", strconv.Itoa(len(rt)))
	w.WriteHeader(status)
	w.Write(rt)
}
