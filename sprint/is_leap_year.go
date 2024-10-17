package sprint

func IsLeapYear(year int) bool {
	if (year%4 == 0) && (year%100 != 0 || year%400 == 0) {
		//Divisible by 4: Yes = Lapyear.
		//Divisible by 100: Yes = Not Lapyear
		//Divisible by 100: Yes but Divisible by 400: Yes = Lapyear
		return true
	} else {
		return false
	}
}
