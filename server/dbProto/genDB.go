package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type TableInfo struct {
	TableName string
	IsSave    bool
}

type DBMgr struct {
	TableInfos []TableInfo
	Table2Save map[string]bool
}

func (this *TableInfo) Show() {
	log.Print("TableName:", this.TableName, " IsSave:", this.IsSave)
}

func (this *DBMgr) Show() {
	for _, v := range this.TableInfos {
		v.Show()
	}
}

func (this *DBMgr) Check() {
	this.Table2Save = make(map[string]bool)
	for _, t := range this.TableInfos {
		_, ok := this.Table2Save[t.TableName]
		if ok {
			log.Panic("Table Name Duplicated ", t.TableName)
		}
		this.Table2Save[t.TableName] = t.IsSave
	}
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("LoadConfig >>>>>>", err)
		}
	}()
	content, err := ioutil.ReadFile("./user.msg")
	if err != nil {
		log.Panic("Read DB Config Error:", err)
	}
	dbMgr := new(DBMgr)
	err = json.Unmarshal(content, dbMgr)
	if err != nil {
		log.Panic("Unmarshal Config Error:", err)
	}

	dbMgr.Show()
	dbMgr.Check()
	//dbMgr.Gen()
}
