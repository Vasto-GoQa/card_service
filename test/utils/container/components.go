package container

import (
	card_service "card_service/test/core/card_service/card"
	"card_service/test/utils/config"
	"card_service/test/utils/logger"

	"go.uber.org/dig"
)

type Components struct {
	CardService *card_service.Client

	Logger logger.Service
	Config config.Config
}

func initComponents(c *dig.Container) (*Components, error) {
	var err error
	components := Components{}

	err = c.Invoke(func(
		cardService *card_service.Client,

		logger logger.Service,
		conf config.Config,
	) {
		components.Config = conf
		components.Logger = logger

		components.CardService = cardService
	})

	if err != nil {
		return nil, err
	}

	return &components, nil
}
