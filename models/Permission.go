package models

import (
	"fmt"
	"regexp"
	"strings"
)

type Permission struct {
	Id       int    `json:"id"`
	Path     string `json:"path"`
	Method   string `json:"method"`
	ServerId int    `json:"server_id"`
}

func (perm *Permission) ConstructUrl(server string) string {
	return server + perm.Path
}

func (perm *Permission) ParseUrl() string {
	fmt.Println("parsing request url")
	fmt.Println(perm.Path)
	if res, err := regexp.MatchString("^/..*/[0-9][0-9]*$", perm.Path); err != nil {
		fmt.Println(err)
		return ""
	} else if res {
		return strings.TrimRight(perm.Path, "1234567890") + "param"
	}
	return perm.Path
}
