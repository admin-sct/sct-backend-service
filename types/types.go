package types

import (
	"context"
	"sct-backend-service/graph/model"
)

type GraphQLService interface {
	SendContactInfo(ctx context.Context, input model.SendContactInfoRequest) (*model.SendContactInfoResponse, error)
}
