package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Printf("job %d success\n", i)
	}

	fmt.Println("main job success")
}
