package sprint

func IntVsFloat(i int, f float32) string {
	if float32(i) < f {
		return "Float"
	} else if float32(i) > f {
		return "Integer"
	} else {
		return "Same"
	}
}
