package auth

import (
	"crypto/md5"
	"fmt"
)

// 一个随机密钥
var _sk string = "it's a secret"

/*
*
* 生成密码
*
 */
func GenPassword(clientid, username string) string {
	return EnMd5(clientid+":"+username, _sk)
}

/*
*
* 验证密码
*
 */
func ValidatePassword(c, u, p string) bool {
	return EnMd5(c+":"+u, _sk) == p
}
func EnMd5(source, secret string) string {
	data := []byte(source + ":" + secret)
	hash := md5.Sum(data)
	return fmt.Sprintf("%x", hash)
}
