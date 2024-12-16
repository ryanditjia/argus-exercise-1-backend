package main

import (
	"errors"

	"github.com/rs/zerolog/log"
	"pkg.world.dev/world-engine/cardinal"

	"argus-exercise-1-backend/component"
	"argus-exercise-1-backend/msg"
	"argus-exercise-1-backend/query"
	"argus-exercise-1-backend/system"
)

func main() {
	w, err := cardinal.NewWorld(cardinal.WithDisableSignatureVerification())
	if err != nil {
		log.Fatal().Err(err).Msg("")
	}

	MustInitWorld(w)

	Must(w.StartGame())
}

// MustInitWorld registers all components, messages, queries, and systems. This initialization happens in a helper
// function so that this can be used directly in tests.
func MustInitWorld(w *cardinal.World) {
	// Register components
	// NOTE: You must register your components here for it to be accessible.
	Must(
		cardinal.RegisterComponent[component.Room](w),
	)

	// Register messages (user action)
	// NOTE: You must register your transactions here for it to be executed.
	Must(
		cardinal.RegisterMessage[msg.CreateRoomMsg, msg.CreateRoomResult](w, "create-room"),
		cardinal.RegisterMessage[msg.DeleteRoomMsg, msg.DeleteRoomResult](w, "delete-room"),
		cardinal.RegisterMessage[msg.MoveMsg, msg.MoveMsgReply](w, "move"),
	)

	// Register queries
	// NOTE: You must register your queries here for it to be accessible.
	Must(
		cardinal.RegisterQuery[query.RoomStateRequest, query.RoomStateResponse](w, "room-state", query.RoomState),
	)

	// Each system executes deterministically in the order they are added.
	// This is a neat feature that can be strategically used for systems that depends on the order of execution.
	// For example, you may want to run the attack system before the regen system
	// so that the player's HP is subtracted (and player killed if it reaches 0) before HP is regenerated.
	Must(cardinal.RegisterSystems(w,
		system.RoomSpawnerSystem,
		system.RoomDeleterSystem,
		system.MoverSystem,
	))

	// Must(cardinal.RegisterInitSystems(w))
}

func Must(err ...error) {
	e := errors.Join(err...)
	if e != nil {
		log.Fatal().Err(e).Msg("")
	}
}
