package main

import "testing"

func TestGetAPIKey(t *testing.T) {
    tests := []struct {
        userID   int
        expected string
    }{
        {1, "APIKEY1234"},
        {2, "UNKNOWN"},
    }

    for _, test := range tests {
        result := GetAPIKey(test.userID)
        if result != test.expected {
            t.Errorf("GetAPIKey