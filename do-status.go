package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
)

const (
	DATE_LAYOUT = "2. Jan 2006, 15:04:05"
	COLOR_LABEL = "#999999"
)

type StatusLine struct {
	Lines []StatusSegment
}

func (l *StatusLine) Add() *StatusSegment {
	var segment StatusSegment = make(StatusSegment)
	l.Lines = append(l.Lines, segment)
	return &segment

}

func (l *StatusLine) AddLabel(label string) {
	l.Add().
		SetFullText(label).
		SetColor(COLOR_LABEL).
		SetSeparator(false).
		SetSeparatorWidth(0).
		SetMinWidthString("9h59m59s")

}

type StatusSegment map[string]interface{}

func (l *StatusSegment) SetFullText(text string) *StatusSegment {
	(*l)["full_text"] = text
	return l

}

func (l *StatusSegment) SetColor(color string) *StatusSegment {
	(*l)["color"] = color
	return l

}

func (l *StatusSegment) SetSeparator(s bool) *StatusSegment {
	(*l)["separator"] = s
	return l

}

func (l *StatusSegment) SetSeparatorWidth(w int) *StatusSegment {
	(*l)["separator_block_width"] = w
	return l

}

func (l *StatusSegment) SetMinWidthString(s string) *StatusSegment {
	(*l)["min_width"] = s
	return l

}

func main() {
	fmt.Println(`{ "version": 1 }`)
	fmt.Println(`[`)
	fmt.Println(`[]`)

	for {
		printLine(getLine())
		time.Sleep(200 * time.Millisecond)
	}

}

func printLine(line StatusLine) {
	bytes, err := json.Marshal(line.Lines)
	if err != nil {
		panic(err)
	}

	fmt.Print(`,`)
	fmt.Print(string(bytes))
	fmt.Println()

}

func getLine() StatusLine {

	var line StatusLine

	line.AddLabel("Uptime: ")
	line.Add().SetFullText(getUptime())
	line.Add().SetFullText(time.Now().Format(DATE_LAYOUT))

	return line

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
