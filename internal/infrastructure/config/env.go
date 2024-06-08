package config

import (
	"github.com/spf13/viper"
)

type EnvConfigs struct {
	// Server
	Route Route

	SecretKey string `mapstructure:"SECRET_KEY"`

	// FDB
	FlezzDBUsername string `mapstructure:"FLEZZ_DB_USERNAME"`
	FlezzDBPassword string `mapstructure:"FLEZZ_DB_PASSWORD"`
	FlezzDBHost     string `mapstructure:"FLEZZ_DB_HOST"`
	FlezzDBPort     string `mapstructure:"FLEZZ_DB_PORT"`
	FlezzDBDatabase string `mapstructure:"FLEZZ_DB_DATABASE"`
}

// var EnvConfigs *EnvConfig
func LoadEnvVariables(serviceName string) (*EnvConfigs, error) {
	envConfigs := &EnvConfigs{}
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(envConfigs)
	if err != nil {
		return nil, err
	}

	route, err := loadRoute(serviceName)
	if err != nil {
		return nil, err
	}
	envConfigs.Route = route
	return envConfigs, nil
}

func loadRoute(serviceName string) (Route, error) {
	viper.SetConfigType("yaml")
	viper.SetConfigFile("gateway.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return Route{}, err
	}

	routes := []Route{}
	
	if viper.UnmarshalKey("gateway.routes", &routes) != nil {
		return Route{}, err
	}

	for _, route := range routes {
		if route.Name == serviceName {
			return route, nil
		}
	}
	return Route{}, nil
}

type Route struct {
	Name       string `mapstructure:"name"`
	Host       string `mapstructure:"host"`
	Protocol   string `mapstructure:"protocol"`
	Context    string `mapstructure:"context"`
	TargetPort string `mapstructure:"targetPort"`
}

type GatewayConfigs struct {
	// Server
	Host        string  `mapstructure:"gatewayHost"`
	GatewayPort string  `mapstructure:"gatewayPort"`
	Routes      []Route `mapstructure:"routes"`
}

func LoadGatewayConfigs() (*GatewayConfigs, error) {
	gatewayConfigs := &GatewayConfigs{}
	viper.SetConfigType("yaml")
	viper.SetConfigFile("gateway.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.UnmarshalKey("gateway", gatewayConfigs)
	if err != nil {
		return nil, err
	}

	return gatewayConfigs, nil
}
