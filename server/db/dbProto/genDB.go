package dbProto

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type TableInfo struct {
	TableName string
	KeyType   string
	IsSave    bool
}

type DBMgr struct {
	TableInfos []TableInfo
	Table2Save map[string]bool
}

func (this *TableInfo) Show() {
	log.Print("TableName:", this.TableName, " KeyType:", this.KeyType, " IsSave:", this.IsSave)
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

func (this *DBMgr) Gen() {
	for _, t := range this.TableInfos {
		this.GenTable(t.TableName, t.KeyType, t.IsSave)
	}
}

func (this *DBMgr) GenTable(name, keyType string, isSave bool) {
	content := make([]byte, 0)
	content = append(content, []byte("package table\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("import (\n")...)
	content = append(content, []byte("	\"gameproject/common\"\n")...)
	content = append(content, []byte("	\"gameproject/server/db/cacheMgr\"\n")...)
	content = append(content, []byte("	\"gameproject/server/db/dbProto\"\n")...)
	content = append(content, []byte("	\"log\"\n")...)
	switch keyType {
	case "uint64":
		content = append(content, []byte("	\"strconv\"\n")...)
	}
	content = append(content, []byte("\n")...)
	content = append(content, []byte("	\"github.com/golang/protobuf/proto\"\n")...)
	content = append(content, []byte(")\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("type "+name+" struct {\n")...)
	content = append(content, []byte("	dbProto."+name+"\n")...)
	content = append(content, []byte("	k string\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func (this *"+name+") IsSave() bool {\n")...)
	if isSave {
		content = append(content, []byte("	return true\n")...)
	} else {
		content = append(content, []byte("	return false\n")...)
	}
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)

	switch keyType {
	case "string":
		content = append(content, []byte("func New"+name+"(t *common.Trans, k string) *"+name+" {\n")...)
		content = append(content, []byte("	r := new("+name+")\n")...)
		content = append(content, []byte("	r.k = \""+name+"_\" + k\n")...)
		content = append(content, []byte("	if t != nil {\n")...)
		content = append(content, []byte("		t.Save(r)\n")...)
		content = append(content, []byte("	}\n")...)
		content = append(content, []byte("	return r\n")...)
		content = append(content, []byte("}\n")...)
		content = append(content, []byte("\n")...)
		content = append(content, []byte("func Get"+name+"(t *common.Trans, k string) *"+name+" {\n")...)
		content = append(content, []byte("	if k == \"\" {\n")...)
		content = append(content, []byte("		return nil\n")...)
		content = append(content, []byte("	}\n")...)
		content = append(content, []byte("	key := \""+name+"_\" + k\n")...)
	case "uint64":
		content = append(content, []byte("func New"+name+"(t *common.Trans, k uint64) *"+name+" {\n")...)
		content = append(content, []byte("	r := new("+name+")\n")...)
		content = append(content, []byte("	r.k = \""+name+"_\" + strconv.FormatUint(k, 10)\n")...)
		content = append(content, []byte("	if t != nil {\n")...)
		content = append(content, []byte("		t.Save(r)\n")...)
		content = append(content, []byte("	}\n")...)
		content = append(content, []byte("	return r\n")...)
		content = append(content, []byte("}\n")...)
		content = append(content, []byte("\n")...)
		content = append(content, []byte("func Get"+name+"(t *common.Trans, k uint64) *"+name+" {\n")...)
		content = append(content, []byte("	if k == 0 {\n")...)
		content = append(content, []byte("		return nil\n")...)
		content = append(content, []byte("	}\n")...)
		content = append(content, []byte("	key := \""+name+"_\" + strconv.FormatUint(k, 10)\n")...)
	default:
		log.Panic("Unkown Key Type table:", name)
	}
	content = append(content, []byte("	t.Lock(key)\n")...)
	content = append(content, []byte("	v := cacheMgr.GetKV(key)\n")...)
	content = append(content, []byte("	if v == \"\" {\n")...)
	content = append(content, []byte("		return nil\n")...)
	content = append(content, []byte("	}\n")...)
	content = append(content, []byte("	\n")...)
	content = append(content, []byte("	oct := common.NewOctets([]byte(v))\n")...)
	content = append(content, []byte("	size := oct.UnmarshalUint32()\n")...)
	content = append(content, []byte("	if size != oct.Remain() {\n")...)
	content = append(content, []byte("		log.Panic(\"get table."+name+" Data Len Error:\", key, \",\", size, \",\", len(v))\n")...)
	content = append(content, []byte("		return nil\n")...)
	content = append(content, []byte("	}\n")...)
	content = append(content, []byte("	data := oct.UnmarshalBytesOnly(size)\n")...)
	content = append(content, []byte("	r := New"+name+"(t, k)\n")...)
	content = append(content, []byte("	err := proto.Unmarshal(data, &r."+name+")\n")...)
	content = append(content, []byte("	if err != nil {\n")...)
	content = append(content, []byte("		log.Panic(\"get DB Data Unmarshal Error:\", r.k)\n")...)
	content = append(content, []byte("		return nil\n")...)
	content = append(content, []byte("	}\n")...)
	content = append(content, []byte("	return r\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	switch keyType {
	case "string":
		content = append(content, []byte("func Select"+name+"(k string) *"+name+" {\n")...)
		content = append(content, []byte("	if k == \"\" {\n")...)
		content = append(content, []byte("		return nil\n")...)
		content = append(content, []byte("	}\n")...)
		content = append(content, []byte("	key := \""+name+"_\" + k\n")...)
	case "uint64":
		content = append(content, []byte("func Select"+name+"(k uint64) *"+name+" {\n")...)
		content = append(content, []byte("	if k == 0 {\n")...)
		content = append(content, []byte("		return nil\n")...)
		content = append(content, []byte("	}\n")...)
		content = append(content, []byte("	key := \""+name+"_\" + strconv.FormatUint(k, 10)\n")...)
	default:
		log.Panic("Unkown Key Type table:", name)
	}
	content = append(content, []byte("	common.Lock(key)\n")...)
	content = append(content, []byte("	defer common.Unlock(key)\n")...)
	content = append(content, []byte("	v := cacheMgr.GetKV(key)\n")...)
	content = append(content, []byte("	if v == \"\" {\n")...)
	content = append(content, []byte("		return nil\n")...)
	content = append(content, []byte("	}\n")...)
	content = append(content, []byte("	\n")...)
	content = append(content, []byte("	oct := common.NewOctets([]byte(v))\n")...)
	content = append(content, []byte("	size := oct.UnmarshalUint32()\n")...)
	content = append(content, []byte("	if size != oct.Remain() {\n")...)
	content = append(content, []byte("		log.Panic(\"select table."+name+" Data Len Error:\", key, \",\", size, \",\", len(v))\n")...)
	content = append(content, []byte("		return nil\n")...)
	content = append(content, []byte("	}\n")...)
	content = append(content, []byte("	data := oct.UnmarshalBytesOnly(size)\n")...)
	content = append(content, []byte("	r := New"+name+"(nil, k)\n")...)
	content = append(content, []byte("	err := proto.Unmarshal(data, &r."+name+")\n")...)
	content = append(content, []byte("	if err != nil {\n")...)
	content = append(content, []byte("		log.Panic(\"select DB Data Unmarshal Error:\", r.k)\n")...)
	content = append(content, []byte("		return nil\n")...)
	content = append(content, []byte("	}\n")...)
	content = append(content, []byte("	return r\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func (this *"+name+") Save() error {\n")...)
	content = append(content, []byte("	if this.k == \"\" {\n")...)
	content = append(content, []byte("		log.Panic(\"DB Data Save Error:\", this.k)\n")...)
	content = append(content, []byte("	}\n")...)
	content = append(content, []byte("	data, err := proto.Marshal(&this."+name+")\n")...)
	content = append(content, []byte("	if err != nil {\n")...)
	content = append(content, []byte("		return err\n")...)
	content = append(content, []byte("	}\n")...)
	content = append(content, []byte("	oct := &common.Octets{}\n")...)
	content = append(content, []byte("	oct.MarshalUint32(uint32(len(data)))\n")...)
	content = append(content, []byte("	oct.MarshalBytesOnly(data)\n")...)
	content = append(content, []byte("	cacheMgr.SetKV(this.k, string(oct.GetBuf()))\n")...)
	content = append(content, []byte("	return nil\n")...)
	content = append(content, []byte("}\n")...)

	filename := "../table/" + name + ".go"
	err := ioutil.WriteFile(filename, content, 0666)
	if err != nil {
		log.Panic("Write "+name+".go Error:", err)
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
	dbMgr.Gen()
}
