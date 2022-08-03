package models

import "github.com/mSh4ke/authorization/api"

type Permission struct {
	Id       int
	Name     string
	ServerId int
}

func (perm *Permission) ConstructUrl(api *api.API) string {
	servers := *api.Config.Servers
	return servers[perm.ServerId] + perm.Name
}
