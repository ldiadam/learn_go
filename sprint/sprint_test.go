package sprint

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFilterBySum(t *testing.T) {
	// Test case
	input := [][]int{{1, 2, 3}, {2, 3, 4}, {3, 4, 5}}
	target := 9
	expected := [][]int{{2, 3, 4}, {3, 4, 5}}

	// Run FilterBySum function
	result := FilterBySum(input, target)

	// Check if the result matches the expected output
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("FilterBySum(%v, %d) = %v; want %v", input, target, result, expected)
	}
}
func TestBulkAtoi(t *testing.T) {
	input := []string{"88", "kood", "-13"}
	expected := []int{88, 0, -13}

	result := BulkAtoi(input)

	// Check if the result matches the expected output
	if len(result) != len(expected) {
		t.Fatalf("Expected length %d but got %d", len(expected), len(result))
	} else {
		fmt.Printf("BulkAtoi(%s) : %v", input, expected)
	}
}

func TestBalanceOut(t *testing.T) {
	input := []bool{true, false, false, false}
	expected := []bool{true, false, false, false, true, true}

	result := BalanceOut(input)

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("FilterBySum(%v = %v; want %v", input, result, expected)
	}

}
