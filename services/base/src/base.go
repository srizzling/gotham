package base

import (
	"golang.org/x/net/context"

	proto "github.com/srizzling/gotham/services/base/proto"
)

// Service Interface for all Services to be running an action
type Service interface {
	HealthCheck(context context.Context, req proto.HealthCheckRequest, res proto.HealthCheckResponse) error
	RunAction(context context.Context, req proto.ActionRequest, res proto.ActionResponse) error
}
