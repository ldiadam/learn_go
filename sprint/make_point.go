package sprint

type Point2 struct {
	X    float32
	Y    float32
	Text string
}

func MakePoint(x, y float32, text string) Point2 {
	return Point2{X: x, Y: y, Text: text}
}
