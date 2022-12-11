package server

import (
	"context"

	"github.com/microservices/incidentgeneration/config"
	"github.com/microservices/incidentgeneration/internal/service"
)

func (c *Container) InjectIncidentHandler(ctx context.Context, cfg *config.Config) service.IFindingsProcessorHandler {
	return nil
}
