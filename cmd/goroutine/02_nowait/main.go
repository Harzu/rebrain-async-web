package main

import "fmt"

func main() {
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Printf("job %d\n", i)
		}
	}()

	fmt.Println("main job")
}
