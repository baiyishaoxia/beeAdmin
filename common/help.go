package common

import (
	"beeAdmin/common/helper"
	"beeAdmin/vendors/caches"
	"fmt"
	"github.com/astaxie/beego/logs"
	"os"
	"strconv"
	"strings"
	"time"
)

//数据data的基本属性
type Response struct {
	Errors []string    `json:"errors,omitempty"`
	Time   string      `json:"time"`
	List   interface{} `json:"list"`
}

//返回json的统一格式
type JSONS struct {
	Status int      `json:"status"`
	Msg    string   `json:"msg"`
	Data   Response `json:"data"`
}

//说明：status: 状态码，msg: 提示语，data: 具体的数据对象数组
func TableJson(status int, msg string, data string) string {
	return "{\"status\":" + strconv.Itoa(status) + ",\"msg\":\"" + msg + "\",\"data\":\"" + data + "\"}"
}

//写入日志到log文件
func WriteError(fileName string, content string) {
	if fileName == "" {
		fileName = "./runtime/logs/error.log"
	}
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0600)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
		if os.IsNotExist(err) {
			f, _ = os.Create(fileName)
		}
	} else {
		fd_time := time.Now().Format("2006-01-02 15:04:05")
		content := strings.Join([]string{"======", fd_time, "=====", content, "\n"}, "")
		_, err = f.Write([]byte(content))
		checkErr(err)
	}
}

//错误输出
func checkErr(err error) {
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
}

//缓存到redis，时长为1小时
func CacheRedis(key string, val string) {
	redis := caches.Cache
	_, err := redis.Set(key, val, time.Hour).Result()
	if err != nil {
		logs.Warn(helper.GetNowTimeStamp() + " - save redis: " + key + " fail!")
	}
}
//字符串数组去重
func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}
