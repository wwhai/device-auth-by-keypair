package auth

import (
	"testing"
)

func Test_auth_device(t *testing.T) {
	password := GenPassword("c1", "user")
	t.Log(ValidatePassword("c1", "user", password))
}
