package front_test

import (
	"testing"
)

func TestRegister(t *testing.T) {
	Req("user/login", "GET", nil)
}
