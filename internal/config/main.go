package config

import (
	"gitlab.com/distributed_lab/kit/comfig"
	"gitlab.com/distributed_lab/kit/copus"
	"gitlab.com/distributed_lab/kit/copus/types"
	"gitlab.com/distributed_lab/kit/kv"
	doormanCfg "github.com/dl-nft-books/doorman/connector/config"
)

type Config interface {
	// base
	comfig.Logger
	types.Copuser
	comfig.Listenerer

	// other configs
	MimeTypesConfigurator
	AWSConfigurator

	// connectors
	doormanCfg.DoormanConfiger
}

type config struct {
	// base
	comfig.Logger
	types.Copuser
	comfig.Listenerer

	// other configs
	MimeTypesConfigurator
	AWSConfigurator

	// connectors
	doormanCfg.DoormanConfiger

	getter kv.Getter
}

func New(getter kv.Getter) Config {
	return &config{
		getter:                getter,
		Copuser:               copus.NewCopuser(getter),
		Listenerer:            comfig.NewListenerer(getter),
		Logger:                comfig.NewLogger(getter, comfig.LoggerOpts{}),
		MimeTypesConfigurator: NewMimeTypesConfigurator(getter),
		AWSConfigurator:       NewAWSConfigurator(getter),
		DoormanConfiger:       doormanCfg.NewDoormanConfiger(getter),
	}
}
