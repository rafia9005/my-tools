package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/fatih/color"
)

type HasilScan struct {
	Path   string
	Status string
}

func ReadPathsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var paths []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		paths = append(paths, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return paths, nil
}

func ScanPaths(baseURL string, paths []string) []HasilScan {
	client := &http.Client{Timeout: 5 * time.Second}
	results := []HasilScan{}

	for _, path := range paths {
		url := fmt.Sprintf("%s/%s", baseURL, path)
		status := CheckPath(client, url)
		results = append(results, HasilScan{Path: path, Status: status})
	}

	return results
}

func CheckPath(client *http.Client, url string) string {
	resp, err := client.Get(url)
	if err != nil {
		return "Error"
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		return "Found"
	}
	return "Not Found"
}

func printResults(results []HasilScan) {
	fmt.Println("\nScan Results:")
	for _, result := range results {
		if result.Status == "Found" {
			color.Green("Path: %s, Status: %s\n", result.Path, result.Status)
		} else {
			color.Red("Path: %s, Status: %s\n", result.Path, result.Status)
		}
	}
}

func main() {
	fmt.Println("Welcome to the Path Scanner Tool")
	fmt.Print("Enter base URL (e.g., https://example.com): ")
	var baseURL string
	fmt.Scanln(&baseURL)

	paths, err := ReadPathsFromFile("path.txt")
	if err != nil {
		fmt.Println("Error reading paths from file:", err)
		return
	}

	results := ScanPaths(baseURL, paths)
	printResults(results)
}

