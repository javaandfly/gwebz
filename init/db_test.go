package init

import (
	"testing"
)

func TestInitServerDB(t *testing.T) {
	type UserInfo struct {
		Username string `gorm:"username" json:"username"`
		Password string `gorm:"password" json:"password"`
	}
	userTest := UserInfo{
		Username: "xiaodong",
		Password: "123456",
	}
	err := InitGlobalDB("mysql", "127.0.0.1", "3306", "test", "root", "123456", "utf8")
	if err != nil {
		panic(err)
	}
	err = GetDB().AutoMigrate(&UserInfo{})
	if err != nil {
		panic(err)
	}
	err = GetDB().Create(&UserInfo{
		Username: "xiaodong",
		Password: "123456",
	}).Error
	if err != nil {
		panic(err)
	}
	userInfo := UserInfo{}
	err = GetDB().Where("username =?", "xiaodong").First(&userInfo).Error
	if err != nil {
		panic(err)
	}

	if userInfo != userTest {
		t.Errorf("TestInitServerDB() = %v, tests %v", userInfo, userTest)
	}
}
