package log

import (
	"errors"

	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/javaandfly/gwebz/utils"
)

func TestInitLog(t *testing.T) {
	serverName := strings.Split(filepath.Base(os.Args[0]), ".")[0]
	serverMark := utils.GetSvrmark("sync" + serverName)

	err := InitLog("test_log/", serverName, serverMark, func(str string) {})
	if err != nil {
		panic(err)
	}
	LogSetCallback(func(s string) {
		LogW("回调函数被调用了")
	})
	LogW("1111")
	LogD("1111")
	LogError(errors.New("12121"))

}
