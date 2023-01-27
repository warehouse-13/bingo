package templates

const (
	pkg = "pkg/"

	flagDir    = pkg + "flags/"
	cmdDir     = pkg + "command/"
	versionDir = pkg + "version/"
	cfgDir     = pkg + "config/"

	mainPkg    = "main.go"
	appCmd     = cmdDir + "app.go"
	exampleCmd = cmdDir + "example.go"
	versionCmd = cmdDir + "version.go"
	flagsPkg   = flagDir + "flags.go"
	cfgPkg     = cfgDir + "config.go"
	versionPkg = versionDir + "version.go"
)

type File struct {
	Name     string
	Contents func(string) []byte
}

func ManagedDirs() []string {
	return []string{cmdDir, flagDir, cfgDir, versionDir}
}

func ManagedFiles() []File {
	return []File{
		{
			Name:     mainPkg,
			Contents: GenerateMain,
		},
		{
			Name:     appCmd,
			Contents: GenerateApp,
		},
		{
			Name:     exampleCmd,
			Contents: GenerateExample,
		},
		{
			Name:     versionCmd,
			Contents: GenerateVersionCmd,
		},
		{
			Name:     flagsPkg,
			Contents: GenerateFlags,
		},
		{
			Name:     cfgPkg,
			Contents: GenerateConfig,
		},
		{
			Name:     versionPkg,
			Contents: GenerateVersionPkg,
		},
	}
}
