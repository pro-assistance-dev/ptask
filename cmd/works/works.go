package works

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type Work struct {
	Name   string `json:"name"`
	Script string `json:"script"`
}

type Works []Work

func (items Works) GetNames() []string {
	names := make([]string, len(items))
	for i, item := range items {
		names[i] = item.GetName()
	}
	return names
}

func (item Work) GetName() string {
	return item.Name
}

func (item Work) Open(scriptsPath string) {
	script := filepath.Join(scriptsPath, item.Script)

	cmd := exec.Command("/bin/sh", "-c", script)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}

var works Works

func GetWorks() Works {
	return works
}

func ParseWorks(scriptsPath string) {
	scriptsFile := filepath.Join(scriptsPath, "works.json")
	jsonFile, err := os.Open(scriptsFile)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}

	defer jsonFile.Close()
	byteValue, _ := io.ReadAll(jsonFile)

	err = json.Unmarshal(byteValue, &works)
	if err != nil {
		log.Fatalf("failed to parse: %s", err)
	}
}
