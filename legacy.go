type Event struct {
	ID int
}
func stage(name string, in <-chan Event, out chan<- Event, metric chan<- string) {
	for e := range in {
		time.Sleep(25 * time.Millisecond)
		metric <- name
		out <- e
	}
}
func main() {
	in := make(chan Event)
	a := make(chan Event)
	b := make(chan Event)
	out := make(chan Event)
	metric := make(chan string)
	go stage("A", in, a, metric)
	go stage("B", a, b, metric)
	go stage("IO", b, out, metric)
	go func() {
		for m := range metric {
			fmt.Println("metric", m)
		}
	}()
	for i := 0; i < 5; i++ {
		in <- Event{ID: i}
		fmt.Println("trace", i, time.Now())
	}
}
