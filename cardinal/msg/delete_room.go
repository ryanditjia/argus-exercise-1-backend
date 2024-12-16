package msg

type DeleteRoomMsg struct {
	Owner string
}

type DeleteRoomResult struct {
	Success bool `json:"success"`
}
