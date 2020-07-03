package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readNums() []int {
	fmt.Println("Enter a sequence of upto 10 integers separated by space: ")
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	sli1 := strings.Fields(str)
	sli2 := make([]int, 0)
	for i := 0; i < len(sli1); i++ {
		num, err := strconv.Atoi(sli1[i])
		if err == nil {
			sli2 = append(sli2, num)
		} else {
			sli2 = make([]int, 0)
			break
		}
	}
	return sli2
}

func swap(numSlice []int, index int) {
	tmp := numSlice[index]
	numSlice[index] = numSlice[index+1]
	numSlice[index+1] = tmp
}

func bubbleSort(numSlice []int) {
	arrLen := len(numSlice)
	for i := 0; i < arrLen; i++ {
		for j := 0; j < arrLen-i-1; j++ {
			if numSlice[j] > numSlice[j+1] {
				swap(numSlice, j)
			}
		}
	}
}

func main() {
	sli := readNums()
	if len(sli) == 0 {
		panic("***** You entered something other than an integer. Exiting program. *****")
	} else {
		bubbleSort(sli)
		fmt.Println("***** Here is the bubble sorted sequence *****\n")
		fmt.Println(sli)
	}
}
