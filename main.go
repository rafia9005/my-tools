package main

import (
	"fmt"
	"os"
	"path-scanner/tools"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("Welcome to the Scanner Tool")
	fmt.Println("Select an option:")
	fmt.Println("1. Scan paths")
	fmt.Println("2. Test SQL Injection")
	fmt.Println("3. Exit")

	var choice int
	fmt.Scanln(&choice)

	switch choice {
	case 1:
		scanPaths()
	case 2:
		testSQLInjection()
	case 3:
		os.Exit(0)
	default:
		fmt.Println("Invalid choice. Please try again.")
	}
}

func scanPaths() {
	fmt.Print("Enter target URL: ")
	var targetURL string
	fmt.Scanln(&targetURL)

	paths, err := tools.ReadPathsFromFile("path.txt")
	if err != nil {
		fmt.Println("Error reading paths from file:", err)
		return
	}

	results := tools.ScanPaths(targetURL, paths)
	printResults(results)
}

func testSQLInjection() {
	fmt.Print("Enter target URL: ")
	var targetURL string
	fmt.Scanln(&targetURL)

	results := tools.TestSQLInjection(targetURL)
	printSQLInjectionResults(results)
}

func printResults(results []tools.HasilScan) {
	fmt.Println("\nScan Results:")
	for _, result := range results {
		if result.Status == "Found" {
			color.Green("Path: %s, Status: %s\n", result.Path, result.Status)
		} else {
			color.Red("Path: %s, Status: %s\n", result.Path, result.Status)
		}
	}
}

func printSQLInjectionResults(results []tools.SQLInjectionResult) {
	fmt.Println("\nSQL Injection Scan Results:")
	for _, result := range results {
		if result.Vulnerable {
			color.Red("URL: %s, Status: Vulnerable\n", result.URL)
		} else {
			color.Green("URL: %s, Status: Not Vulnerable\n", result.URL)
		}
	}
}

