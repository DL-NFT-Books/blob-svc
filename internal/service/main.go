package service

import (
	"net"
	"net/http"

	"gitlab.com/distributed_lab/logan/v3"

	"github.com/dl-nft-books/blob-svc/internal/config"

	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/logan/v3/errors"
)

type service struct {
	log       *logan.Entry
	copus     types.Copus
	listener  net.Listener
	mimeTypes *config.MimeTypes
	aws       *config.AWSConfig
}

func (s *service) run(cfg config.Config) error {
	s.log.Info("Service started")
	r := s.router(cfg)

	if err := s.copus.RegisterChi(r); err != nil {
		return errors.Wrap(err, "cop failed")
	}

	return http.Serve(s.listener, r)
}

func newService(cfg config.Config) *service {
	return &service{
		log:       cfg.Log(),
		copus:     cfg.Copus(),
		listener:  cfg.Listener(),
		mimeTypes: cfg.MimeTypes(),
		aws:       cfg.AWSConfig(),
	}
}

func Run(cfg config.Config) {
	if err := newService(cfg).run(cfg); err != nil {
		panic(err)
	}
}
