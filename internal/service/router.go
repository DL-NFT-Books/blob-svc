package service

import (
	"github.com/go-chi/chi"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/tokend/nft-books/blob-svc/internal/service/middlewares"

	"gitlab.com/tokend/nft-books/blob-svc/internal/config"
	"gitlab.com/tokend/nft-books/blob-svc/internal/service/handlers"
	"gitlab.com/tokend/nft-books/blob-svc/internal/service/helpers"
)

func (s *service) router(cfg config.Config) chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			helpers.CtxLog(s.log),
			helpers.CtxMimeTypes(s.mimeTypes),
			helpers.CtxAwsConfig(s.aws),
			helpers.CtxDoormanConnector(cfg.DoormanConnector()),
		),
	)

	r.Route("/integrations", func(r chi.Router) {
		r.Route("/documents", func(r chi.Router) {
			r.With(middlewares.CheckAccessToken).Post("/", handlers.CreateDocument)
			r.Route("/{key}", func(r chi.Router) {
				r.Get("/", handlers.GetDocumentByKey)
				r.With(middlewares.CheckAccessToken).Delete("/", handlers.DeleteFile)
			})
		})
	})

	return r
}
