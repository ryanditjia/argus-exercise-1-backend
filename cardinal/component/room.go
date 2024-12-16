package component

type Room struct {
	Owner   string `json:"owner"`
	PlayerX int    `json:"player_x"`
	PlayerY int    `json:"player_y"`
	GoalX   int    `json:"goal_x"`
	GoalY   int    `json:"goal_y"`
}

func (Room) Name() string {
	return "Room"
}
