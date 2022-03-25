package main

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"sync"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go leakyFunction(&wg)
	wg.Wait()
}

func leakyFunction(wg *sync.WaitGroup) {
	defer wg.Done()
	s := make([]string, 3)
	for i := 0; i < 100000000; i++ {
		s = append(s, "debug me!!!")
	}
}
