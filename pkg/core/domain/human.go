package domain

type Human struct {
	Dna      []string `json:"dna"`
	IsMutant bool     `json:"-"`
}
