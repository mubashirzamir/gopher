package structs

import (
	"math"
	"reflect"
	"testing"
)

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t testing.TB, shape Shape, want float64) {
		t.Helper()

		got := shape.Perimeter()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	PerimeterTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Length: 10, Width: 10}, want: 40},
		{shape: Circle{Radius: 10}, want: 2 * math.Pi * 10},
	}

	for _, test := range PerimeterTests {
		t.Run(reflect.TypeOf(test.shape).String(), func(t *testing.T) {
			checkPerimeter(t, test.shape, test.want)
		})
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()

		got := shape.Area()
		if got != want {
			t.Errorf("%#v got %g want %g", shape, got, want)
		}
	}

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{shape: Rectangle{Length: 12, Width: 6}, want: 72.0},
		{shape: Circle{Radius: 10}, want: math.Pi * math.Pow(10, 2)},
		{shape: Triangle{Base: 12, Height: 6}, want: 36.0},
	}

	for _, test := range areaTests {
		t.Run(reflect.TypeOf(test.shape).String(), func(t *testing.T) {
			checkArea(t, test.shape, test.want)
		})
	}
}
