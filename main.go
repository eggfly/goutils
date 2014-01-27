package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"time"
)

func testRamDB() {
	t := func(count int) {
		var factor = float64(count) / 1000 / 1000
		fmt.Println(factor)
		m := make(map[string]interface{})
		t := time.Now()
		for i := 0; i < count; i++ {
			m[string(i)] = i
		}
		fmt.Println(float64(time.Now().Sub(t).Nanoseconds()) / float64(count))

		t = time.Now()
		for i := 0; i < count; i++ {
			_ = m[string(i)]
		}
		fmt.Println(float64(time.Now().Sub(t).Nanoseconds()) / float64(count))
		runtime.GC()
	}
	// runtime.GOMAXPROCS(runtime.NumCPU())
	t(2 * 1000 * 1000)
	t(4 * 1000 * 1000)
	t(6 * 1000 * 1000)
	t(8 * 1000 * 1000)
	t(10 * 1000 * 1000)
	t(20 * 1000 * 1000)
	t(40 * 1000 * 1000)
	t(60 * 1000 * 1000)
	t(80 * 1000 * 1000)
	t(100 * 1000 * 1000)
}

func testFileOps() {
	t := time.Now()
	for i := 0; i < 10000; i++ {
		fmt.Println(i)
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		file, err := os.Open("main.go")
		_ = file
		if err != nil {
			log.Panic(err)
			log.Println(err)
		}
	}
	fmt.Println(time.Now().Sub(t))
}

func main() {
	// testRamDB()
	testFileOps()
}
