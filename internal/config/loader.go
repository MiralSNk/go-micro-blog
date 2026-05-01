package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

// Load читает yml файл и выгружает данные в структуру
func Load(path string, out any) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("failed: not found config -> %s, %v", path, err)
	}

	return yaml.Unmarshal(data, out)
}
