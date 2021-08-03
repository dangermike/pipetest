package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "emit" {
		emit(1<<22, 1<<35)
	} else {
		consume(1 << 22)
	}
}

func consume(bufferSize int) {
	buf := make([]byte, bufferSize)
	total := 0
	var err error
	start := time.Now()
	for err == nil {
		var cnt int
		cnt, err = os.Stdin.Read(buf)
		total += cnt
	}
	duration := time.Since(start)
	fmt.Print("total:        ")
	fmt.Println(total)
	fmt.Print("duration:     ")
	fmt.Println(duration.String())
	fmt.Print("rate (GiBps): ")

	fmt.Printf("%0.03f\n", (float64(total)/duration.Seconds())/float64(1<<30))
}

func emit(dataSize int, totalBytes int) {
	data := make([]byte, dataSize)
	for i := 0; i < len(data)-1; i++ {
		data[i] = 'x'
	}
	data[len(data)-1] = '\n'

	for total := 0; total < totalBytes; {
		cnt, err := os.Stdout.Write(data)
		if err != nil {
			return
		}
		total += cnt
	}
}
