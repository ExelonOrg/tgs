package structs

// TODO: add property for project name,
type Config struct {
	Stacks map[string]Stack `json:"stacks"`
}
type App struct {
	Dependencies []string `json:"dependencies"`
}

type AppType map[string]App

type Stack struct {
	Environments            []string
	StateStorageAccountName string
	StateStorageRG          string
}

type Group map[string][]string

type Pattern map[string]AppType

type Dependency struct {
	config_path string `hcl:"config_path,label"`
}
