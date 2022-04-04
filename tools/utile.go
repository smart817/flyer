package tools

import (
	"regexp"

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

//邮箱格式验证
func VEmail(email string) bool {
	//pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	pattern := `^[0-9a-z][_.0-9a-z-]{0,31}@([0-9a-z][0-9a-z-]{0,30}[0-9a-z]\.){1,4}[a-z]{2,4}$`

	reg := regexp.MustCompile(pattern)
	return reg.MatchString(email)
}

//手机号码验证
func VPhnoe(mobileNum string) bool {
	regular := "^((13[0-9])|(14[5,7])|(15[0-3,5-9])|(17[0,3,5-8])|(18[0-9])|166|198|199|(147))\\d{8}$"

	reg := regexp.MustCompile(regular)
	return reg.MatchString(mobileNum)
}
