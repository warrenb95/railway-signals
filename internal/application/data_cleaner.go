package application

import (
	"bytes"
	"encoding/json"
	"io"
	"math"
)

func CleanJSON(input io.Reader) ([]byte, error) {
	allBytes, err := io.ReadAll(input)
	if err != nil {
		return nil, err
	}
	allBytes = bytes.ReplaceAll(allBytes, []byte("NaN"), []byte("null"))

	var data []any
	if err := json.Unmarshal(allBytes, &data); err != nil {
		return nil, err
	}

	// Function to traverse and clean NaN values
	cleanSlice(data)

	// Encode back to JSON
	output, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	return output, nil
}

func cleanMap(m map[string]any) {
	for key, value := range m {
		switch v := value.(type) {
		case float64:
			if math.IsNaN(v) {
				m[key] = nil
			}
		case map[string]any:
			cleanMap(v)
		case []any:
			cleanSlice(v)
		}
	}
}

func cleanSlice(s []any) {
	for i, value := range s {
		switch v := value.(type) {
		case float64:
			if math.IsNaN(v) {
				s[i] = nil
			}
		case map[string]any:
			cleanMap(v)
		case []any:
			cleanSlice(v)
		}
	}
}
