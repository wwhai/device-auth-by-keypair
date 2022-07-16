package auth

import (
	"fmt"
	"testing"
)

func Test_auth_device(t *testing.T) {
	password := GenPassword("c1", "user")
	t.Log(ValidatePassword("c1", "user", password))
}
func Test_gen_auth_kpair(t *testing.T) {
	fmt.Println("# 测试设备数据")
	fmt.Println("| ClientID | Username | Password |")
	fmt.Println("|---|---|---|")
	for i := 0; i < 30; i++ {
		clientid := fmt.Sprintf("c:%d", i)
		username := fmt.Sprintf("u:%d", i)
		password := GenPassword(clientid, username)
		fmt.Printf("|%s|%s|%s|\n", clientid, username, password)

	}

}
