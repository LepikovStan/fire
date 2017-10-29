package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"flag"
)

func readFile(path string) []string {
	var result []string
	inFile, _ := os.Open(path)
	defer inFile.Close()
	scanner := bufio.NewScanner(inFile)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		result = append(result, scanner.Text())
	}
	return result
}

func makeMatrix(mlength int) [][]string {
	matrix := make([][]string, mlength)
	for i := 0; i < mlength; i++ {
		matrix[i] = make([]string, mlength)
		for j := 0; j < mlength; j++ {
			matrix[i][j] = "turn off"
		}
	}
	return matrix
}

func makeMatrixD(mlength int) [][]int {
	matrix := make([][]int, mlength)
	for i := 0; i < mlength; i++ {
		matrix[i] = make([]int, mlength)
		for j := 0; j < mlength; j++ {
			matrix[i][j] = 0
		}
	}
	return matrix
}

func parseCommand(command string) (string, []int64, []int64) {
	splittedCommand := strings.Split(command, " ")
	splittedCommandLength := len(splittedCommand)
	var commandType string
	var firstPoint []string
	var secondPoint []string

	if splittedCommandLength == 4 {
		commandType = "toggle"
		firstPoint = strings.Split(splittedCommand[1], ",")
		secondPoint = strings.Split(splittedCommand[3], ",")
	} else {
		commandType = fmt.Sprintf("turn %s", splittedCommand[1])
		firstPoint = strings.Split(splittedCommand[2], ",")
		secondPoint = strings.Split(splittedCommand[4], ",")
	}
	x1, _ := strconv.ParseInt(firstPoint[0], 10, 32)
	x2, _ := strconv.ParseInt(secondPoint[0], 10, 32)
	y1, _ := strconv.ParseInt(firstPoint[1], 10, 32)
	y2, _ := strconv.ParseInt(secondPoint[1], 10, 32)

	xCoords := []int64{ x1, x2 }
	yCoords := []int64{ y1, y2 }

	return commandType, xCoords, yCoords
}

func execCommands(commands []string, matrix[][]string) [][]string {
	for _, command := range commands {
		commandType, xCoords, yCoords := parseCommand(command)
		for y := yCoords[0]; y <= yCoords[1]; y++ {
			for x := xCoords[0]; x <= xCoords[1]; x++ {
				if commandType == "toggle" {
					if matrix[y][x] == "turn on" {
						matrix[y][x] = "turn off"
					} else {
						matrix[y][x] = "turn on"
					}
				} else {
					matrix[y][x] = commandType
				}
			}
		}
	}
	return matrix
}

func execCommandsD(commands []string, matrix[][]int) [][]int {
	for _, command := range commands {
		commandType, xCoords, yCoords := parseCommand(command)
		for y := yCoords[0]; y <= yCoords[1]; y++ {
			for x := xCoords[0]; x <= xCoords[1]; x++ {
				if commandType == "toggle" {
					matrix[y][x] += 2
				} else if commandType == "turn on" {
					matrix[y][x]++
				} else {
					matrix[y][x]--
				}

				if matrix[y][x] < 0 {
					matrix[y][x] = 0
				}
			}
		}
	}
	return matrix
}

func simple() {
	matrixLength := 1000
	matrix := makeMatrix(matrixLength)
	commands := readFile("input.txt")
	matrix = execCommands(commands, matrix)
	turnOnCounter := 0

	for _, line := range(matrix) {
		for _, cell := range(line) {
			if cell == "turn on" {
				turnOnCounter++
			}
		}
	}

	fmt.Println("Result = ", turnOnCounter)
}

func difficult() {
	fmt.Println()
	matrixLength := 1000
	matrix := makeMatrixD(matrixLength)
	commands := readFile("input.txt")
	matrix = execCommandsD(commands, matrix)

	brightnessCounter := 0
	for _, line := range(matrix) {
		for _, brightness := range(line) {
			brightnessCounter += brightness
		}
	}

	fmt.Println("Result = ", brightnessCounter)
}

var d bool
func initFlags() {
	flag.BoolVar(&d, "d", false, "")
	flag.Parse()
}

func main() {
	start := time.Now()
	initFlags()
	if d {
		difficult()
	} else {
		simple()
	}
	end := time.Now()
	fmt.Println(end.Sub(start))
}

// 400410
// 15343601
