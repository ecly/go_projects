package main

import(
	"fmt"
	"os"
	"bufio"
	"sort"
	"strconv"
)

const defaultPrintAmount int = 10

// Input filename and optionally amount of word's occurences to print
func main(){
	args := os.Args[1:]
	file, err := os.Open(args[0])
	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)//set split

	counter := make(map[string]int)
	
	//While words are available, continue scanning
	for scanner.Scan() {
		word := scanner.Text()
		counter[word]++
	}

	//safely parse command line argument
	//regarding amount of word's occurences to print
	amountToPrint := defaultPrintAmount
	if len(args) == 2 {
		inputAmount, err := strconv.Atoi(args[1])
		if err == nil {
			amountToPrint = inputAmount
		}
	}

	sortedList := sortMap(counter)
	prettyPrint(sortedList, amountToPrint)
}

func sortMap(counter map[string]int) PairList{
	list := make(PairList, len(counter))
	i := 0 //index counter
	for key, val := range counter{
		list[i] = Pair{key, val}
		i++
	}

	//sort it in descending order by value as implemented in interface
	sort.Sort(sort.Reverse(list))
	return list
}

func prettyPrint(list PairList, amountToPrint int) {
	//Don't allow more than actual word amount
	if amountToPrint > len(list){
		amountToPrint = len(list)
	}

	for i:= 0; i < amountToPrint; i++ {
		fmt.Printf("%s occured %d times\n", list[i].Key, list[i].Value)
	}
}

// Pair used for sorting
type Pair struct{
	Key string
	Value int
}

// PairList on which we implement the sort interface
type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value < p[j].Value }