package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

const (
	DATE_LAYOUT = "2. Jan 2006 15:04:05"
)

func main() {
	fmt.Println()
	fmt.Println()

	for {
		printStatus()
		time.Sleep(200 * time.Millisecond)
	}

}

func printStatus() {
	sts := ""

	sts += "Up: " + getUptime()
	sts += " | "
	sts += time.Now().Format(DATE_LAYOUT)

	fmt.Print(sts)
	fmt.Println()
}

func getUptime() string {
	bytes, err := ioutil.ReadFile("/proc/uptime")

	if err != nil {
		return fmt.Sprint(err)
	}

	secondstr := strings.Split(string(bytes), " ")[0]
	seconds, err := strconv.ParseFloat(secondstr, 64)
	duration := time.Duration(int(seconds) * int(time.Second))

	return duration.String()

}
