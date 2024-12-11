package system

import (
	"fmt"

	"pkg.world.dev/world-engine/cardinal"
	"pkg.world.dev/world-engine/cardinal/filter"
	"pkg.world.dev/world-engine/cardinal/types"

	comp "argus-exercise-1-backend/component"
)

// queryTargetRoom queries for the target room’s entity ID and room component.
func queryTargetRoom(world cardinal.WorldContext, targetOwner string) (types.EntityID, *comp.Room, error) {
	var roomID types.EntityID
	var room *comp.Room
	var err error
	searchErr := cardinal.NewSearch().Entity(
		filter.Exact(filter.Component[comp.Room]())).Each(world,
		func(id types.EntityID) bool {
			var currentRoom *comp.Room
			currentRoom, err = cardinal.GetComponent[comp.Room](world, id)
			if err != nil {
				return false
			}

			// Terminates the search if the room is found
			if currentRoom.Owner == targetOwner {
				roomID = id
				room = currentRoom
				return false
			}

			// Continue searching if the room is not the target room
			return true
		})
	if searchErr != nil {
		return 0, nil, err
	}
	if err != nil {
		return 0, nil, err
	}
	if room == nil {
		return 0, nil, fmt.Errorf("room %q does not exist", targetOwner)
	}

	return roomID, room, err
}
