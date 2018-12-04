package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../myInput.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sum := 0

	for scanner.Scan() {
		line := scanner.Text()
		i, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}
		if line[0] == '+' {
			sum = sum + i
		} else if line[0] == '-' {
			sum = sum - i
		} else {
			fmt.Println("error")
		}
	}
	fmt.Println(sum)

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
