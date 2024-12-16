package system

import (
	"crypto/rand"
	"fmt"
	"math/big"

	"pkg.world.dev/world-engine/cardinal"

	comp "argus-exercise-1-backend/component"
	"argus-exercise-1-backend/msg"
)

const (
	xLimit = 9
	yLimit = 9
	// player starts at bottom center
	playerInitialX = 4
	playerInitialY = 8
)

func secureIntn(xLimit int) (int, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(xLimit)))
	if err != nil {
		return 0, err
	}
	return int(n.Int64()), nil
}

func randomizeGoal() (int, int, error) {
	goalX, err := secureIntn(xLimit)
	if err != nil {
		return 0, 0, err
	}

	goalY, err := secureIntn(yLimit)
	if err != nil {
		return 0, 0, err
	}

	if goalX == playerInitialX && goalY == playerInitialY {
		return randomizeGoal() // Recursive call if the goal is same as player’s initial position
	}

	return goalX, goalY, nil
}

func RoomSpawnerSystem(world cardinal.WorldContext) error {
	return cardinal.EachMessage[msg.CreateRoomMsg, msg.CreateRoomResult](
		world,
		func(create cardinal.TxData[msg.CreateRoomMsg]) (msg.CreateRoomResult, error) {
			goalX, goalY, err := randomizeGoal()
			if err != nil {
				return msg.CreateRoomResult{}, fmt.Errorf("error randomizing goal: %w", err)
			}
			room := comp.Room{
				Owner:   create.Tx.PersonaTag,
				PlayerX: playerInitialX,
				PlayerY: playerInitialY,
				GoalX:   goalX,
				GoalY:   goalY,
			}
			id, err := cardinal.Create(world, room)
			if err != nil {
				return msg.CreateRoomResult{}, fmt.Errorf("error creating room: %w", err)
			}

			err = world.EmitEvent(map[string]any{
				"event": "new_room",
				"id":    id,
			})
			if err != nil {
				return msg.CreateRoomResult{}, err
			}
			return msg.CreateRoomResult{Success: true}, nil
		})
}
