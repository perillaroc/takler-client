package cmd

import "os"

const (
	TaklerHost = "TAKLER_HOST"
	TaklerPort = "TAKLER_PORT"
	TaklerName = "TAKLER_NAME"
)

func getHost(host string) string {
	if len(host) > 0 {
		return host
	}
	host = os.Getenv(TaklerHost)
	if len(host) > 0 {
		return host
	}
	return ""
}

func getPort(port string) string {
	if len(port) > 0 {
		return port
	}
	port = os.Getenv(TaklerPort)
	if len(port) > 0 {
		return port
	}
	return ""
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
