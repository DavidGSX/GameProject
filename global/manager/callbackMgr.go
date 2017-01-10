package manager

import (
	"gameproject/global/config"
	"io"
	"log"
	"net/http"
	"strconv"
)

func InitHttpCallback(cfg *config.GlobalConfig) {
	ip := cfg.BaseConfig.CallbackIP
	port := cfg.BaseConfig.CallbackPort

	http.HandleFunc("/", CallbackServer)

	log.Println("HttpCallback Listen ", ip, port)
	err := http.ListenAndServe(ip+":"+strconv.Itoa(int(port)), nil)
	if err != nil {
		log.Fatal("HttpCallback Listen Error:", err)
	}

}

func CallbackServer(w http.ResponseWriter, req *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Println("CallbackServer -> ", err)
		}
	}()

	req.ParseForm()
	result := GetPlatMgr().ProcessCallback(req.URL.Path, req.Form)
	io.WriteString(w, result)
}
