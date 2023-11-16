package cmd

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

const (
	TaklerHost        = "TAKLER_HOST"
	TaklerPort        = "TAKLER_PORT"
	TaklerName        = "TAKLER_NAME"
	TaklerConnectFile = "TAKLER_CONNECT_FILE"
	DefaultHost       = "localhost"
	DefaultPort       = "33083"
)

type Address struct {
	Hostname string `yaml:"hostname"`
	Ip       string `yaml:"ip"`
	Port     string `yaml:"port"`
}

type Server struct {
	Address Address `yaml:"address"`
}

type ConnectConfig struct {
	Server Server `yaml:"server"`
}

func loadConnectConfig(filePath string) (*ConnectConfig, error) {
	buf, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	c := &ConnectConfig{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %w", filePath, err)
	}

	return c, nil
}

func getHostAndPort(host string, port string) (string, string) {
	resultHost := DefaultHost
	resultPort := DefaultPort

	envHost := os.Getenv(TaklerHost)
	if len(envHost) > 0 {
		resultHost = envHost
	}
	envPort := os.Getenv(TaklerPort)
	if len(envPort) > 0 {
		resultPort = envPort
	}

	connectFilePath := os.Getenv(TaklerConnectFile)
	if len(connectFilePath) > 0 {
		connectConfig, err := loadConnectConfig(connectFilePath)
		if err != nil {
			log.Fatalf("load connect config filed: %s", connectFilePath)
		}
		resultHost = connectConfig.Server.Address.Hostname
		resultPort = connectConfig.Server.Address.Port
	}

	if len(host) > 0 {
		resultHost = host
	}
	if len(port) > 0 {
		resultPort = port
	}

	return resultHost, resultPort
}

// Deprecated: getHost is deprecated. Use getHostAndPort instead.
func getHost(host string) string {
	if len(host) > 0 {
		return host
	}
	host = os.Getenv(TaklerHost)
	if len(host) > 0 {
		return host
	}
	return DefaultHost
}

// Deprecated: getPort is deprecated. Use getHostAndPort instead.
func getPort(port string) string {
	if len(port) > 0 {
		return port
	}
	port = os.Getenv(TaklerPort)
	if len(port) > 0 {
		return port
	}
	return DefaultPort
}

func getNodePath(nodePath string) string {
	if len(nodePath) > 0 {
		return nodePath
	}
	nodePath = os.Getenv(TaklerName)
	if len(nodePath) > 0 {
		return nodePath
	}
	return ""
}
