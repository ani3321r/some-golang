package main

import (
	"fmt"
	//"math/rand"
	"sync"
	"time"
)

var m1 = sync.RWMutex{}
var m = sync.Mutex{}      //mutual exclution
var wg = sync.WaitGroup{} //they are just counters
var dbData = []string{"a1", "a2", "a3", "a4", "a5"}
var results = []string{}

func main() {
	t0 := time.Now()
	for i := 0; i < len(dbData); i++ {
		wg.Add(1)    //this incriments the counter
		go dbCall(i) //the go keyword is to run all the tasks concurrently
	}
	wg.Wait() //this will wait for all tasks to go 0 and then will execute
	fmt.Printf("\nExecution time: %v", time.Since(t0))
	fmt.Printf("\nResult %v", results)
}

func dbCall(i int) {
	//var delay float32 = rand.Float32() * 2000
	var delay float32 = 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	fmt.Println("\nresult of db: ", dbData[i])
	// m.Lock() //until the lock is released for the previous call the next call is not being made so it is for memory protection purpose
	// results = append(results, dbData[i])
	// m.Unlock()
	save(dbData[i])
	log()
	wg.Done() //this decrements the counter
}

func save(result string) {
	m.Lock() //this is full lock
	results = append(results, result)
	m.Unlock()
}

func log() {
	m1.RLock() //this allows to read from other slices
	fmt.Printf("\nNew Result are %v", results)
	m1.RUnlock()
}
