package structs

type Config struct {
	BaseModules []string         `mapstructure:"base_modules"`
	Groups      map[string]Group `mapstructure:"groups"`
}
type App struct {
	Dependencies []string `mapstructure:"dependencies"`
}

type AppType map[string]App

type Stack map[string]AppType

type Group map[string]Stack

type base_modules []string

func Testing() {

	// if err != nil {
	// 	os.Exit(1)
	// }
}
