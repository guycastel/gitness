// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package server

import (
	"github.com/bradrydzewski/my-app/internal/cron"
	"github.com/bradrydzewski/my-app/internal/router"
	"github.com/bradrydzewski/my-app/internal/server"
	"github.com/bradrydzewski/my-app/internal/store/database"
	"github.com/bradrydzewski/my-app/internal/store/memory"
	"github.com/bradrydzewski/my-app/types"
)

// Injectors from wire.go:

func initSystem(config *types.Config) (*system, error) {
	db, err := database.ProvideDatabase(config)
	if err != nil {
		return nil, err
	}
	executionStore := database.ProvideExecutionStore(db)
	pipelineStore := database.ProvidePipelineStore(db)
	userStore := database.ProvideUserStore(db)
	systemStore := memory.New(config)
	handler := router.New(executionStore, pipelineStore, userStore, systemStore)
	serverServer := server.ProvideServer(config, handler)
	nightly := cron.NewNightly()
	serverSystem := newSystem(serverServer, nightly)
	return serverSystem, nil
}
