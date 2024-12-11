package system

import (
	"pkg.world.dev/world-engine/cardinal"

	"argus-exercise-1-backend/msg"
)

func RoomDeleterSystem(world cardinal.WorldContext) error {
	return cardinal.EachMessage[msg.DeleteRoomMsg, msg.DeleteRoomResult](
		world,
		func(del cardinal.TxData[msg.DeleteRoomMsg]) (msg.DeleteRoomResult, error) {
			roomID, _, err := queryTargetRoom(world, del.Msg.Owner)
			if err != nil {
				return msg.DeleteRoomResult{}, err
			}

			err = cardinal.Remove(world, roomID)
			if err != nil {
				return msg.DeleteRoomResult{}, err
			}

			return msg.DeleteRoomResult{Success: true}, nil
		})
}
