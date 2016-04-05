package main

import (
	"bufio"
	"flag"
	"log"
	"os"
	"os/exec"
	"strings"
	"text/template"
)

subject := "Weekly Uberspace Report"

var accessLogPath = flag.String("access", "", "path to an access_log file")
var mailAccount = flag.String("mail", "", "mail address that should receive the report")

func main() {
	flag.Parse()

	if *accessLogPath == "" || *mailAccount == "" {
		log.Fatal("'access' and 'mail' flags are required.")
	}

	accessLog := parseAccessLog(*accessLogPath)

	tmpl, err := template.ParseFiles("mail.template")

	if err != nil {
		log.Fatal(err)
	}


	cmd := exec.Command("mailx", "-s", subject, *mailAccount)
	in, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Start()
	tmpl.Execute(in, accessLog)
	in.Close()
	err = cmd.Wait()
	if err != nil {
		log.Println("Make sure mailx is installed and configured.")
		log.Fatal(err)
	}
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
