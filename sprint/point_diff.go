package sprint

type Point3 struct {
	X    float32
	Y    float32
	Text string
}

func PointDiff(p1, p2 Point3) Point3 {
	if p1.X > p2.X {
		return p1
	} else if p1.X == p2.X {
		if p1.Y > p2.Y {
			return p1
		} else {
			return p2
		}
	} else {
		return p2
	}
}
