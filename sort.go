package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func concSort(numSlice []int, wg *sync.WaitGroup) {
	fmt.Println(numSlice)
	sort.Ints(numSlice)
	wg.Done()
}

func main() {
	fmt.Println("Enter integers for sorting (At least 4):")
	reader := bufio.NewReader(os.Stdin)
	temp_str, _ := reader.ReadString('\n')
	temp_arr := strings.Fields(temp_str)
	main_arr := make([]int, 0)
	for _, i := range temp_arr {
		num, err := strconv.Atoi(i)
		if err == nil {
			main_arr = append(main_arr, num)
		} else {
			main_arr = make([]int, 0)
			break
		}
	}
	var wg sync.WaitGroup
	arr_len := len(main_arr)
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go concSort(main_arr[i*arr_len/4:arr_len*(i+1)/4], &wg)
	}
	wg.Add(1)
	go concSort(main_arr[arr_len*3/4:], &wg)
	wg.Wait()
	sort.Ints(main_arr)
	fmt.Println(main_arr)
}
