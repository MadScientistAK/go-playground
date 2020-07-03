package main

import (
	b "bufio"
	f "fmt"
	"os"
	"strings"
	s "strings"
)

type names struct {
	fname string
	lname string
}

func main() {
	sli := make([]names, 0)
	f.Println("Enter a file name: ")
	r := b.NewReader(os.Stdin)
	filename, _ := r.ReadString('\n')
	filename = strings.TrimSpace(filename)
	filehandle, err := os.Open(filename)
	if err != nil {
		f.Println("File doesn't exist!")
	} else {
		scanner := b.NewScanner(filehandle)
		for scanner.Scan() {
			text := s.Split(scanner.Text(), " ")
			name := names{text[0], text[1]}
			sli = append(sli, name)
		}
		defer filehandle.Close()
		for _, n := range sli {
			f.Println(n.fname, n.lname)
		}
	}

}
