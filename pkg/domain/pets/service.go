package pets

import (
	"context"

	api "github.com/abhisek/go-microservice-starter/gen"
	log "github.com/sirupsen/logrus"
)

type petServer struct {
	api.PetServiceServer
}

func NewPetService() (*petServer, error) {
	log.Debug("Initializing new pet service")
	return &petServer{}, nil
}

func (server *petServer) CreatePet(context context.Context, info *api.PetInfo) (*api.Pet, error) {
	log.Debug("Creating pet..")
	return &api.Pet{}, nil
}
