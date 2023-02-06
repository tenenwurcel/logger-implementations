package main

import (
	"github.com/dock-tech/dock-gocore-lib/json"
	"github.com/dock-tech/dock-gocore-lib/log"
	"os"
)

type Fields struct {
	Message string `json:"message,omitempty"`
	Name    string `json:"name,omitempty"`
	Age     int    `json:"age,omitempty"`
	Field1  string `json:"field1,omitempty"`
	Field2  string `json:"field2,omitempty"`
}

func (lf Fields) JSON() map[string]any {
	jsonMap := make(map[string]any)
	b := json.Marshal(lf)
	json.Unmarshal(b, &jsonMap)

	return jsonMap
}

func main() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
	log.InitializeLogs(log.Fields{"app_name": "dummy", "app_version": "0.0.1"})

	log.Debugj(Fields{
		Message: "Hello Worlds",
		Name:    "John Doe",
		Age:     42,
	})
}
