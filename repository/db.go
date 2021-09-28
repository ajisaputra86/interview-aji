package repository

import (
	"context"
	"io"

	"github.com/fajrirahmat/interview-aji/model"
)

type DB interface {
	io.Closer
	ListLocation(ctx context.Context) ([]*model.Location, error)
	AddCheckin(ctx context.Context, request *model.CheckInOutRequest) ([]*model.CheckInOutRequest, error)
}
