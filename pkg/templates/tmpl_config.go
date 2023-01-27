package templates

func GenerateConfig(_ string) []byte {
	content := `package config

type Config struct {
	// ExampleOpt is an example config option.
	ExampleOpt string
}
`
	return []byte(content)
}
