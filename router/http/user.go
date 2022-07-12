package http

import (
	server "helloKit/server/http"
	svc "helloKit/service"
	transport "helloKit/transport/http/user"
)

func init() {
	registerUserHandler()
}

func registerUserHandler() {
	server.RegisterHandler("/user/create", transport.MakeCreateHandler(svc.NewUserSvc()))
	server.RegisterHandler("/user/delete", transport.MakeDeleteHandler(svc.NewUserSvc()))
}
