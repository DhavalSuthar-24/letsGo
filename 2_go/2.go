// package main

// import (
// 	"fmt"
// 	"time"
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// 
// )

// func worker(id int) {
// 	fmt.Printf("Worker %d started\n", id)
// 	time.Sleep(1 * time.Second)
// 	fmt.Printf("Worker %d done\n", id)
// }

// func main() {
// 	for i := 1; i <= 1000; i++ {
// 		go worker(i) // Try 1000 threads in Java—it’ll crash!
// 	}
// 	time.Sleep(2 * time.Second)
// }
package main

import (
	"fmt"
	"log"
)

func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, fmt.Errorf("cannot divide by zero") // Create error
    }
    return a / b, nil // No error
}

func main() {
    // Example 1: Basic error handling
    _, err := divide(2, 0)
    if err != nil {
        log.Printf("Division failed: %v", err) // Log the error
        // Continue execution instead of returning
    }

    // Example 2: Wrapping errors
    result, err := divide(10, 0)
    if err != nil {
        wrappedErr := fmt.Errorf("calculation failed: %w", err)
        log.Println(wrappedErr)
        // You could return here if this wasn't main()
    } else {
        fmt.Println("Result:", result)
    }
}