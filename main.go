package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
    // Check if the number of arguments is correct and therefore returns error if number is missing
    if len(os.Args) != 2 {
        fmt.Println("ERROR!")
        return
    }

    // Converts argument to an integer and checks it according to conditions
    count, err := strconv.Atoi(os.Args[1])
    if err != nil || count < 0 || count > 10 {
        fmt.Println("ERROR!")
        return
    }

    // Loop to print the nickname based on number of times specified
    for i := 0; i < count; i++ {
        fmt.Println("gentle-gerbil")  
    }
}
