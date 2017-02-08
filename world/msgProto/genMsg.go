package msgProto

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
	content = append(content, []byte("import (\n")...)
	content = append(content, []byte("	\"gameproject/common\"\n")...)
	content = append(content, []byte("	\"gameproject/world/transMgr\"\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("	\"github.com/golang/protobuf/proto\"\n")...)
	content = append(content, []byte(")\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("// 避免与协议的函数名称重复，函数的命名有点特殊\n")...)
	content = append(content, []byte("type MsgInfo interface {\n")...)
	content = append(content, []byte("	Clone() MsgInfo\n")...)
	content = append(content, []byte("	MsgType() uint32\n")...)
	content = append(content, []byte("	GetMsg() proto.Message\n")...)
	content = append(content, []byte("	Sets(s ISend)\n")...)
	content = append(content, []byte("	Gets() ISend\n")...)
	content = append(content, []byte("	Unmarshal(data []byte) error\n")...)
	content = append(content, []byte("	Send(MsgInfo) error\n")...)
	content = append(content, []byte("	Process(t *transMgr.Trans) bool\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("type ISend interface {\n")...)
	content = append(content, []byte("	Send(b []byte)\n")...)
	content = append(content, []byte("	SendByZoneIds(zoneIds []uint32, b []byte)\n")...)
	content = append(content, []byte("	SetZoneId(z uint32)\n")...)
	content = append(content, []byte("	GetZoneId() uint32\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func GetMsgByte(msg MsgInfo) (error, []byte) {\n")...)
	content = append(content, []byte("	data, err := proto.Marshal(msg.GetMsg())\n")...)
	content = append(content, []byte("	if err != nil {\n")...)
	content = append(content, []byte("		return err, nil\n")...)
	content = append(content, []byte("	}\n")...)
	content = append(content, []byte("	oct := &common.Octets{}\n")...)
	content = append(content, []byte("	oct.MarshalUint32(uint32(len(data)))\n")...)
	content = append(content, []byte("	oct.MarshalUint32(msg.MsgType())\n")...)
	content = append(content, []byte("	oct.MarshalBytesOnly(data)\n")...)
	content = append(content, []byte("	return nil, oct.GetBuf()\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("var MsgInfos map[int]MsgInfo\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func Init() {\n")...)
	content = append(content, []byte("	MsgInfos = make(map[int]MsgInfo)\n")...)
	for _, k := range types {
		v := this.Type2Name[k]
		content = append(content, []byte("	MsgInfos["+strconv.Itoa(k)+"] = new("+v+")\n")...)
		this.GenMsgInfo(v)
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

func (this *MsgMgr) GenMsgInfo(name string) {
	content := make([]byte, 0)
	content = append(content, []byte("package message\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("import (\n")...)
	content = append(content, []byte("	\"gameproject/common\"\n")...)
	content = append(content, []byte("	\"gameproject/world/msgProto\"\n")...)
	content = append(content, []byte("	\"gameproject/world/transMgr\"\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("	\"github.com/golang/protobuf/proto\"\n")...)
	content = append(content, []byte(")\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("type "+name+" struct {\n")...)
	content = append(content, []byte("	msgProto."+name+"\n")...)
	content = append(content, []byte("	s ISend  // Server缩写\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func (this *"+name+") Clone() MsgInfo {\n")...)
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
	content = append(content, []byte("func (this *"+name+") Sets(s ISend) {\n")...)
	content = append(content, []byte("	this.s = s\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func (this *"+name+") Gets() ISend {\n")...)
	content = append(content, []byte("	return this.s\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func (this *"+name+") Unmarshal(data []byte) error {\n")...)
	content = append(content, []byte("	err := proto.Unmarshal(data, &this."+name+")\n")...)
	content = append(content, []byte("	return err\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func (this *"+name+") Send(msg MsgInfo) error {\n")...)
	content = append(content, []byte("	data, err := proto.Marshal(msg.GetMsg())\n")...)
	content = append(content, []byte("	if err != nil {\n")...)
	content = append(content, []byte("		return err\n")...)
	content = append(content, []byte("	}\n")...)
	content = append(content, []byte("	oct := &common.Octets{}\n")...)
	content = append(content, []byte("	oct.MarshalUint32(uint32(len(data)))\n")...)
	content = append(content, []byte("	oct.MarshalUint32(msg.MsgType())\n")...)
	content = append(content, []byte("	oct.MarshalBytesOnly(data)\n")...)
	content = append(content, []byte("	this.Gets().Send(oct.GetBuf())\n")...)
	content = append(content, []byte("	return nil\n")...)
	content = append(content, []byte("}\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("func (this *"+name+") Process(t *transMgr.Trans) bool {\n")...)
	if len(name) > 0 && name[0] == 'W' {
		content = append(content, []byte("	// do nothing\n")...)
		content = append(content, []byte("	return false\n")...)
	} else {
		content = append(content, []byte("	p := new("+name+"Process)\n")...)
		content = append(content, []byte("	p.msg = this\n")...)
		content = append(content, []byte("	p.trans = t\n")...)
		content = append(content, []byte("	return p.Process()\n")...)
	}
	content = append(content, []byte("}\n")...)

	filename := "../message/" + name + ".go"
	err := ioutil.WriteFile(filename, content, 0666)
	if err != nil {
		log.Panic("Write "+name+".go Error:", err)
	}
}

func GenMsgProcess(name string) {
	// S开头的协议不用生成Process文件
	if len(name) > 0 && name[0] == 'W' {
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
	content = append(content, []byte("	\"gameproject/world/transMgr\"\n")...)
	content = append(content, []byte("	\"log\"\n")...)
	content = append(content, []byte(")\n")...)
	content = append(content, []byte("\n")...)
	content = append(content, []byte("type "+name+"Process struct {\n")...)
	content = append(content, []byte("	msg   *"+name+"\n")...)
	content = append(content, []byte("	trans *transMgr.Trans\n")...)
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
	content, err := ioutil.ReadFile("./world.msg")
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
