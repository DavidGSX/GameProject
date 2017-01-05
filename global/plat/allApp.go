package plat

import (
	"gameproject/global/config"
	"log"
	"net/url"
)

/*
	Author    string // 平台类型，填账号的后缀
*/
type AllApp struct {
	Plat
}

func (this *AllApp) Init(cfg *config.PlatConfig) {
	this.Author = cfg.Author

	if this.Author == "" {
		log.Panic("AllApp Author not set")
	}
}

func (this *AllApp) Clone() IClass {
	return new(AllApp)
}

func (this *AllApp) Authorize(uid, token string) string {
	return "ok"
}

func (this *AllApp) Callback(param url.Values) string {
	defer func() {
		if err := recover(); err != nil {
			log.Println("AllApp.Callback -> ", err)
		}
	}()

	return "AllApp Callback"
}
