package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"sync"
)

func backgroundTask(wg *sync.WaitGroup) {
	defer wg.Done() // 4. Tell WaitGroup this task is done

	url := "http://127.0.0.1:8080/test"
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error fetching posts: %v", err)
	}

	defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// fmt.Println(string(body))
	io.Copy(os.Stdout, resp.Body)

}

func main() {
	var wg sync.WaitGroup // 1. Create a WaitGroup

	wg.Add(1) // 2. We’re launching 1 goroutine

	go backgroundTask(&wg) // 3. Start the goroutine

	wg.Wait() // 5. Wait for all tasks to finish

	fmt.Println("Main function done")
}
