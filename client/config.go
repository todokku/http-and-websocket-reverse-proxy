package main

import (
	"encoding/json"
	"os"
)

// Config for client
type Config struct {
	Target    string `json:"target"`
	RemoteAddr   string `json:"remoteaddr"`
	Key          string `json:"key"`
	SockBuf      int    `json:"sockbuf"`
	KeepAlive    int    `json:"keepalive"`
}

func parseJSONConfig(config *Config, path string) error {
	file, err := os.Open(path) // For read access.
	if err != nil {
		return err
	}
	defer file.Close()

	return json.NewDecoder(file).Decode(config)
}
