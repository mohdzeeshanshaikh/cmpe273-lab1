package shapes

import "testing"
import "fmt"

func calc_area(s Shape) float64 {
	return s.area()
}

func check(t *testing.T, lhs float64, rhs float64) {
	if lhs != rhs {
		t.Error("Expected ", lhs, " got ", rhs)
	}
}

func calc_perimeter(s Shape) float64 {
	return s.perimeter()
}

func TestShapes(t *testing.T) {
	c := Circle{0, 0, 10}
	r := Rectangle{0, 0, 100, 100}

	check(t, calc_area(c), 314.1592653589793)
	check(t, calc_area(r), 10000)

	check(t, calc_perimeter(c), 62.83185307179586)
	check(t, calc_perimeter(r), 400)

	c1 := Circle{10, 10, 100}
	fmt.Println("c1.area = ", c1.area())
	r1 := Rectangle{100, 100, 200, 100}
	fmt.Println("r1.area = ", r1.area())

	c2 := Circle{10, 10, 100}
	fmt.Println("c2.perimeter = ", c2.perimeter())
	r2 := Rectangle{100, 100, 200, 100}
	fmt.Println("r2.perimeter = ", r2.perimeter())
}
