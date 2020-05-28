package commands

type MovieCommand struct {
	Movie struct {
		Title    string `json:"Title"`
		Duration int64  `json:"Duration"`
		UrlImg   string `json:"UrlImg"`
		Synopsis string `json:"Synopsis"`
	} `json:"Movie"`
	Dni int `json:"Dni"`
}
