// Package utils contains very cool functions for utility stuff
package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ReadAndParseLines reads a file line by line and parses each line using the provided parse function
//
// Parameters:
//   - filePath: The path to the file to be read
//   - parseFunc: A function that parses the line
//
// Returns:
//   - []T: A slice containing the parsed lines
//   - error: An error if u fucked up somewhere
func ReadAndParseLines[T any](filePath string, parseFunc func(string) (T, error)) ([]T, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var values []T
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		val, err := parseFunc(line)
		if err != nil {
			return nil, err
		}
		values = append(values, val)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return values, nil
}

// ReadSplitAndParseLines reads a file line by line, splits each line using a separator,
// and parses the resulting parts using provided parsing functions
// (cus i cant be bothered to make weird ahh parsing functions to do this with the other function)
//
// Parameters:
//   - filePath: The path to the file to be read
//   - separator: The string used to split each line into two parts
//   - leftParseFunc: A function that parses the left part of each split line
//   - rightParseFunc: A function that parses the right part of each split line
//
// Returns:
//   - []T: A slice containing the parsed left parts of each line
//   - []U: A slice containing the parsed right parts of each line
//   - error: An error if u fucked up somewhere
func ReadSplitAndParseLines[T, U any](filePath string, separator string, leftParseFunc func(string) (T, error), rightParseFunc func(string) (U, error)) ([]T, []U, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	var left []T
	var right []U
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, separator)
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid line format: %s", line)
		}

		leftVal, err := leftParseFunc(parts[0])
		if err != nil {
			return nil, nil, err
		}

		rightVal, err := rightParseFunc(parts[1])
		if err != nil {
			return nil, nil, err
		}

		left = append(left, leftVal)
		right = append(right, rightVal)
	}

	if err := scanner.Err(); err != nil {
		return nil, nil, err
	}

	return left, right, nil
}
