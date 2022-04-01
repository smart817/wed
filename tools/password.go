package tools

import (
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

//生成一个新的uuid
func NewUUID() uuid.UUID {
	return uuid.New()
}

//密码加密
func EncodePassword(Password string) string {
	hash, _ := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost) //加密处理
	// 保存在数据库的密码，虽然每次生成都不一样，只需保存一份便可 长度60
	return string(hash)
}

//密码验证
func DecodePassword(hashPassword, Password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPassword), []byte(Password)) //验证（对比）
	if err != nil {
		return false
	} else {
		return true
	}
}
