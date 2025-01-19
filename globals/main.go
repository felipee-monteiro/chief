package globals 

import (
	"errors"
	"os"
)

func GetGlobal(key string) (string, error) {
	value, err := os.LookupEnv(key) 

	if err == false {
		return "", errors.New("A variável não existe ou está vazia")
	}

	return value, nil
}