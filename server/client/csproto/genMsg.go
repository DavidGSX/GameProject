package csproto

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
	content = append(content, []byte("package msgMgr\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("import (\n")...)
	content = append(content, []byte("	\"gameproject/common\"\n")...)
	content = append(content, []byte("	\"log\"\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("	\"github.com/golang/protobuf/proto\"\n")...)
	content = append(content, []byte(")\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("// 避免与协议的函数名称重复，函数的命名有点特殊\n")...)
	content = append(content, []byte("type MsgInfo interface {\n")...)
	content = append(content, []byte("	Clone() MsgInfo\n")...)
	content = append(content, []byte("	MsgType() uint32\n")...)
	content = append(content, []byte("	GetMsg() proto.Message\n")...)
	content = append(content, []byte("	Setr(r uint64)\n")...)
	content = append(content, []byte("	Getr() uint64\n")...)
	content = append(content, []byte("	Setl(s ISend)\n")...)
	content = append(content, []byte("	Getl() ISend\n")...)
	content = append(content, []byte("	Setg(s ISend)\n")...)
	content = append(content, []byte("	Getg() ISend\n")...)
	content = append(content, []byte("	Setw(w ISend)\n")...)
	content = append(content, []byte("	Getw() ISend\n")...)
	content = append(content, []byte("	Unmarshal(data []byte) error\n")...)
	content = append(content, []byte("	Send2Link(MsgInfo) error\n")...)
	content = append(content, []byte("	Process(t *common.Trans) bool\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("type ISend interface {\n")...)
	content = append(content, []byte("	Send(x []byte)\n")...)
	content = append(content, []byte("	SetUserId(u string)\n")...)
	content = append(content, []byte("	GetUserId() string\n")...)
	content = append(content, []byte("	SetRoleId(r uint64)\n")...)
	content = append(content, []byte("	GetRoleId() uint64\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("var MsgInfos map[int]MsgInfo\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func AddMsg(t int, msg MsgInfo) {\n")...)
	content = append(content, []byte("	if MsgInfos == nil {\n")...)
	content = append(content, []byte("		MsgInfos = make(map[int]MsgInfo)\n")...)
	content = append(content, []byte("	}\n")...)
	content = append(content, []byte("	if _, ok := MsgInfos[t]; ok {\n")...)
	content = append(content, []byte("		log.Panic(\"Duplicate Msg Type\", t)\n")...)
	content = append(content, []byte("	}\n")...)
	content = append(content, []byte("	MsgInfos[t] = msg\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func GetMsg(t int) MsgInfo {\n")...)
	content = append(content, []byte("	if MsgInfos == nil {\n")...)
	content = append(content, []byte("		return nil\n")...)
	content = append(content, []byte("	}\n")...)
	content = append(content, []byte("	return MsgInfos[t].Clone()\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func MarshalMsg(msg MsgInfo) ([]byte, error) {\n")...)
	content = append(content, []byte("	data, err := proto.Marshal(msg.GetMsg())\n")...)
	content = append(content, []byte("	if err != nil {\n")...)
	content = append(content, []byte("		return nil, err\n")...)
	content = append(content, []byte("	}\n")...)
	content = append(content, []byte("	oct := &common.Octets{}\n")...)
	content = append(content, []byte("	oct.MarshalUint32(uint32(len(data)))\n")...)
	content = append(content, []byte("	oct.MarshalUint32(msg.MsgType())\n")...)
	content = append(content, []byte("	oct.MarshalBytesOnly(data)\n")...)
	content = append(content, []byte("	return oct.GetBuf(), nil\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("type IProcess interface {\n")...)
	content = append(content, []byte("	Clone() IProcess\n")...)
	content = append(content, []byte("	SetMsg(m MsgInfo)\n")...)
	content = append(content, []byte("	SetTrans(t *common.Trans)\n")...)
	content = append(content, []byte("	Process() bool\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("var ProcInfos map[string]IProcess\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func AddProc(s string, p IProcess) {\n")...)
	content = append(content, []byte("	if ProcInfos == nil {\n")...)
	content = append(content, []byte("		ProcInfos = make(map[string]IProcess)\n")...)
	content = append(content, []byte("	}\n")...)
	content = append(content, []byte("	if _, ok := ProcInfos[s]; ok {\n")...)
	content = append(content, []byte("		log.Panic(\"Duplicate Proc Type\", s)\n")...)
	content = append(content, []byte("	}\n")...)
	content = append(content, []byte("	ProcInfos[s] = p\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func GetProc(s string) IProcess {\n")...)
	content = append(content, []byte("	if ProcInfos == nil {\n")...)
	content = append(content, []byte("		return nil\n")...)
	content = append(content, []byte("	}\n")...)
	content = append(content, []byte("	return ProcInfos[s].Clone()\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)

	err := ioutil.WriteFile("../msgMgr/msgMgr.go", content, 0666)
	if err != nil {
		log.Panic("Write MsgMgr.go Error:", err)
	}

	msgContent := make([]byte, 0)
	msgContent = append(msgContent, []byte("package csmsg\n")...)
	msgContent = append(msgContent, []byte("\n")...)
	msgContent = append(msgContent, []byte("import \"gameproject/server/client/msgMgr\"\n")...)
	msgContent = append(msgContent, []byte("\n")...)
	msgContent = append(msgContent, []byte("func Init() {\n")...)

	procContent := make([]byte, 0)
	procContent = append(procContent, []byte("package csproc\n")...)
	procContent = append(procContent, []byte("\n")...)
	procContent = append(procContent, []byte("import \"gameproject/server/client/msgMgr\"\n")...)
	procContent = append(procContent, []byte("\n")...)
	procContent = append(procContent, []byte("func Init() {\n")...)
	for _, k := range types {
		v := this.Type2Name[k]
		this.GenMsgInfo(v)
		GenMsgProcess(v)
		msgContent = append(msgContent, []byte("	msgMgr.AddMsg("+strconv.Itoa(k)+", new("+v+"))\n")...)
		if len(v) > 0 && v[0] != 'S' {
			procContent = append(procContent, []byte("	msgMgr.AddProc(\""+v+"\", new("+v+"Process))\n")...)
		}
	}
	msgContent = append(msgContent, []byte("}\n")...)
	procContent = append(procContent, []byte("}\n")...)

	err = ioutil.WriteFile("../csmsg/msgInit.go", msgContent, 0666)
	if err != nil {
		log.Panic("Write msgInit.go Error:", err)
	}

	err = ioutil.WriteFile("../csproc/procInit.go", procContent, 0666)
	if err != nil {
		log.Panic("Write procInit.go Error:", err)
	}
}

func (this *MsgMgr) GenMsgInfo(name string) {
	content := make([]byte, 0)
	content = append(content, []byte("package csmsg\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("import (\n")...)
	content = append(content, []byte("	\"gameproject/common\"\n")...)
	content = append(content, []byte("	\"gameproject/server/client/csproto\"\n")...)
	content = append(content, []byte("	\"gameproject/server/client/msgMgr\"\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("	\"github.com/golang/protobuf/proto\"\n")...)
	content = append(content, []byte(")\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("type "+name+" struct {\n")...)
	content = append(content, []byte("	csproto."+name+"\n")...)
	content = append(content, []byte("	l msgMgr.ISend  // Link缩写\n")...)
	content = append(content, []byte("	g msgMgr.ISend  // Global缩写\n")...)
	content = append(content, []byte("	w msgMgr.ISend  // World缩写\n")...)
	content = append(content, []byte("	r uint64 // RoleId缩写\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func (this *"+name+") Clone() msgMgr.MsgInfo {\n")...)
	content = append(content, []byte("	return new("+name+")\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func (this *"+name+") MsgType() uint32 {\n")...)
	t, ok := this.Name2Type[name]
	if ok {
		content = append(content, []byte("	return "+strconv.Itoa(t)+"\n")...)
	} else {
		log.Panic("Get MsgType from Name Error:", name)
	}
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func (this *"+name+") GetMsg() proto.Message {\n")...)
	content = append(content, []byte("	return &this."+name+"\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("// 避免与协议的函数名称重复，以下函数命名有点特殊\n")...)
	content = append(content, []byte("func (this *"+name+") Setr(r uint64) {\n")...)
	content = append(content, []byte("	this.r = r\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func (this *"+name+") Getr() uint64 {\n")...)
	content = append(content, []byte("	return this.r\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func (this *"+name+") Setl(s msgMgr.ISend) {\n")...)
	content = append(content, []byte("	this.l = s\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func (this *"+name+") Getl() msgMgr.ISend {\n")...)
	content = append(content, []byte("	return this.l\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func (this *"+name+") Setg(s msgMgr.ISend) {\n")...)
	content = append(content, []byte("	this.g = s\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func (this *"+name+") Getg() msgMgr.ISend {\n")...)
	content = append(content, []byte("	return this.g\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func (this *"+name+") Setw(w msgMgr.ISend) {\n")...)
	content = append(content, []byte("	this.w = w\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func (this *"+name+") Getw() msgMgr.ISend {\n")...)
	content = append(content, []byte("	return this.w\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func (this *"+name+") Unmarshal(data []byte) error {\n")...)
	content = append(content, []byte("	err := proto.Unmarshal(data, &this."+name+")\n")...)
	content = append(content, []byte("	return err\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func (this *"+name+") Send2Link(msg msgMgr.MsgInfo) error {\n")...)
	content = append(content, []byte("	data, err := msgMgr.MarshalMsg(msg)\n")...)
	content = append(content, []byte("	if err != nil {\n")...)
	content = append(content, []byte("		return err\n")...)
	content = append(content, []byte("	}\n")...)
	content = append(content, []byte("	this.Getl().Send(data)\n")...)
	content = append(content, []byte("	return nil\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func (this *"+name+") Process(t *common.Trans) bool {\n")...)
	if len(name) > 0 && name[0] == 'S' {
		content = append(content, []byte("	// do nothing\n")...)
		content = append(content, []byte("	return false\n")...)
	} else {
		content = append(content, []byte("	p := msgMgr.GetProc(\""+name+"\")\n")...)
		content = append(content, []byte("	p.SetMsg(this)\n")...)
		content = append(content, []byte("	p.SetTrans(t)\n")...)
		content = append(content, []byte("	return p.Process()\n")...)
	}
	content = append(content, []byte("}\n")...)

	fileName := "../csmsg/" + name + ".go"
	err := ioutil.WriteFile(fileName, content, 0666)
	if err != nil {
		log.Panic("Write "+name+".go Error:", err)
	}
}

func GenMsgProcess(name string) {
	// S开头的协议不用生成Process文件
	if len(name) > 0 && name[0] == 'S' {
		return
	}

	// 文件存在，可能有写具体的处理逻辑，生成代码会覆盖，所以直接返回
	filename := "../csproc/" + name + "Process.go"
	_, err := os.Stat(filename)
	if err == nil || os.IsNotExist(err) == false {
		return
	}

	content := make([]byte, 0)
	content = append(content, []byte("package csproc\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("import (\n")...)
	content = append(content, []byte("	\"gameproject/common\"\n")...)
	content = append(content, []byte("	\"gameproject/server/client/csmsg\"\n")...)
	content = append(content, []byte("	\"gameproject/server/client/msgMgr\"\n")...)
	content = append(content, []byte("	\"log\"\n")...)
	content = append(content, []byte(")\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("type "+name+"Process struct {\n")...)
	content = append(content, []byte("	msg   *csmsg."+name+"\n")...)
	content = append(content, []byte("	trans *common.Trans\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func (this *"+name+"Process) Clone() msgMgr.IProcess {\n")...)
	content = append(content, []byte("	return new("+name+"Process)\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func (this *"+name+"Process) SetMsg(m msgMgr.MsgInfo) {\n")...)
	content = append(content, []byte("	this.msg = m.(*csmsg."+name+")\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func (this *"+name+"Process) SetTrans(t *common.Trans) {\n")...)
	content = append(content, []byte("	this.trans = t\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func (this *"+name+"Process) Process() bool {\n")...)
	content = append(content, []byte("	defer func() {\n")...)
	content = append(content, []byte("		if err := recover(); err != nil {\n")...)
	content = append(content, []byte("			log.Println(\""+name+"Process Error:\", err)\n")...)
	content = append(content, []byte("		}\n")...)
	content = append(content, []byte("	}()\n")...)
	content = append(content, []byte("	\n")...)
	content = append(content, []byte("	log.Println(\"to do "+name+"Process\")\n")...)
	content = append(content, []byte("	return true\n")...)
	content = append(content, []byte("}\n")...)

	err = ioutil.WriteFile(filename, content, 0666)
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
