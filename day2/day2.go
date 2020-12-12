package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type PasswordPolicy struct {
	low, high  uint8
	filterChar string
	password   string
}

func fileToFilters(path string) ([]PasswordPolicy, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var pwPols []PasswordPolicy
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pwPols = append(pwPols, tokenize(scanner.Text()))
	}
	return pwPols, scanner.Err()
}

func tokenize(policyLine string) PasswordPolicy {
	splitPolicy := strings.Split(policyLine, ":")
	filter, password := splitPolicy[0], splitPolicy[1]
	strings.TrimSpace(filter)
	strings.TrimSpace(password)
	splitFilter := strings.Split(policyLine, " ")
	lowHigh, filterChar := splitFilter[0], splitFilter[1]
	splitRange := strings.Split(lowHigh, "-")
	lowInt, lowErr := strconv.Atoi(splitRange[0])
	highInt, highErr := strconv.Atoi(splitRange[1])
	low := uint8(lowInt)
	high := uint8(highInt)
	if lowErr != nil || highErr != nil {
		errMsg := fmt.Sprintf("Invalid range. low, high: %s, %s", lowErr, highErr)
		panic(errMsg)
	}
	retval := PasswordPolicy{low, high, filterChar, password}
	fmt.Printf("Password policy as struct: %s\n", retval)
	return retval
}

func main() {
	fileToFilters("./day2_input.txt")
}
