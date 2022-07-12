package endpoint

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"helloKit/errors"
	user "helloKit/model"
	"helloKit/service"
)

func MakeDeleteEndpoint(svc service.UserSvc) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req, ok := request.(*user.DeleteReq)
		if !ok {
			return nil, errors.EndpointTypeError
		}
		resp, err := svc.Delete(ctx, req)
		if err != nil {
			return nil, err
		}
		return resp, nil
	}

}
