package Array

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T){
	t.Run("Array",func(t *testing.T){
		numberArray := [5]int{1,2,3,4,6}
		got :=ArraySum(numberArray)
		want := 16
		if got != want {
			t.Errorf("want %d ,got %d",want,got)
		}
	})
	t.Run("Slice",func(t *testing.T) {
		numberSlice := []int{1,2,3,4,6}
	
		got := SliceSum(numberSlice)
		want := 16

		if got != want {
			t.Errorf("want %d , got %d", want,got)
		}
	})
	t.Run("SumAll",func(t *testing.T) {
		got := sumAll([]int{1,9},[]int{2,10})
		want := []int{10,12}
		if !reflect.DeepEqual(got,want){
			t.Errorf("want %v , got %v",want,got)
		}
	})
	t.Run("SumAllTails",func(t *testing.T) {
		got := sumAllTails([]int{1,9},[]int{2,10})
		want := []int{9,10}
		if !reflect.DeepEqual(got,want){
			t.Errorf("want %v , got %v",want,got)
		}
	})

	
	
}