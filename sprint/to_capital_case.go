package sprint

func ToCapitalCase(s string) string {
	var result string

	n := len(s)

	//first letter will alwyas capital, so its capital true
	capital := true

	for i := 0; i < n; i++ {
		char := s[i]

		//if char s[i] is lowecase
		if char >= 'a' && char <= 'z' {
			//check wether this char should be capital or not
			if capital {
				//if capital, -32 in ascii
				result += string(char - 32)
			} else {
				//if not capital, leave as is
				result += string(char)
			}
			//since this is not symbol or anything, the next char should be not capital
			capital = false

			//this is just the reverse
		} else if char >= 'A' && char <= 'Z' {
			if capital {
				result += string(char)
			} else {
				result += string(char + 32)
			}
			//same here, since this is not symbol or anything, the next char should be not capital
			capital = false
		} else if char >= '0' && char <= '9' {
			result += string(char)
			//same here, since this is not symbol or anything, the next char should be not capital
			capital = false
		} else {
			result += string(char)
			// since this is gonna be symbol (not a-z A-Z) the next char has to be capital
			capital = true
		}
	}

	return result
}
