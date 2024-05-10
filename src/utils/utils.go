package utils

import (
	"crypto/sha256"
	"encoding/hex"
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

func CreateHash(email, telephone, apto string) string {
	h := sha256.New()

	totalText := fmt.Sprintf("%s%s%s", email, telephone, apto)
	
	h.Write([]byte(totalText))

	hashText := hex.EncodeToString(h.Sum(nil))

	return hashText 
}