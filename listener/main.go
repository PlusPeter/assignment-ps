package main

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
)

// myMessage is the format of the published messages.
type myMessage struct {
	Type string `json:"type"`
	Data int    `json:"data"`
}

func main() {

}

// printToFile is a function that represents a call to an external service that
// can have temporary failures but will eventually succeed.
func printToFile(f *os.File, s string) error {
	x := rand.Int31n(10)
	if x < 6 {
		return errors.New("sorry couldn't make it")
	}

	fs := fmt.Sprintf("processed msg: %s\n", s)
	if _, err := f.WriteString(fs); err != nil {
		panic(fmt.Sprintf("unexpected error while writing on file, err: %v", err))
	}

	return nil
}
