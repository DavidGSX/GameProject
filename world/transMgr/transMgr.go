package transMgr

import (
	"gameproject/world/lockMgr"
	"log"
	"time"
)

type tMsg interface {
	Process(t *Trans) bool
	MsgType() uint32
}

type tSave interface {
	Save() error
}

type Trans struct {
	begin  time.Time
	keys   []string
	tables []tSave
}

func NewTrans() *Trans {
	t := &Trans{}
	t.begin = time.Now()
	t.keys = make([]string, 0)
	t.tables = make([]tSave, 0)
	return t
}

func (this *Trans) Process(msg tMsg) {
	defer this.Unlock()
	if msg.Process(this) {
		for _, t := range this.tables {
			t.Save()
		}
	}
	if time.Since(this.begin) > 50*time.Millisecond {
		log.Println("msg ", msg.MsgType(), " process use ", time.Since(this.begin))
	}
}

func (this *Trans) Lock(names ...string) {
	tmpKey := make([]string, 0)
	for _, name := range names {
		var bFind bool = false
		for _, key := range this.keys {
			if key == name {
				bFind = true
				break
			}
		}
		if bFind == false {
			tmpKey = append(tmpKey, name)
		}
	}

	if len(tmpKey) == 0 {
		return
	}

	this.keys = append(this.keys, tmpKey...)
	lockMgr.Lock(tmpKey...)
}

func (this *Trans) Unlock() {
	if len(this.keys) == 0 {
		return
	}
	lockMgr.Unlock(this.keys...)
}

func (this *Trans) Save(t tSave) {
	if t == nil {
		return
	}
	this.tables = append(this.tables, t)
}
