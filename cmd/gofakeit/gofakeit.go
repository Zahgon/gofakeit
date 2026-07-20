package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/brianvoe/gofakeit/v7"
)

var errNoFuncRunMsg = errors.New("could not find function to run\nrun gofakeit help or gofakeit list for available functions")

func main() {
	var loop int = 1
	args := os.Args[1:]

	cleanArgs := []string{}
	for i := 0; i < len(args); i++ {

		if strings.Contains(args[i], "-loop") {

			split := strings.Split(args[i], "=")

			var err error
			loop, err = strconv.Atoi(split[1])
			if err != nil {
				fmt.Println("Error converting loop flag to int")
				os.Exit(1)
			}
		}

		if !strings.HasPrefix(args[i], "-") {
			cleanArgs = append(cleanArgs, args[i])
		}
	}
	args = cleanArgs

	out, err := mainFunc(0, args, loop)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Printf("%s", out)
}

func mainFunc(seed uint64, args []string, loop int) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func runFunction(faker *gofakeit.Faker, function string, args []string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func stringInSlice(a string, list []string) bool { _ = "STUB: not implemented"; return false }

func listOutput(selectedCategory string, selectedFunction string) string {
	_ = "STUB: not implemented"
	return ""
}
