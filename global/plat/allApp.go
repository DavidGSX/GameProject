package plat

import (
	"gameproject/global/config"
	"log"
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

func (this *AllApp) Authorize(uid, token string) {

}

func (this *AllApp) Callback(msg string) {

}
