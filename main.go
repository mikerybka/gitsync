package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/mikerybka/github"
)

func main() {
	config, err := ReadConfig("/etc/gitsync")
	if err != nil {
		panic(err)
	}
	webhookHandler := github.WebhookHandler(config.HandleWebhook)
	err = http.ListenAndServe(":9022", webhookHandler)
	if err != nil {
		panic(err)
	}
}

func ReadConfig(dir string) (*Config, error) {
	path := filepath.Join(dir, "config.json")
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var c Config
	err = json.Unmarshal(b, &c)
	if err != nil {
		return nil, err
	}
	return &c, nil
}

type Config map[string][]string // maps branch:repo to fs location

func (c *Config) HandleWebhook(w *github.Webhook) error {
	b, _ := json.MarshalIndent(w, "", "  ")
	s := string(b)
	fmt.Println(s)
	return fmt.Errorf(s)
}
