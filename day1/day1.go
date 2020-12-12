package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func textSlicer(fileLines []string) []int {
	var luckyNumbers []int
	for _, line := range fileLines {
		i, err := strconv.Atoi(line)
		if err != nil {
			errMsg := fmt.Sprintf("There is an error in numberSlicer. %s isn't a number or something. You figure it out!", line)
			panic(errMsg)
		}
		luckyNumbers = append(luckyNumbers, i)
	}
	return luckyNumbers
}

func findTwenty20Doublet(s []int) [2]int {
	var winningNumbers [2]int
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s); j++ {
			if j == i {
				continue
			} else if (s[i] + s[j]) != 2020 {
				continue
			} else {
				winningNumbers[0] = s[i]
				winningNumbers[1] = s[j]
				break
			}
		}
	}
	if winningNumbers == [2]int{0, 0} {
		panic("a problem")
	}
	return winningNumbers
}

func findTwenty20Triplet(s []int) [3]int {
	var winningNumbers [3]int
	for i := 0; i < len(s)-2; i++ {
		for j := i + 1; j < len(s)-1; j++ {
			for k := j + 2; k < len(s); k++ {
				if s[i]+s[j]+s[k] == 2020 {
					winningNumbers[0] = s[i]
					winningNumbers[1] = s[j]
					winningNumbers[2] = s[k]
				}
			}
		}
	}
	if winningNumbers == [3]int{0, 0, 0} {
		panic("...it has the words 'DON'T PANIC' in large, friendly letters...")
	}
	return winningNumbers
}

func main() {
	aPath := "./day1-input.txt"
	ourString, err := readLines(aPath)
	if err != nil {
		panic("oh good god! There's no file! abandon ship!!!")
	}
	ourSlice := textSlicer(ourString)
	theWinningingDoublet := findTwenty20Doublet(ourSlice)
	theWinningingTriplet := findTwenty20Triplet(ourSlice)
	theDoubletWinner := theWinningingDoublet[0] * theWinningingDoublet[1]
	theTripletWinner := theWinningingTriplet[0] * theWinningingTriplet[1] * theWinningingTriplet[2]
	fmt.Println(theDoubletWinner, "\n")
	fmt.Println(theTripletWinner, "\n")
}
