package utils

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// ExportData represents the data structure for export
type ExportData struct {
	Query     string      `json:"query"`
	Type      string      `json:"type"`
	Timestamp string      `json:"timestamp"`
	Count     int         `json:"count"`
	Results   interface{} `json:"results"`
}

// ExportToFile exports data to the specified format
func ExportToFile(data interface{}, query, searchType, format, filename string) error {
	// Show export progress
	spinner := NewSpinner("Exporting results")
	spinner.Start()
	defer spinner.Stop()

	if filename == "" {
		// Generate default filename
		timestamp := time.Now().Format("2006-01-02_15-04-05")
		filename = fmt.Sprintf("%s_%s_%s.%s", searchType, strings.ReplaceAll(query, " ", "_"), timestamp, format)
	}

	// Ensure file extension matches format
	if !strings.HasSuffix(strings.ToLower(filename), "."+format) {
		filename += "." + format
	}

	// Create export data structure
	exportData := ExportData{
		Query:     query,
		Type:      searchType,
		Timestamp: time.Now().Format(time.RFC3339),
		Results:   data,
	}

	// Get count based on data type
	switch v := data.(type) {
	case []interface{}:
		exportData.Count = len(v)
	default:
		exportData.Count = 0
	}

	var err error
	switch strings.ToLower(format) {
	case "json":
		err = exportToJSON(exportData, filename)
	case "csv":
		err = exportToCSV(data, query, searchType, filename)
	case "txt":
		err = exportToTXT(data, query, searchType, filename)
	default:
		err = fmt.Errorf("unsupported format: %s", format)
	}

	if err != nil {
		return WrapInternalError("export data", err)
	}

	PrintSuccess(fmt.Sprintf("Results exported to '%s' successfully", filename))
	return nil
}

// exportToJSON exports data to JSON format
func exportToJSON(data ExportData, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

// exportToCSV exports data to CSV format
func exportToCSV(data interface{}, query, searchType, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header
	header := []string{"Title", "URL", "Description"}
	if err := writer.Write(header); err != nil {
		return fmt.Errorf("failed to write header: %w", err)
	}

	// Write data based on type
	switch v := data.(type) {
	case []interface{}:
		for _, item := range v {
			if record, ok := item.(map[string]interface{}); ok {
				row := []string{
					getStringValue(record, "Title"),
					getStringValue(record, "URL"),
					getStringValue(record, "Description"),
				}
				if err := writer.Write(row); err != nil {
					return fmt.Errorf("failed to write row: %w", err)
				}
			}
		}
	}

	return nil
}

// exportToTXT exports data to plain text format
func exportToTXT(data interface{}, query, searchType, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer file.Close()

	// Write header
	fmt.Fprintf(file, "Search Results\n")
	fmt.Fprintf(file, "==============\n")
	fmt.Fprintf(file, "Query: %s\n", query)
	fmt.Fprintf(file, "Type: %s\n", searchType)
	fmt.Fprintf(file, "Timestamp: %s\n", time.Now().Format(time.RFC3339))
	fmt.Fprintf(file, "\n")

	// Write data
	switch v := data.(type) {
	case []interface{}:
		for i, item := range v {
			if record, ok := item.(map[string]interface{}); ok {
				fmt.Fprintf(file, "%d. %s\n", i+1, getStringValue(record, "Title"))
				if url := getStringValue(record, "URL"); url != "" {
					fmt.Fprintf(file, "   URL: %s\n", url)
				}
				if desc := getStringValue(record, "Description"); desc != "" {
					fmt.Fprintf(file, "   Description: %s\n", desc)
				}
				fmt.Fprintf(file, "%s\n", strings.Repeat("-", 40))
			}
		}
	}

	return nil
}

// getStringValue safely extracts string value from map
func getStringValue(record map[string]interface{}, key string) string {
	if val, ok := record[key]; ok {
		if str, ok := val.(string); ok {
			return str
		}
	}
	return ""
}

// ConvertToInterfaceSlice converts any slice to []interface{}
func ConvertToInterfaceSlice(slice interface{}) []interface{} {
	switch v := slice.(type) {
	case []interface{}:
		return v
	default:
		// For now, return empty slice for unsupported types
		// This can be extended to handle specific types
		return []interface{}{}
	}
}
