package customhandlers

import (
	"github.com/aku-ato/scanner/app/customhandlers/openapi"
	"github.com/go-masonry/mortar/providers/groups"
	"go.uber.org/fx"
)

func FxOption() fx.Option {
	return fx.Provide(
		fx.Annotated{
			Group:  groups.ExternalHTTPHandlerFunctions,
			Target: openapi.HandlerEndpoint,
		},
		// Register Here other custom http endpoints
	)
}
