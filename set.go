package setenv

import (
	"errors"
	"os"
	"strings"
)

func readFile(envFile string) (string, error) {
	workingDirectory, err := os.Getwd()

	if err != nil {
		return "", errors.Unwrap(err)
	}

	filePath := workingDirectory + string(os.PathSeparator) + envFile

	file, err := os.Open(filePath)

	defer file.Close()

	if err != nil {
		return "", errors.Unwrap(err)
	}

	fileStat, _ := os.Stat(filePath)

	envFileData := make([]byte, fileStat.Size())

	count, err := file.Read(envFileData)

	if err != nil {
		return "", errors.Unwrap(err)
	}

	envFileDataString := string(envFileData[:count])

	return envFileDataString, nil
}

// SetEnv sets the corresponding environmental variable key value pairs
func SetEnv() {
	envFile := ".env"
	fileString, err := readFile(envFile)

	if err != nil {
		return
	}

	for _, variable := range strings.Split(fileString, "\r\n") {

		if strings.TrimSpace(variable) == "" {
			continue
		}

		variableParts := strings.Split(variable, "=")
		value := strings.Replace(variableParts[1], `"`, "", -1)
		value = strings.Replace(value, `'`, "", -1)
		value = strings.TrimSpace(value)
		key := strings.TrimSpace(variableParts[0])

		err := os.Setenv(key, value)

		if err != nil {
			continue
		}
	}
}
