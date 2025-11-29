package common

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	tests := []struct {
		name        string
		target      int
		searchSpace []int
		wantIndex   int
		wantErr     bool
	}{
		{
			name:        "target found in middle",
			target:      5,
			searchSpace: []int{1, 3, 5, 7, 9},
			wantIndex:   2,
			wantErr:     false,
		},
		{
			name:        "target found at beginning",
			target:      1,
			searchSpace: []int{1, 3, 5, 7, 9},
			wantIndex:   0,
			wantErr:     false,
		},
		{
			name:        "target found at end",
			target:      9,
			searchSpace: []int{1, 3, 5, 7, 9},
			wantIndex:   4,
			wantErr:     false,
		},
		{
			name:        "target not found",
			target:      6,
			searchSpace: []int{1, 3, 5, 7, 9},
			wantIndex:   0,
			wantErr:     true,
		},
		{
			name:        "empty search space",
			target:      5,
			searchSpace: []int{},
			wantIndex:   0,
			wantErr:     true,
		},
		{
			name:        "single element found",
			target:      5,
			searchSpace: []int{5},
			wantIndex:   0,
			wantErr:     false,
		},
		{
			name:        "single element not found",
			target:      3,
			searchSpace: []int{5},
			wantIndex:   0,
			wantErr:     true,
		},
		{
			name:        "duplicate elements - returns first occurrence",
			target:      5,
			searchSpace: []int{1, 3, 5, 5, 5, 7, 9},
			wantIndex:   2,
			wantErr:     false,
		},
		{
			name:        "all elements same - returns first",
			target:      5,
			searchSpace: []int{5, 5, 5, 5, 5},
			wantIndex:   0,
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIndex, err := BinarySearch(tt.target, tt.searchSpace)
			if (err != nil) != tt.wantErr {
				t.Errorf("BinarySearch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotIndex != tt.wantIndex {
				t.Errorf("BinarySearch() = %v, want %v", gotIndex, tt.wantIndex)
			}
		})
	}
}

func TestBinarySearchStrings(t *testing.T) {
	tests := []struct {
		name        string
		target      string
		searchSpace []string
		wantIndex   int
		wantErr     bool
	}{
		{
			name:        "string found",
			target:      "cat",
			searchSpace: []string{"apple", "cat", "dog", "elephant"},
			wantIndex:   1,
			wantErr:     false,
		},
		{
			name:        "string not found",
			target:      "bird",
			searchSpace: []string{"apple", "cat", "dog", "elephant"},
			wantIndex:   0,
			wantErr:     true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotIndex, err := BinarySearch(tt.target, tt.searchSpace)
			if (err != nil) != tt.wantErr {
				t.Errorf("BinarySearch() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotIndex != tt.wantIndex {
				t.Errorf("BinarySearch() = %v, want %v", gotIndex, tt.wantIndex)
			}
		})
	}
}
