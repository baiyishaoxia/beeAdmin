package helper

import (
	"fmt"
)

func RedisErrorInfo(src string, infos interface{}) string {
	return fmt.Sprintf("[%s]: %s - Redis Query Fail : %v", GetNowTimeStamp(), src, infos)
}

func DBErrorInfo(src string, sql string, err interface{}) string {
	return fmt.Sprintf("[%s]: %s - DB Query Fail : sql: '%s' | error: '%s'", GetNowTimeStamp(), src, sql, err)
}

func UnmarshalErrorInfo(src string, infos interface{}) string {
	return fmt.Sprintf("[%s]: %s - Umarshal Fail : %v", GetNowTimeStamp(), src, infos)
}

func ESQueryErrorInfo(src string, infos interface{}) string {
	return fmt.Sprintf("[%s]: %s - ES Query Error : %v", GetNowTimeStamp(), src, infos)
}
