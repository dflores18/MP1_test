package config_file

import (
	"bufio"
	"os"
	"strings"
)

type Config struct {
	maxDelay int
	minDelay int
	id       int
	ip       string
	port     string
}

func ReadFile() ([]Config, error) {
	var configs []Config
	file, err := os.Open("configuration.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	delayVals := getDelayVals(strings.Split(scanner.Text(), " "))
	scanner.Scan()
	counter := 0
	for scanner.Scan() {
		temp := strings.Split(scanner.Text(), " ")
		configs[counter].minDelay = delayVals[0]
		configs[counter].maxDelay = delayVals[1]
		configs[counter].id, err = strcov.Atoi(temp[0])
		configs[counter].port = temp[1]
		configs[counter].ip = temp[2]

	}
	return configs, scanner.Err()
}

/*
func config_file() Config{
	lineInput, _ := ReadFile("configuration.txt")
	parts := string.Split(lineInput, " ")
	for i:=1; i < len(parts);i++{
		go newConnection()
	}

	maxDelay, minDelay := parts[0][0], parts[1][0]

	for i := 1; i < len(parts); i++{
		for j := 0; j < len(parts); j++ {
			ip, port := parts[i][0], parts[1][j]
		}
	}

}
