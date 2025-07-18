package race

func Race() {
	ch := make(chan bool)
	data := map[string]int{}

	go func() {
		data["0"] = 1
		ch <- true
	}()
	data["0"] = 2
	<-ch
}
