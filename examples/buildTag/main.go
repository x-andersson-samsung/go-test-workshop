package main

import "log"

func main() {
	if err := WriteReport("Test report"); err != nil {
		panic(err)
	}

	report, err := ReadReport()
	if err != nil {
		panic(err)
	}
	log.Println(report)
}
