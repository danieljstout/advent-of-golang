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
	filter = strings.TrimSpace(filter)
	password = strings.TrimSpace(password)
	splitFilter := strings.Split(filter, " ")
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
	return retval
}

func processPolicies(policies []PasswordPolicy) (uint16) {
	var compliantPasswords uint16
	for _, policy := range policies {
		charMap := getCharCounts(policy.password)
		if(checkCompliance(charMap, policy)) {
			compliantPasswords += 1
		}
	}
	return compliantPasswords
}

func getCharCounts(password string) (map[string]uint8) {
	passwordCharCount := map[string]uint8{}
	for _, char := range strings.Split(password, "") {
		if _, ok:= passwordCharCount[char]; ok {
			passwordCharCount[char] += 1
		} else {
			passwordCharCount[char] = 1
		}
	}
	return passwordCharCount
}

func checkCompliance(passwordCharCount map[string]uint8, currentPolicy PasswordPolicy) (bool) {
	var filterCharCount uint8
	if _, ok := passwordCharCount[currentPolicy.filterChar]; ok {
		filterCharCount = passwordCharCount[currentPolicy.filterChar]
		return currentPolicy.low <= filterCharCount && currentPolicy.high >= filterCharCount
	} else {
		return false
	}
}

func main() {
	filters, err := fileToFilters("./day2_input.txt")
	if err != nil {
		panic("Your file is bad and you should feel bad")
	}
	compliantPasswords := processPolicies(filters)
	fmt.Printf("Number of compliant passwords: %i\n", compliantPasswords)
}
