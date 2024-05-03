package utils

import (
	"fmt"
	"os"
)

func GetQueryAsString(file string) (string, error) {
	queryFile := fmt.Sprintf("src/sql/%s.sql", file)
	queryBytes, err := os.ReadFile(queryFile)
	if err != nil {
		return "", err
	}

	query := string(queryBytes)

	return query, nil
}
