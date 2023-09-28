package cmd

import "os"

const (
	TaklerHost  = "TAKLER_HOST"
	TaklerPort  = "TAKLER_PORT"
	TaklerName  = "TAKLER_NAME"
	DefaultHost = "localhost"
	DefaultPort = "33083"
)

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
