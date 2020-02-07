package main

// example 1
// import (
// 	"fmt"
// 	"time"
// )

// func say(s string) {
// 	for i := 0; i < 3; i++ {
// 		fmt.Println(s)
// 	}
// }

// func main() {
// 	go say("hey")
// 	go say("there")
// 	time.Sleep(time.Second)
// }

// example 2
import (
	"fmt"
	"sync"
	"time"
)

func findUserById(uid string) string {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println("user found")
	return "John"
}

func checkSubStatus(uid string) bool {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println("subscription found")
	return true
}

func getQuotes(uid string) int {
	defer wg.Done()
	time.Sleep(time.Second)
	fmt.Println("quote found")
	return 3000
}

func elapsed(what string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", what, time.Since(start))
	}
}

var wg sync.WaitGroup

func main() {
	defer elapsed("quote API")()
	wg.Add(1)
	go findUserById("a1")
	wg.Add(1)
	go checkSubStatus("a1")
	wg.Wait()

	wg.Add(1)
	go getQuotes("a1")
	wg.Wait()
}
