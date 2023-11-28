package front_test

import (
	"net/url"
	"testing"
)

func TestRegister(t *testing.T) {
	Req("user/login", "GET", nil)
}

func TestGetUserFriends(t *testing.T) {
	params := url.Values{}
	params.Add("user_id", "1")
	Req("user/getUserFriends?"+params.Encode(), "GET", nil)
}
