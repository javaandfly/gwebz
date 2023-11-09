package utils

import (
	"fmt"
	"os"
	"reflect"
	"strings"
)

// InArray :给定元素值 是否在 指定的数组中
func InArray(needle interface{}, hystack interface{}) bool {
	if harr, ok := ToSlice(hystack); ok {
		for _, item := range harr {
			if item == needle {
				return true
			}
		}
	}
	return false
}

// 通用转换为数组
func ToSlice(arr interface{}) ([]interface{}, bool) {
	v := reflect.ValueOf(arr)
	if v.Kind() != reflect.Slice {
		return nil, false
	}
	l := v.Len()
	ret := make([]interface{}, l)
	for i := 0; i < l; i++ {
		ret[i] = v.Index(i).Interface()
	}
	return ret, true
}

func GetSvrmark(svrname string, serverid ...string) string {
	hostname, _ := os.Hostname()
	if pidx := strings.Index(string(hostname), "."); pidx > 0 {
		hostname = string([]byte(hostname)[:pidx-1])
	}
	if len(serverid) > 0 && len(serverid[0]) > 0 {
		return fmt.Sprintf("%s-%s", svrname, serverid[0])
	}
	pid := os.Getpid()
	return fmt.Sprintf("%s-%s-%d", hostname, svrname, pid)
}
