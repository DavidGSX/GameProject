package protocol

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
)

type MsgInfo struct {
	MsgType int
	MsgName string
}

type MsgMgr struct {
	MsgInfos  []MsgInfo
	Type2Name map[int]string
	Name2Type map[string]int
}

func (this *MsgInfo) Show() {
	log.Println("MsgType:", this.MsgType, "MsgName:", this.MsgName)
}

func (this *MsgMgr) Show() {
	for _, v := range this.MsgInfos {
		v.Show()
	}
}

func (this *MsgMgr) Check() {
	this.Type2Name = make(map[int]string)
	this.Name2Type = make(map[string]int)
	for _, msg := range this.MsgInfos {
		_, ok := this.Type2Name[msg.MsgType]
		if ok {
			log.Panic("Msg Type Duplicated ", msg.MsgType)
		}
		this.Type2Name[msg.MsgType] = msg.MsgName

		_, ok = this.Name2Type[msg.MsgName]
		if ok {
			log.Panic("Msg Name Duplicated ", msg.MsgName)
		}
		this.Name2Type[msg.MsgName] = msg.MsgType
	}
}

func (this *MsgMgr) Gen() {
	types := make([]int, 0)
	for k, _ := range this.Type2Name {
		types = append(types, k)
	}
	sort.Ints(types)

	content := make([]byte, 0)
	content = append(content, []byte("package message\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("type MsgInfo interface {\n")...)
	content = append(content, []byte("	Clone() MsgInfo\n")...)
	content = append(content, []byte("	SetRoleId(r uint64)\n")...)
	content = append(content, []byte("	SetLink(s ISend)\n")...)
	content = append(content, []byte("	SetGlobal(s ISend)\n")...)
	content = append(content, []byte("	Unmarshal(data []byte) error\n")...)
	content = append(content, []byte("	Process()\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("type ISend interface {\n")...)
	content = append(content, []byte("	Send(x []byte)\n")...)
	content = append(content, []byte("	SetUserId(u string)\n")...)
	content = append(content, []byte("	GetUserId() string\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("var MsgInfos map[int]MsgInfo\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func Init() {\n")...)
	content = append(content, []byte("	MsgInfos = make(map[int]MsgInfo)\n")...)
	for _, k := range types {
		v := this.Type2Name[k]
		content = append(content, []byte("	MsgInfos["+strconv.Itoa(k)+"] = new("+v+")\n")...)
		GenMsgInfo(v)
		GenMsgProcess(v)
	}
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func GetMsg(t int) MsgInfo {\n")...)
	content = append(content, []byte("	if MsgInfos == nil {\n")...)
	content = append(content, []byte("		return nil\n")...)
	content = append(content, []byte("	}\n")...)
	content = append(content, []byte("	return MsgInfos[t]\n")...)
	content = append(content, []byte("}\n")...)

	err := ioutil.WriteFile("../message/msgMgr.go", content, 0666)
	if err != nil {
		log.Panic("Write MsgMgr.go Error:", err)
	}
}

func GenMsgInfo(name string) {
	content := make([]byte, 0)
	content = append(content, []byte("package message\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("import (\n")...)
	content = append(content, []byte("	\"gameproject/server/protocol\"\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("	\"github.com/golang/protobuf/proto\"\n")...)
	content = append(content, []byte(")\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("type "+name+" struct {\n")...)
	content = append(content, []byte("	Link   ISend\n")...)
	content = append(content, []byte("	Global ISend\n")...)
	content = append(content, []byte("	RoleId uint64\n")...)
	content = append(content, []byte("	Proto  protocol."+name+"\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func (this *"+name+") Clone() MsgInfo {\n")...)
	content = append(content, []byte("	return new("+name+")\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func (this *"+name+") SetRoleId(r uint64) {\n")...)
	content = append(content, []byte("	this.RoleId = r\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func (this *"+name+") SetLink(s ISend) {\n")...)
	content = append(content, []byte("	this.Link = s\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func (this *"+name+") SetGlobal(s ISend) {\n")...)
	content = append(content, []byte("	this.Global = s\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func (this *"+name+") Unmarshal(data []byte) error {\n")...)
	content = append(content, []byte("	err := proto.Unmarshal(data, &this.Proto)\n")...)
	content = append(content, []byte("	return err\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func (this *"+name+") Process() {\n")...)
	if len(name) > 0 && name[0] == 'S' {
		content = append(content, []byte("	//")...)
	}
	content = append(content, []byte("	new("+name+"Process).Process(this)\n")...)
	content = append(content, []byte("}\n")...)

	filename := "../message/" + name + ".go"
	err := ioutil.WriteFile(filename, content, 0666)
	if err != nil {
		log.Panic("Write MsgMgr.go Error:", err)
	}
}

func GenMsgProcess(name string) {
	// S开头的协议不用生成Process文件
	if len(name) > 0 && name[0] == 'S' {
		return
	}

	// 文件存在，可能有写具体的处理逻辑，生成代码会覆盖，所以直接返回
	filename := "../message/" + name + "Process.go"
	_, err := os.Stat(filename)
	if err == nil || os.IsNotExist(err) == false {
		return
	}

	content := make([]byte, 0)
	content = append(content, []byte("package message\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("import (\n")...)
	content = append(content, []byte("	\"log\"\n")...)
	content = append(content, []byte(")\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("type "+name+"Process struct {\n")...)
	content = append(content, []byte("	msg *"+name+"\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func (this *"+name+"Process) Process(msg *"+name+") {\n")...)
	content = append(content, []byte("	this.msg = msg\n")...)
	content = append(content, []byte("	log.Println(\"to do "+name+"Process\")\n")...)
	content = append(content, []byte("}\n")...)

	err = ioutil.WriteFile(filename, content, 0666)
	if err != nil {
		log.Panic("Write MsgMgr.go Error:", err)
	}
}

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("LoadConfig >>>>>>", err)
		}
	}()
	content, err := ioutil.ReadFile("./login.msg")
	if err != nil {
		log.Panic("Read Msg Config Error:", err)
	}
	msgMgr := new(MsgMgr)
	err = json.Unmarshal(content, msgMgr)
	if err != nil {
		log.Panic("Unmarshal Config Error:", err)
	}

	msgMgr.Show()
	msgMgr.Check()
	msgMgr.Gen()
}
