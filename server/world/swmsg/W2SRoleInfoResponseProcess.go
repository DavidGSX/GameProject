package swmsg

import (
	"gameproject/common"
	"gameproject/server/client/csproto"
	"log"

	"github.com/golang/protobuf/proto"
)

type W2SRoleInfoResponseProcess struct {
	msg   *W2SRoleInfoResponse
	trans *common.Trans
}

func (this *W2SRoleInfoResponseProcess) Process() bool {
	defer func() {
		if err := recover(); err != nil {
			log.Println("W2SRoleInfoResponseProcess Error:", err)
		}
	}()

	sendInfo := &csproto.SServerRoleInfos{}
	sendInfo.RoleId = this.msg.RoleId

	if this.msg.Info != nil {
		sendInfo.Info = make([]*csproto.SServerRoleInfos_RoleInfo, 0)
		for _, rInfo := range this.msg.Info {
			info := &csproto.SServerRoleInfos_RoleInfo{}
			info.ZoneId = rInfo.ZoneId
			info.RoleId = rInfo.RoleId
			info.RoleName = rInfo.RoleName
			info.Level = rInfo.Level
			info.School = rInfo.School
			info.Sex = rInfo.Sex
			info.Lasttime = rInfo.Lasttime
			sendInfo.Info = append(sendInfo.Info, info)
		}
	}

	data, err := proto.Marshal(sendInfo)
	if err != nil {
		log.Panic("Marshal SServerRoleInfos Error,", err)
	}
	oct := &common.Octets{}
	oct.MarshalUint32(uint32(len(data)))
	oct.MarshalUint32(1014)
	oct.MarshalBytesOnly(data)
	this.msg.Gets().SendByRoleIds([]uint64{sendInfo.RoleId}, oct.GetBuf())
	return true
}
