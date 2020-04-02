package commands

type MovieCommand struct {
	Movie struct {
		Title    string `json:"title"`
		Duration int64  `json:"duration"`
		Synopsis string `json:"synopsis"`
	} `json:"movie"`
	Dni int `json:"dni"`
}
