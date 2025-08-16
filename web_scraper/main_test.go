package main

import (
	"encoding/json"
	"io/ioutil"
	"testing"
)

// Model represents the structure of a model in the JSON data.

type Model struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	ContextLength int    `json:"context_length"`
}


// Response represents the top-level structure of the JSON data.

type Response struct {
	Data []Model `json:"data"`
}

func TestParseModels(t *testing.T) {
	// Read the JSON file
	data, err := ioutil.ReadFile("tests/openrouter-api.json")
	if err != nil {
		t.Fatalf("Failed to read file: %v", err)
	}

	// Unmarshal the JSON data
	var response Response
	err = json.Unmarshal(data, &response)
	if err != nil {
		t.Fatalf("Failed to unmarshal JSON: %v", err)
	}

	// Check if we have at least one model
	if len(response.Data) == 0 {
		t.Fatal("No models found in the JSON data")
	}

	// Check the first model's data
	model := response.Data[0]
	if model.ID == "" {
		t.Error("Expected model ID to be non-empty")
	}
	if model.Name == "" {
		t.Error("Expected model name to be non-empty")
	}
	if model.Description == "" {
		t.Error("Expected model description to be non-empty")
	}
	if model.ContextLength == 0 {
		t.Error("Expected model context length to be non-zero")
	}
}