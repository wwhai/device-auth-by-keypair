# 设备验证

这是个非常简单的设备秘钥生成和验证器。
```sh
func Test_auth_device(t *testing.T) {
	password := GenPassword("c1", "user")
	t.Log(ValidatePassword("c1", "user", password))
}

```