//go:build !mock

package main

import "os"

func WriteReport(report string) error {
	f, err := os.OpenFile("report.txt", os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(report)
	return err
}

func ReadReport() (string, error) {
	data, err := os.ReadFile("report.txt")
	if err != nil {
		return "", err
	}

	return string(data), nil
}
