package utils

import (
	"bufio"
	"errors"
	"os"
	"strings"
)

func File_put_contents(filename string, text string) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	f.WriteString(text)
	f.Close()
}

func File_append_contents(filename string, text string) {
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	f.WriteString(text)
	f.Close()
}

func Fs_exists(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func File_get_lines(path string) ([]string, error) {
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

func FindLine(filename string, tofind string) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if strings.Contains(text, tofind) {
			return text, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", errors.New("Nothnig found")
}
