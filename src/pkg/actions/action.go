package actions

import "context"

type Action interface {
	Execute(ctx context.Context) error
}
