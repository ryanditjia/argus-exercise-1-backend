package system

import (
	"pkg.world.dev/world-engine/cardinal"

	comp "argus-exercise-1-backend/component"
	"argus-exercise-1-backend/msg"
)

func MoverSystem(world cardinal.WorldContext) error {
	return cardinal.EachMessage[msg.MoveMsg, msg.MoveMsgReply](
		world,
		func(move cardinal.TxData[msg.MoveMsg]) (msg.MoveMsgReply, error) {
			roomID, room, err := queryTargetRoom(world, move.Tx.PersonaTag)
			if err != nil {
				return msg.MoveMsgReply{}, err
			}

			switch move.Msg.Direction {
			case msg.Up:
				if room.PlayerY > 0 {
					room.PlayerY--
				}
			case msg.Down:
				if room.PlayerY < yLimit-1 {
					room.PlayerY++
				}
			case msg.Left:
				if room.PlayerX > 0 {
					room.PlayerX--
				}
			case msg.Right:
				if room.PlayerX < xLimit-1 {
					room.PlayerX++
				}
			}

			err = cardinal.SetComponent[comp.Room](world, roomID, room)

			if err != nil {
				return msg.MoveMsgReply{}, err
			}

			return msg.MoveMsgReply{Success: true}, nil
		})
}
