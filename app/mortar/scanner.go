package mortar

import (
	"com.facehunter/scanner/app/controllers"
	"com.facehunter/scanner/app/services"
	"com.facehunter/scanner/app/validations"
	api "com.facehunter/scanner/pkg/proto/api/v1"
	"context"
	serverInt "github.com/go-masonry/mortar/interfaces/http/server"
	"github.com/go-masonry/mortar/providers/groups"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type scannerServiceDeps struct {
	fx.In

	// API Implementations, "Register" them as GRPCServiceAPI
	ScannerService api.ScannerServiceServer
}

func ServiceAPIsAndOtherDependenciesFxOption() fx.Option {
	return fx.Options(
		// GRPC Service APIs registration
		fx.Provide(fx.Annotated{
			Group:  groups.GRPCServerAPIs,
			Target: serviceGRPCServiceAPIs,
		}),
		// GRPC Gateway Generated Handlers registration
		fx.Provide(fx.Annotated{
			Group:  groups.GRPCGatewayGeneratedHandlers + ",flatten", // "flatten" does this [][]serverInt.GRPCGatewayGeneratedHandlers -> []serverInt.GRPCGatewayGeneratedHandlers
			Target: serviceGRPCGatewayHandlers,
		}),
		// All other dependencies
		serviceDependencies(),
	)
}

func serviceGRPCServiceAPIs(deps scannerServiceDeps) serverInt.GRPCServerAPI {
	return func(srv *grpc.Server) {
		api.RegisterScannerServiceServer(srv, deps.ScannerService)
		// Any additional gRPC Implementations should be called here
	}
}

func serviceGRPCGatewayHandlers() []serverInt.GRPCGatewayGeneratedHandlers {
	return []serverInt.GRPCGatewayGeneratedHandlers{
		// Register service REST API
		func(mux *runtime.ServeMux, localhostEndpoint string) error {
			return api.RegisterScannerServiceHandlerFromEndpoint(context.Background(), mux, localhostEndpoint, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
		},
		// Any additional gRPC gateway registrations should be called here
	}
}

func serviceDependencies() fx.Option {
	return fx.Provide(
		services.CreateScannerServiceService,
		controllers.CreateScannerServiceController,
		validations.CreateScannerServiceValidations,
	)
}
