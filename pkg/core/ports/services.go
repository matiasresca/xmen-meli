package ports

type HumanService interface {
	IsMutant(dna []string) bool
}
