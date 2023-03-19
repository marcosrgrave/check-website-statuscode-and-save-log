package src

import "os"

func WriteToLogFile(filePath string, log string) error {
	// https://pkg.go.dev/os
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(log + "\n")
	if err != nil {
		return err
	}

	return nil
}
