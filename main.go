package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("This should be the last simple test! If you see this we are good to go.")
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Press enter to exit.")
	reader.ReadString('\n')
}
