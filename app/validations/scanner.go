package validations

import (
	"context"

	api "github.com/aku-ato/scanner/pkg/proto/api/v1"
	"github.com/go-masonry/mortar/interfaces/auth/jwt"
	"github.com/go-masonry/mortar/interfaces/log"
	"go.uber.org/fx"
)

type ScannerServiceValidations interface {
	api.ScannerServiceServer
}

type ScannerServiceValidationsDeps struct {
	fx.In

	jwt.TokenExtractor
	Logger log.Logger
}

type ScannerServiceValidationsImpl struct {
	*api.UnimplementedScannerServiceServer // if keep this one added even when you change your interface this code will compile
	deps                                   ScannerServiceValidationsDeps
}

func CreateScannerServiceValidations(deps ScannerServiceValidationsDeps) ScannerServiceValidations {
	return &ScannerServiceValidationsImpl{
		deps: deps,
	}
}

func (impl *ScannerServiceValidationsImpl) Start(ctx context.Context, req *api.StartRequest) (*api.StartResponse, error) {
	impl.deps.Logger.Debug(ctx, "start called")
	return &api.StartResponse{}, nil
}
