package cache

import (
	"gameproject/common"
	"log"
	"strconv"
)

var (
	ZoneId uint32
)

func SetZoneId(z uint32) {
	ZoneId = z
}

func GetUniqIdByTableName(key string) (r uint64) {
	common.Lock(key)
	defer common.Unlock(key)

	v := GetKV(key)
	if v == "" {
		r = uint64(ZoneId)
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
