package container

import (
	card_service "card_service/test/core/card_service/card"
	"card_service/test/utils/config"
	"card_service/test/utils/logger"

	"go.uber.org/dig"
)

func BuildContainer() (*Components, error) {
	c := dig.New()
	servicesConstructors := []interface{}{
		card_service.NewClient,
		config.NewConfig,
		logger.NewLoggerService,
	}

	for _, service := range servicesConstructors {
		err := c.Provide(service)
		if err != nil {
			return nil, err
		}
	}
	components, componentsError := initComponents(c)

	if componentsError != nil {
		return nil, componentsError
	}

	return components, nil
}
