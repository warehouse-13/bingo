package templates

import "fmt"

func GenerateVersionPkg(name string) []byte {
	content := `package version

// PackageName is the name of the package, all commands fall under this name.
const PackageName = "%s"

var (
	Version    = "undefined" // Specifies the cli version
	BuildDate  = "undefined" // Specifies the build date
	CommitHash = "undefined" // Specifies the git commit hash
)
`
	return []byte(fmt.Sprintf(content, name))
}
