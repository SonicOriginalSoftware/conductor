package lib

import "os"

const defaultPort = "8080"

func lookupPort() (port string) {
	found := false
	if port, found = os.LookupEnv("PORT"); !found {
		port = defaultPort
	}
	return
}
