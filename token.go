package main

import (
	"fmt"
	"io/ioutil"
)

func Token(mode string) string {
	token, err := ioutil.ReadFile("./plutobot-token.tk")
	if err != nil {
		token, err = ioutil.ReadFile("~/.config/plutobot-token.tk")
		if err != nil {
			fmt.Println(err)
			return ""
		}
	}
	return string(token)
}
