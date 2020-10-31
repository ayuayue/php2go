package array

import (
	"testing"
)
func TestArrayDiffInt(t *testing.T) {
	a := []int{1,2,3,4}
	b := []int{3,4,2,5}
	diff := DiffInt(a,b)
	expected := 1
	if diff[0] != expected {
		t.Errorf("expected %v but get %v",expected,diff)
	}
}
func TestArrayDiffInt64(t *testing.T) {
	a := []int64{1,2,3,4}
	b := []int64{3,4,2,5}
	diff := DiffInt64(a,b)
	expected := int64(1)
	if diff[0] != expected {
		t.Errorf("expected %v but get %v",expected,diff)
	}
}
func TestArrayDiffString(t *testing.T) {
	a := []string{"1","2","3","4"}
	b := []string{"3","4","2","5"}
	diff := DiffString(a,b)
	expected := "1"
	if diff[0] != expected {
		t.Errorf("expected %v but get %v",expected,diff)
	}
}