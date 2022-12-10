package helpers

import (
	"encoding/json"
	"os"
)

func ReadAccountsJson(fileName string) ([]string, error) {
	var sc []string

	content, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, &sc)
	if err != nil {
		return nil, err
	}

	return sc, nil
}
