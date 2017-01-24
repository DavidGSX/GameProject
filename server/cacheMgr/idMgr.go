package cacheMgr

import (
	"gameproject/server/config"
	"gameproject/server/lockMgr"
	"log"
	"strconv"
)

func GetUniqIdByTableName(key string) (r uint64) {
	lockMgr.Lock(key)
	defer lockMgr.Unlock(key)

	v := GetKV(key)
	if v == "" {
		r = uint64(config.GetConfig().GetZoneId())
	} else {
		last, err := strconv.ParseUint(v, 10, 64)
		if err != nil {
			log.Panic("GetUniqIdByTableName Err, table:", key, " dbValue", v)
		}
		r = last + 4096
	}
	SetKV(key, strconv.FormatUint(r, 10))
	return r
}

func GetNextRoleId() uint64 {
	return GetUniqIdByTableName("SYS_ROLEID")
}
