package main

import (
	"00/GameWorld/config"
	"bufio"
	"errors"
	"os"
	"strconv"
)

func run() error {
	if len(os.Args) != 4 {
		return errors.New("ERROR1")
	}

	height, err := strconv.Atoi(os.Args[1])
	if err != nil {
		return errors.New("ERROR2")
	}

	width, err := strconv.Atoi(os.Args[2])
	if err != nil {
		return errors.New("ERROR3")
	}

	percent, err := strconv.Atoi(os.Args[3])
	if err != nil {
		return errors.New("ERROR4")
	}

	//fmt.Println(height, width, percent)
	file, err := os.Create("config.txt")
	if err != nil {
		return errors.New("Fail create file")
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	str := strconv.Itoa(height) + "x" + strconv.Itoa(width) + " " + strconv.Itoa(percent) + "%"

	_, err = writer.WriteString(str)
	if err != nil {
		return errors.New("no write")
	}

	err = writer.Flush()

	if err != nil {
		return errors.New("error")
	}
	return nil
}

func main() {
	config.New()
	run()
}
