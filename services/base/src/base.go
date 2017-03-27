package base

import (
	"context"

	proto "github.com/srizzling/gotham/services/base/proto"
)

// Service Interface for all Services to be running an action
type Service interface {
	RunAction(context context.Context, req proto.ActionRequest, res proto.ActionResponse) error
}
