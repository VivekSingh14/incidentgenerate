package app

import (
	"context"
	"fmt"

	"github.com/microservices/incidentgeneration/config"
	"github.com/microservices/incidentgeneration/internal/server"
)

func Start() {
	fmt.Println("Starting the incident generator service !!")
	appconfig, configerr := config.Get()

	if configerr != nil {
		fmt.Println("error occured while loading config")
	} else {
		fmt.Println("config loaded succesfully...")
	}
	ctx := context.Background()

	initializeMetric(appconfig)
	registerIncidentProcessor(ctx, appconfig)

}

func initializeMetric(appconfig *config.Config) {
	fmt.Println("metrics initialized successfully....")
}

func registerIncidentProcessor(ctx context.Context, appconfig *config.Config) {
	c := server.ServiceContainer()
	service := c.InjectIncidentHandler(ctx, appconfig)
	fmt.Println(service)

}
