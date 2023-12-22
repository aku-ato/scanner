package services

import (
	"com.facehunter/scanner/app/controllers"
	"com.facehunter/scanner/app/validations"
	"context"

	"github.com/go-masonry/mortar/interfaces/monitor"

	api "com.facehunter/scanner/pkg/proto/api/v1"

	"github.com/go-masonry/mortar/interfaces/log"
	"go.uber.org/fx"
)

type scannerServiceServiceDeps struct {
	fx.In

	Logger log.Logger

	Validations validations.ScannerServiceValidations
	Controller  controllers.ScannerServiceController
	Metrics     monitor.Metrics `optional:"true"`
}

type scannerServiceServiceImpl struct {
	api.UnimplementedScannerServiceServer // if keep this one added even when you change your interface this code will compile
	deps                                  scannerServiceServiceDeps
}

func CreateScannerServiceService(deps scannerServiceServiceDeps) api.ScannerServiceServer {
	return &scannerServiceServiceImpl{
		deps: deps,
	}
}

func (impl *scannerServiceServiceImpl) Start(ctx context.Context, req *api.StartRequest) (*api.StartResponse, error) {
	_, err := impl.deps.Validations.Start(ctx, req)
	if err != nil {
		impl.deps.Logger.WithError(err).Debug(ctx, "validation failed")
		return nil, err
	}
	return impl.deps.Controller.Start(ctx, req)
}
