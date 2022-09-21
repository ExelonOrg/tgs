package structs

type Config struct {
	Stacks map[string][]string `json:"stacks"`
}
type App struct {
	Dependencies []string `json:"dependencies"`
}

type AppType map[string]App

type Stack []string

type Group map[string][]string

type base_modules []string

type Pattern map[string]AppType
