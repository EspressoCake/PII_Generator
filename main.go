package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

const (
	bound = 400000
)

var (
	executionStart = time.Now()
)

func main() {
	fileToWrite := CreateFile()
	if fileToWrite == nil {
		log.Fatal(fileToWrite)
	}

	fields := make([]string, bound)
	var wg sync.WaitGroup

	fmt.Printf("Generating %d records.\n", bound)

	for i := 0; i < bound; i += 40000 {
		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			for index := start; index < end; index++ {
				metadata := [7]string{
					gofakeit.FirstName(),
					gofakeit.LastName(),
					strings.Title(gofakeit.Address().Street),
					gofakeit.Address().City,
					gofakeit.Address().State,
					gofakeit.CreditCardNumber(nil),
					gofakeit.SSN(),
				}

				fields[index] = strings.Join(metadata[:], `,`)
			}
		}(i, i+40000)
	}
	wg.Wait()

	WriteToFile(fileToWrite, fields)

	fmt.Printf("Execution completed in: %f seconds\n", time.Since(executionStart).Seconds())

}

func CreateFile() *os.File {
	c, err := os.Create("PII_DATA.csv")
	if err != nil {
		return nil
	}
	return c
}

func WriteToFile(file *os.File, data []string) {
	defer file.Close()

	file.WriteString("First,Last,Address,City,State,CC,SSN\n")
	file.WriteString(strings.Join(data, "\n"))
}
