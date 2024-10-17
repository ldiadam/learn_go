package sprint

type Circle struct {
	Radius    float32
	Diameter  float32
	Area      float32
	Perimeter float32
}

func GetCircle(r float32) Circle {
	radius := r
	diameter := radius * 2
	area := 3.14 * radius * radius
	perimeter := 2 * 3.14 * radius
	result := Circle{Radius: radius, Diameter: diameter, Area: area, Perimeter: perimeter}
	return result
}
