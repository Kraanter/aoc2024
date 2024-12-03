package week1

import (
	"os"
	"strconv"
	"strings"
)

const input = "7 6 4 2 1\n1 2 7 8 9\n9 7 6 2 1\n1 3 2 4 5\n8 6 4 4 1\n1 3 6 7 9"

func parseInputDay2() [][]int {
	rows := [][]int{}
	row := []int{}

	var currNumBuilder strings.Builder

	content, _ := os.ReadFile("week1/inputs/day2.txt")

	for _, val := range content {
		if val == 10 {
			num, _ := strconv.Atoi(currNumBuilder.String())
			row = append(row, num)

			rows = append(rows, row)

			currNumBuilder.Reset()
			row = []int{}
			continue
		}

		if val == 32 {
			num, _ := strconv.Atoi(currNumBuilder.String())
			row = append(row, num)
			currNumBuilder.Reset()
			continue
		}

		currNumBuilder.WriteString(string(val))
	}

	return rows
}

func checkSafe(isInc bool, row []int, i int) bool {
	if row[i+1] == row[i] {
		return false
	}

	if isInc {
		return row[i+1]-row[i] <= 3
	} else {
		return row[i]-row[i+1] <= 3
	}

}

func DayTwoPuzzleOne() int {
	rows := parseInputDay2()
	amountSafe := 0

	for _, row := range rows {
		isInc := false
		isSafe := false

		for i, num := range row {
			if i+1 >= len(row) {
				break
			}

			if i == 0 {
				isInc = num < row[i+1]
			} else if isInc && num > row[i+1] {
				isSafe = false
				break
			} else if !isInc && num < row[i+1] {
				isSafe = false
				break
			}

			isSafe = checkSafe(isInc, row, i)

			if !isSafe {
				break
			}
		}

		if isSafe {
			amountSafe++
		}
	}

	return amountSafe
}

func DayTwoPuzzleTwo() int {

	rows := parseInputDay2()
	amountSafe := 0

	for _, row := range rows {
		isInc := false
		isSafe := false
		passUsed := false

		for i, num := range row {
			if i+1 >= len(row) {
				break
			}

			if i == 0 {
				isInc = num < row[i+1]
			} else if isInc && num > row[i+1] {
				isSafe = false
				break
			} else if !isInc && num < row[i+1] {
				isSafe = false
				break
			}

			isSafe = checkSafe(isInc, row, i)

			if !isSafe {
				if !passUsed {
					passUsed = true
					continue
				}

				break
			}
		}

		if isSafe {
			amountSafe++
		}
	}

	return amountSafe
}