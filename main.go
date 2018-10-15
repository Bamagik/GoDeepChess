package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter Move: ")
	for scanner.Scan() && strings.ToLower(scanner.Text()) != "exit" {
		text := scanner.Text()
		fmt.Println(text)
		fmt.Print("Enter Move: ")
	}
}
