//go:build mock

package main

import "log"

var data string

func WriteReport(report string) error {
	log.Printf("Writing mock report: %s", report)
	data = report
	return nil
}

func ReadReport() (string, error) {
	log.Printf("Reading mock report: %s", data)
	return data, nil
}
