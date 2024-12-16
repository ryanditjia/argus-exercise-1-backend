package msg

type Direction string

const (
	Up    Direction = "up"
	Down  Direction = "down"
	Left  Direction = "left"
	Right Direction = "right"
)

type MoveMsg struct {
	Direction Direction
}

type MoveMsgReply struct {
	Success bool `json:"success"`
}
