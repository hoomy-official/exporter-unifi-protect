package v1

import (
	"github.com/hoomy-official/go-shared/pkg/net/rest"
	"github.com/hoomy-official/go-shared/pkg/net/ws"
)

type V1 struct {
	Auth    *APIAuth
	Users   *APIUser
	Sensors *APISensor
	Live    *APILive
}

func NewV1(cl rest.Requester, w ws.Dialer) *V1 {
	return &V1{
		Auth:    NewAPIAuth(cl),
		Users:   NewAPIUsers(cl),
		Sensors: NewAPISensors(cl),
		Live:    NewAPILive(w),
	}
}
