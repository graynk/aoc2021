package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func part1(binaryReports []string) int64 {
	reportSize := len(binaryReports[0])

	gammaBinary := ""
	epsilonBinary := ""

	for i := 0; i < reportSize; i++ {
		zeroCount := 0
		oneCount := 0
		for _, report := range binaryReports {
			if rune(report[i]) == '0' {
				zeroCount++
			} else {
				oneCount++
			}
		}
		if zeroCount > oneCount {
			gammaBinary += "0"
			epsilonBinary += "1"
		} else {
			gammaBinary += "1"
			epsilonBinary += "0"
		}
	}

	gamma, err := strconv.ParseInt(gammaBinary, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	epsilon, err := strconv.ParseInt(epsilonBinary, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	return epsilon * gamma
}

func part2(binaryReports []string) int64 {
	reportSize := len(binaryReports[0])

	oxygenReports := make([]string, 0, len(binaryReports))
	oxygenReports = append(oxygenReports, binaryReports...)

	for i := 0; i < reportSize && len(oxygenReports) > 1; i++ {
		zeroBitReports := make([]string, 0, len(oxygenReports))
		oneBitReports := make([]string, 0, len(oxygenReports))
		for _, report := range oxygenReports {
			if rune(report[i]) == '0' {
				zeroBitReports = append(zeroBitReports, report)
			} else {
				oneBitReports = append(oneBitReports, report)
			}
		}
		if len(oneBitReports) >= len(zeroBitReports) {
			oxygenReports = oneBitReports
		} else {
			oxygenReports = zeroBitReports
		}
	}

	if len(oxygenReports) != 1 {
		log.Fatalf("for oxygen we got left with not 1: %d", len(oxygenReports))
	}

	oxygenReport, err := strconv.ParseInt(oxygenReports[0], 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	co2Reports := make([]string, 0, len(binaryReports))
	co2Reports = append(co2Reports, binaryReports...)

	for i := 0; i < reportSize && len(co2Reports) > 1; i++ {
		zeroBitReports := make([]string, 0, len(co2Reports))
		oneBitReports := make([]string, 0, len(co2Reports))
		for _, report := range co2Reports {
			if rune(report[i]) == '0' {
				zeroBitReports = append(zeroBitReports, report)
			} else {
				oneBitReports = append(oneBitReports, report)
			}
		}
		if len(oneBitReports) >= len(zeroBitReports) {
			co2Reports = zeroBitReports
		} else {
			co2Reports = oneBitReports
		}
	}

	if len(co2Reports) != 1 {
		log.Fatalf("for co2 we got left with not 1: %d", len(co2Reports))
	}

	co2Report, err := strconv.ParseInt(co2Reports[0], 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	return oxygenReport * co2Report
}

func main() {
	input, err := os.Open("./day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	fileScanner := bufio.NewScanner(input)
	report := make([]string, 0, 1000)
	for fileScanner.Scan() {
		report = append(report, fileScanner.Text())
	}

	fmt.Println(part1(report))
	fmt.Println(part2(report))
}
