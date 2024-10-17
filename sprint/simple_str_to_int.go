package sprint

func SimpleStrToInt(s string) int {
	result := 0 //Initialize result
	i := 0      //init i
	n := len(s) // count length of s (string)

	//looping every char inside string
	for i < n {
		//access ascii code for that char
		char := s[i]

		//check is char is not inside number range
		if char < '0' || char > '9' {
			break
		}

		//converting to string by substracting char to 0 (ascii 49)
		digit := int(char - '0')
		result = result*10 + digit
		i++
	}
	return result
}
