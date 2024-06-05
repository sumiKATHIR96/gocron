package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var (
	minuteValues     = generateDefaultValue(0, 59)
	hourValues       = generateDefaultValue(0, 23)
	dayofMonthValues = generateDefaultValue(1, 31)
	monthValues      = generateDefaultValue(1, 12)
	weekValues       = generateDefaultValue(0, 6)
)

func generateDefaultValue(start int, end int) []string {
	var result []string
	for i := start; i <= end; i++ {
		result = append(result, strconv.Itoa(i))
	}
	return result
}
func validateNumber(numStr string, allowedValues []string) (string, error) {
	num, err := strconv.Atoi(numStr)
	if err != nil {
		return " ", fmt.Errorf("invalid number: %s", numStr)
	}
	min, err := strconv.Atoi(allowedValues[0])
	if err != nil {
		return " ", fmt.Errorf("invalid number: %s", numStr)
	}
	max, err := strconv.Atoi(allowedValues[len(allowedValues)-1])
	if err != nil {
		return " ", fmt.Errorf("invalid number: %s", numStr)
	}
	if num < min || num > max {
		return " ", fmt.Errorf("value out of range: %d not in [%d,%d]", num, min, max)
	}
	return numStr, nil
}

func cronField(field string, allowedValues []string) (string, error) {
	var result []string
	switch {
	case field == "*":
		return strings.Join(allowedValues, " "), nil
	case strings.Contains(field, ","):
		values := strings.Split(field, ",")
		for _, value := range values {
			result = append(result, value)
		}
		return strings.Join(result, " "), nil
	case strings.Contains(field, "/"):
		subfields := strings.Split(field, "/")
		if subfields[0] == "*" {
			subfields[0] = allowedValues[0] + "-" + allowedValues[len(allowedValues)-1]
		}
		valuetoiterate := strings.Split(subfields[0], "-")
		_, err := validateNumber(subfields[1], allowedValues)
		if err != nil {
			return "", err
		}
		increment, _ := strconv.Atoi(subfields[1])
		_, err = validateNumber(valuetoiterate[0], allowedValues)
		if err != nil {
			return "", err
		}
		start, _ := strconv.Atoi(valuetoiterate[0])
		_, err = validateNumber(valuetoiterate[1], allowedValues)
		if err != nil {
			return "", err
		}
		end, _ := strconv.Atoi(valuetoiterate[1])
		for i := start; i <= end; i += increment {
			result = append(result, strconv.Itoa(i))
		}
		return strings.Join(result, " "), nil
	case strings.Contains(field, "-"):
		subfields := strings.Split(field, "-")
		_, err := validateNumber(subfields[0], allowedValues)
		if err != nil {
			return "", err
		}
		start, _ := strconv.Atoi(subfields[0])
		_, err = validateNumber(subfields[1], allowedValues)
		if err != nil {
			return "", err
		}
		end, _ := strconv.Atoi(subfields[1])
		for i := start; i <= end; i++ {
			result = append(result, strconv.Itoa(i))
		}
		return strings.Join(result, " "), nil
	case field == "?":
		return "any", nil
	case strings.Contains(field, "#"):
		subfields := strings.Split(field, "#")
		_, err := validateNumber(subfields[0], allowedValues)
		if err != nil {
			return "", err
		}
		weekday, _ := strconv.Atoi(subfields[0])
		_, err = validateNumber(subfields[1], allowedValues)
		if err != nil {
			return "", err
		}
		nth, _ := strconv.Atoi(subfields[1])
		return fmt.Sprintf("the %dth %d", nth, weekday), nil
	default:
		_, err := validateNumber(field, allowedValues)
		if err != nil {
			return field, err
		}

		return field, err
	}
	return "", nil
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please give valid string")
		os.Exit(1)
	}

	input := os.Args[1]
	//input:="*/15 0 1,16 * 1-5 /usr/bin/find"
	//input:="15 10 15 * ? /usr/bin/find"
	fields := strings.Fields(input)
	fmt.Printf(strings.Join(fields, "..."))
	if len(fields) < 6 {
		fmt.Println("Kindly give valid input")
		os.Exit(1)
	}
	minute, err := cronField(fields[0], minuteValues)
	if err != nil {
		fmt.Println("Kindly give valid input", err)
		os.Exit(1)
	}
	hour, err := cronField(fields[1], hourValues)
	if err != nil {
		fmt.Println("Kindly give valid input", err)
		os.Exit(1)
	}
	dayOfMonth, err := cronField(fields[2], dayofMonthValues)
	if err != nil {
		fmt.Println("Kindly give valid input", err)
		os.Exit(1)
	}
	month, err := cronField(fields[3], monthValues)
	if err != nil {
		fmt.Println("Kindly give valid input", err)
		os.Exit(1)
	}
	dayofWeek, err := cronField(fields[4], weekValues)
	if err != nil {
		fmt.Println("Kindly give valid input", err)
		os.Exit(1)
	}
	fmt.Println()
	fmt.Printf("%-14s%s\n", "minutes", minute)
	fmt.Printf("%-14s%s\n", "hour", hour)
	fmt.Printf("%-14s%s\n", "day of month", dayOfMonth)
	fmt.Printf("%-14s%s\n", "month", month)
	fmt.Printf("%-14s%s\n", "day of week", dayofWeek)
	fmt.Printf("%-14s%s\n", "command", strings.Join(fields[5:], " "))
}
