package main

import (
    "fmt"
    "testing"
)

func TestParseInput(t *testing.T) {
    colony, err := ParseInput("input.txt")
    if err != nil {
        t.Errorf("ParseInput failed: %v", err)
    }
    fmt.Println("Colony parsed:", colony)
}
