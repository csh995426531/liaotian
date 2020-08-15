package handler

import (
	"liaotian/basic/config"
	"net/http"
	"testing"
)

func TestMain(m *testing.M) {
	config.Init(func(options *config.Options) {
		options.Path = "../"
	})

	m.Run()
}

func TestLogin(t *testing.T) {

	loginTests := []struct{
		Name		string
		Password 	string

	} {
		{"张三", "aa123123123"},
		{"张四", "aa123123123"},
	}

	for _, test := range loginTests {
		t.Run(test.Name, func(t *testing.T) {
			resp := http.Response{}
			Login(resp, &http.Request{

			})
		})
	}
}
