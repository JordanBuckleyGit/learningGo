package main

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

// Serve the home page with the file upload form
func homePage(w http.ResponseWriter, r *http.Request) {
	// Serve the simple HTML form for file upload
	tmpl := `<html>
				<head><title>Upload CSV</title></head>
				<body>
					<h1>Upload CSV File</h1>
					<form enctype="multipart/form-data" action="/upload" method="post">
						<input type="file" name="file" accept=".csv">
						<input type="submit" value="Upload">
					</form>
				</body>
			</html>`
	fmt.Fprint(w, tmpl)
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	// Ensure it's a POST request
	if r.Method != "POST" {
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the uploaded file
	file, _, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Error processing file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Read CSV data
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		http.Error(w, "Error reading CSV", http.StatusInternalServerError)
		return
	}

	// Analyze CSV data
	sum, count := analyzeCSV(records)

	resultPage := `<html>
					<head><title>Results</title></head>
					<body>
						<h1>CSV Analysis Results</h1>
						<p>Sum of numbers: {{.Sum}}</p>
						<p>Count of numbers: {{.Count}}</p>
						<a href="/">Upload another file</a>
					</body>
				</html>`
	tmpl := template.Must(template.New("result").Parse(resultPage))
	tmpl.Execute(w, map[string]float64{
		"Sum":   sum,
		"Count": float64(count),
	})
}

func analyzeCSV(records [][]string) (float64, int) {
	var sum float64
	var count int

	for i, row := range records {
		if i == 0 {
			continue
		}
		for _, value := range row {
			num, err := strconv.ParseFloat(value, 64)
			if err == nil {
				sum += num
				count++
			}
		}
	}
	return sum, count
}

func main() {
	// Route for the home page (file upload form)
	http.HandleFunc("/", homePage)

	http.HandleFunc("/upload", uploadFile)

	// Start the server
	fmt.Println("Server starting on port 8081...")
	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
