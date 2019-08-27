package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	file, err := os.Open("iso3166-1")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text()[:50])
		country = scanner.Text()[:48]
		b = scanner.Text()[48:50]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
