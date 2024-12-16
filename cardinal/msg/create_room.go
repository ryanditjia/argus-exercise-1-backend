package msg

type CreateRoomMsg struct{}

type CreateRoomResult struct {
	Success bool `json:"success"`
}
