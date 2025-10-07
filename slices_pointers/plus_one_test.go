package slices_pointers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPlusOneZero(T *testing.T) {
	input := []int{0}
	expected := []int{1}
	actual := plusOne(input)
	assert.Equal(T, expected, actual)
}

func TestPlusOneUsualNum(T *testing.T) {
	input := []int{1, 2, 5}
	expected := []int{1, 2, 6}
	actual := plusOne(input)
	assert.Equal(T, expected, actual)
}

func TestPlusOneWithLastNine(T *testing.T) {
	input := []int{1, 1, 9}
	expected := []int{1, 2, 0}
	actual := plusOne(input)
	assert.Equal(T, expected, actual)
}

func TestPlusOne999(T *testing.T) {
	input := []int{9, 9, 9}
	expected := []int{1, 0, 0, 0}
	actual := plusOne(input)
	assert.Equal(T, expected, actual)
}
