package main

import (
	"io/ioutil"
	"fmt"
)

func Token(mode string) string {
	token, err := ioutil.ReadFile("~/.config/plutobot-token.tk")
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(token)
}
