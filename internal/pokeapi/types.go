package pokeapi

type LocationArea struct {
	Name string `json:"name"`
}

type RespLocationAreas struct {
	Next     *string        `json:"next"`
	Previous *string        `json:"previous"`
	Results  []LocationArea `json:"results"`
}
