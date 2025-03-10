package constants

const (
	// App
	AppName            = "Tulip"
	AppVersion         = "1.0.0"
	AppRootDir         = ".tulip"
	AppConfigDir       = "config"
	AppConfigFile      = "config.yml"

	// Proxy
    ProxyNetworkName   = "tulip"
    ProxyContainerName = "tulip_proxy"
    ProxyUrl           = "http://localhost:8855"
	ProxyConfigDir     = "proxy"
	ProxyDockerFile    = "docker-compose.yml"
	ProxyTraefikFile   = "traefik.yml"
)
