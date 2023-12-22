package controllers

import (
	"context"

	api "com.facehunter/scanner/pkg/proto/api/v1"

	"github.com/go-masonry/mortar/interfaces/log"
	"go.uber.org/fx"
)

type ScannerServiceController interface {
	api.ScannerServiceServer
}

type ScannerServiceControllerDeps struct {
	fx.In

	Logger log.Logger
}

type ScannerServiceControllerImpl struct {
	*api.UnimplementedScannerServiceServer // if keep this one added even when you change your interface this code will compile
	deps                                   ScannerServiceControllerDeps
}

func CreateScannerServiceController(deps ScannerServiceControllerDeps) ScannerServiceController {
	return &ScannerServiceControllerImpl{
		deps: deps,
	}
}

func (w *ScannerServiceControllerImpl) Start(ctx context.Context, req *api.StartRequest) (*api.StartResponse, error) {
	w.deps.Logger.Debug(ctx, "start called")
	return &api.StartResponse{
		Message: "start called",
	}, nil
}
