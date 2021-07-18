package controller

type Param struct {
	ConfigFilePath string
	LogLevel       string
	StatePath      string
	DryRun         bool
	HCLFilePaths   []string
}
