package main

import (
	"grep/cmd"
	"reflect"
	"testing"
)

func Test_SearchInFile(t *testing.T) {
	// given
	filepath := "test.txt"
	searchText := "search_string"
	want := []string{"I found the search_string in the file.", "Another line also contains the search_string"}

	// when
	output, _ := cmd.SearchInFile(searchText, filepath)

	// then
	if !reflect.DeepEqual(output, want) {
		t.Errorf("wanted:%v but got: %v", want, output)
	}
}
