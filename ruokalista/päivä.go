package ruokalista

// Päivä kertoo yhden päivän ruuat
type Päivä struct {
	Viikonpäivä string `json:"paiva"`
	Perus       string `json:"kotiruoka"`
	Veg         string `json:"kasvisruoka"`
}

// Viikko kertoo yhden viikon ruuat
type Viikko []Päivä
