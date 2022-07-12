package user

import (
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"helloKit/endpoint"
	user "helloKit/model"
	"helloKit/service"
	transport "helloKit/transport/http"
	"net/http"
)

func decodeCreateRequest(_ context.Context, r *http.Request) (interface{}, error) {
	if err := transport.FormCheckAccess(r); err != nil {
		return nil, err
	}
	r.ParseForm()
	req := &user.CreateReq{}
	err := transport.ParseForm(r.Form, req)
	if err != nil {
		return nil, err
	}
	r.Body.Close()
	return req, nil
}

func encodeCreateResponse(_ context.Context, w http.ResponseWriter, resp interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(resp)
}

func MakeCreateHandler(svc service.UserSvc) http.Handler {
	handler := httptransport.NewServer(
		endpoint.MakeCreateEndpoint(svc),
		decodeCreateRequest,
		encodeCreateResponse,
		transport.ErrorServerOption(), // 自定义错误处理
	)
	return handler
}
