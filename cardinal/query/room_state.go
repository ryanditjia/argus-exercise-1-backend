package query

import (
	"fmt"

	"pkg.world.dev/world-engine/cardinal/filter"
	"pkg.world.dev/world-engine/cardinal/types"

	comp "argus-exercise-1-backend/component"

	"pkg.world.dev/world-engine/cardinal"
)

type RoomStateRequest struct {
	Owner string
}

type RoomStateResponse struct {
	PlayerX int `json:"player_x"`
	PlayerY int `json:"player_y"`
	GoalX   int `json:"goal_x"`
	GoalY   int `json:"goal_y"`
}

func RoomState(world cardinal.WorldContext, req *RoomStateRequest) (*RoomStateResponse, error) {
	var roomState *comp.Room
	var err error
	searchErr := cardinal.NewSearch().Entity(
		filter.Exact(filter.Component[comp.Room]())).
		Each(world, func(id types.EntityID) bool {
			var room *comp.Room
			room, err = cardinal.GetComponent[comp.Room](world, id)
			if err != nil {
				return false
			}

			// Terminates the search if the room is found
			if room.Owner == req.Owner {
				roomState, err = cardinal.GetComponent[comp.Room](world, id)
				if err != nil {
					return false
				}
				return false
			}

			// Continue searching if the room is not the target room
			return true
		})
	if searchErr != nil {
		return nil, searchErr
	}
	if err != nil {
		return nil, err
	}

	if roomState == nil {
		return nil, fmt.Errorf("room %s does not exist", req.Owner)
	}

	return &RoomStateResponse{
		PlayerX: roomState.PlayerX,
		PlayerY: roomState.PlayerY,
		GoalX:   roomState.GoalX,
		GoalY:   roomState.GoalY,
	}, nil
}
