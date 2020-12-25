package examples

import (
	beforeEach "github.com/jknair0/beforeeach"
	"testing"
)

var it = beforeEach.Create(setUp, tearDown)

var calculator *Calculator

func setUp() {
	calculator = NewCalculator()
}

func tearDown() {
	calculator = nil
}

func TestCalculator_Add(t *testing.T) {
	it(func() {
		calculator.AddOffset(10)
		result := calculator.Add(5, 5)
		if result != 20 {
			t.Error("test failed")
		}
	})
}

func TestCalculator_AddOffset(t *testing.T) {
	it(func() {
		calculator.AddOffset(-20)
		result := calculator.Subtract(5, 5)
		if result != -20 {
			t.Error("test failed")
		}
	})
}
