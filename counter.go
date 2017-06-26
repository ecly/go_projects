package main

import(
	"fmt"
	"os"
	"bufio"
)

func main(){
	args := os.Args[1:]
	file, err := os.Open(args[0])
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)//set split
	
	status := scanner.Scan()
	if status == false {
		fmt.Println("Scan failed")
	}

	fmt.Println("First word: ", scanner.Text())
}