type Span struct {
	ID int
	Stage string
}

func process(stage string, in <-chan int, out chan<- int, m chan<- string, l chan<- string, t chan<- Span) {
	for id := range in {
		start := time.Now()
		time.Sleep(30 * time.Millisecond)
		m <- stage
		l <- fmt.Sprintf("event %d %s", id, stage)
		t <- Span{ID: id, Stage: stage}
		fmt.Println(stage, id, time.Since(start))
		out <- id
	}
}

func main() {
	in := make(chan int)
	a := make(chan int)
	b := make(chan int)
	out := make(chan int)
	metrics := make(chan string)
	logs := make(chan string)
	traces := make(chan Span)

	go process("StageA", in, a, metrics, logs, traces)
	go process("StageB", a, b, metrics, logs, traces)
	go process("IO", b, out, metrics, logs, traces)

	for i := 0; i < 5; i++ {
		in <- i
	}
	time.Sleep(time.Second)
}
