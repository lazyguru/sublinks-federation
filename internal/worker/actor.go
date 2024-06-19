package worker

import (
	"encoding/json"
	"errors"
	"sublinks/sublinks-federation/internal/log"
	"sublinks/sublinks-federation/internal/model"
	"sublinks/sublinks-federation/internal/service"
)

type ActorWorker struct {
	log.Logger
	userService      *service.UserService
	communityService *service.CommunityService
}

func NewActorWorker(logger log.Logger, userService *service.UserService, communityService *service.CommunityService) *ActorWorker {
	return &ActorWorker{
		Logger:           logger,
		userService:      userService,
		communityService: communityService,
	}
}

func (w *ActorWorker) Process(msg []byte) error {
	actor := model.Actor{}
	err := json.Unmarshal(msg, &actor)
	if err != nil {
		w.Logger.Error("Error unmarshalling actor", err)
		return err
	}
	switch actor.ActorType {
	case "Person":
		person := model.Person{
			Id:           actor.Id,
			Username:     actor.Username,
			Name:         actor.Name,
			Bio:          actor.Bio,
			MatrixUserId: actor.MatrixUserId,
			PublicKey:    actor.PublicKey,
			PrivateKey:   actor.PrivateKey,
		}
		if !w.userService.Save(&person) {
			w.Logger.Error("Error saving person", nil)
			return errors.New("Error saving person")
		}
	case "Group":
		group := model.Group{
			Id:         actor.Id,
			Username:   actor.Username,
			Name:       actor.Name,
			Bio:        actor.Bio,
			PublicKey:  actor.PublicKey,
			PrivateKey: actor.PrivateKey,
		}
		if !w.communityService.Save(&group) {
			w.Logger.Error("Error saving community", nil)
			return errors.New("Error saving ommunity")
		}
	}
	return nil
}
