package openapi

import (
	"com.facehunter/scanner/app/conf"
	"com.facehunter/scanner/third_party"
	"github.com/go-masonry/mortar/interfaces/cfg"
	"github.com/go-masonry/mortar/interfaces/log"
	"github.com/go-masonry/mortar/providers/types"
	"go.uber.org/fx"
	"io/fs"
	"mime"
	"net/http"
)

type DocumentationHandler struct {
	logger  log.Logger
	cfg     cfg.Config
	docPath string
}

type documentationHandlerDeps struct {
	fx.In
	Logger log.Logger
	Cfg    cfg.Config
}

func CreateDocumentationHandler(deps documentationHandlerDeps) *DocumentationHandler {
	return &DocumentationHandler{
		logger:  deps.Logger,
		cfg:     deps.Cfg,
		docPath: deps.Cfg.Get(conf.SwaggerRestPath).String(),
	}
}

func (h *DocumentationHandler) getOpenAPIHandler() http.Handler {
	err := mime.AddExtensionType(".svg", "image/svg+xml")
	if err != nil {
		panic("couldn't add mime: " + err.Error())
	}
	// Use subdirectory in embedded files
	subFS, err := fs.Sub(third_party.OpenAPI, "OpenAPI")
	if err != nil {
		panic("couldn't create sub filesystem: " + err.Error())
	}

	return http.StripPrefix(h.docPath, http.FileServer(http.FS(subFS)))
}

func (h *DocumentationHandler) handleRequest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.getOpenAPIHandler().ServeHTTP(w, r)
	}
}

func HandlerEndpoint(h *DocumentationHandler) types.HTTPHandlerFuncPatternPair {
	handlerFunc := h.handleRequest()

	return types.HTTPHandlerFuncPatternPair{
		Pattern:     h.docPath,
		HandlerFunc: handlerFunc,
	}
}
