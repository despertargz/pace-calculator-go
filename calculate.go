package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	distance, _ := strconv.ParseFloat(os.Args[1], 64)
	pace := os.Args[2]

	time := calculate(pace, distance)
	fmt.Println(time)
}

func toSeconds(pace string) float64 {
	split := strings.Split(pace, ":")
	minute, _ := strconv.Atoi(split[0])
	second, _ := strconv.Atoi(split[1])

	return float64((minute * 60) + second)
}

func toTime(numSeconds float64) string {
	seconds := int(numSeconds)

	hours := seconds / 3600
	seconds = seconds - (hours * 3600)

	minutes := seconds / 60
	seconds = seconds - (minutes * 60)

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}

func calculate(pace string, distance float64) string {
	paceSeconds := toSeconds(pace)
	timeSeconds := paceSeconds * distance
	return toTime(timeSeconds)
}
