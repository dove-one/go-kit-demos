package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"helloKit/errors"
	user "helloKit/model"
	"helloKit/service"
)

func MakeCreateEndpoint(svc service.UserSvc) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*user.CreateReq)
		if !ok {
			return nil, errors.EndpointTypeError
		}
		resp, err := svc.Create(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}
}
