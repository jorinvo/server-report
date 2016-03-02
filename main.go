package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	var accessLogPath = flag.String("access", "access_log", "path to an access_log file")
	flag.Parse()

	accessLog := parseAccessLog(*accessLogPath)

	fmt.Println("Most accessed resources:")

	for _, p := range accessLog.Top(20) {
		fmt.Println(p.int, p.string)
	}

	fmt.Println("")
	fmt.Println("Total: ", accessLog.Total())
}

func parseAccessLog(path string) Histogram {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	histogram := Histogram{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		method := fields[5][1:]
		uri := fields[6]
		histogram.Add(method + " " + uri)
	}

	return histogram
}
