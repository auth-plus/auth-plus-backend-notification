package config

import "os"

type FirebaseEnv struct {
	CredentialPath string
	AppName        string
}

type ProviderEnv struct {
	Firebase FirebaseEnv
}

type AppEnv struct {
	Name string
	Port string
	Env  string
}

type Environment struct {
	App       AppEnv
	Providers ProviderEnv
}

func GetEnv() Environment {
	app := AppEnv{
		Name: os.Getenv("APP_NAME"),
		Port: os.Getenv("APP_PORT"),
		Env:  os.Getenv("GO_ENV"),
	}
	firebase := FirebaseEnv{
		CredentialPath: os.Getenv("GOOGLE_APPLICATION_CREDENTIALS"),
		AppName:        os.Getenv("GOOGLE_APPLICATION_NAME"),
	}
	providers := ProviderEnv{
		Firebase: firebase,
	}
	env := Environment{
		App:       app,
		Providers: providers,
	}
	return env
}
