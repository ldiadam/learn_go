package sprint

func GenRange(min, max int) []int {
	result := []int{}
	for i := max; i < max; i++ {
		result = append(result, i)
	}

	return result
}

func RemoveEl(arr []float64, from, to int) []float64 {
	result := []float64{}

	for index, value := range arr {
		if index >= from || index < to {
			result = append(result, value)
		}
	}

	return result
}

func BalOut(arr []bool) []bool {
	result := arr

	ct := 0
	cf := 0
	need := 0
	var insert bool

	for _, value := range arr {
		if value == true {
			ct++
		} else {
			cf++
		}
	}

	if ct > cf {
		need = ct - cf
		insert = false
	} else {
		need = cf - ct
		insert = true
	}

	for i := 0; i < need; i++ {
		result = append(result, insert)
	}

	return result
}

func FilBySum(arr [][]int, limit int) [][]int {
	var result [][]int
	var temp int
	for _, parent := range arr {
		for _, i := range parent {
			temp = temp + i
		}
		if temp >= limit {
			result = append(result, parent)
		}
		temp = 0
	}

	return result
}

func SortInt(table []int) []int {
	for i := 0; i < len(table)-1; i++ {
		for j := 0; j < len(table)-i-1; j++ {
			if table[j] > table[j+1] {
				table[j], table[j+1] = table[j+1], table[j]
			}
		}
	}
	return table
}

func Str_int(s string) int {
	if len(s) <= 0 {
		return 0
	}

	i := 0
	neg := false
	result := 0

	if s[i] == '-' {
		neg = true
		i++
	}

	for ; i < len(s); i++ {
		char := s[i]

		if char < '0' || char > '9' {
			return 0
		}
		digit := int(s[i] - '0')
		result = result*10 + digit
	}

	if neg {
		result = result * -1
	}

	return result
}

func BulkA(arr []string) []int {
	result := []int{}
	for _, v := range arr {
		convert := Str_int(v)
		result = append(result, convert)
	}

	return result
}

func Comb_n(n int) []string {
	digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	var result []string

	var combRecursive func(current string, startIdx int)
	combRecursive = func(current string, startIdx int) {
		if len(current) == n {
			result = append(result, current)
			return
		}

		for i := startIdx; i < len(digits); i++ {
			combRecursive(current+digits[i], i+1)
		}
	}

	combRecursive("", 0)
	return result
}

func AnotherCombn(n int) []string {
	digits := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	var result []string

	var combR func(currentString string, startIdx int)
	combR = func(currentString string, startIdx int) {
		if len(currentString) == n {
			result = append(result, currentString)
		}

		for i := startIdx; i < len(digits); i++ {
			combR(currentString+digits[i], i+1)
		}
	}

	combR("", 0)

	return result
}

func Str_len(s string) []int {
	bytecount := len(s)
	runecount := len([]rune(s))

	return []int{bytecount, runecount}
}

func Str_rev(s string) string {
	var result string
	n := len(s)

	if n <= 0 {
		return ""
	}

	for i := 0; i < n; i++ {
		result = string(s[i]) + result
	}

	return result
}

func Lower(s string) bool {
	if len(s) <= 0 {
		return true
	}

	for i := 0; i < len(s); i++ {
		if s[i] < 'a' || s[i] > 'z' {
			return false
		}
	}

	return true
}

func Numeric(s string) bool {
	if len(s) <= 0 {
		return true
	}

	for i := 0; i < len(s); i++ {
		char := s[i]

		if char < '0' || char > '9' {
			return false
		}
	}

	return true
}

func ToUp(s string) string {
	n := len(s)
	var result string

	for i := 0; i < n; i++ {
		char := s[i]
		if char >= 'a' && char <= 'z' {
			result = result + string(char-32)
		} else {
			result = result + string(char)
		}
	}

	return result
}

func ToCap(s string) string {
	var result string

	toCaps := true

	for i := 0; i < len(s); i++ {
		char := s[i]

		newchar := s[i]

		if char >= 'a' && char <= 'z' {
			if toCaps {
				newchar = char - 32
			}
			toCaps = false
		} else if char > 'A' && char < 'Z' {
			if !toCaps {
				newchar = char + 32
			}
			toCaps = false
		} else {
			toCaps = true
		}

		result = result + string(newchar)

	}

	return result
}

func Str_conc(strs []string, sep string) string {
	var result string
	for index, v := range strs {
		if index < len(strs)-1 {
			result += v + sep
		} else {
			result += v
		}
	}

	return result
}

func Split_space(s string) []string {
	var result []string
	var tempword string
	for i := 0; i < len(s); i++ {
		if s[i] != ' ' {
			tempword += string(s[i])
		} else {
			result = append(result, tempword)
			tempword = ""
		}
	}

	result = append(result, tempword)

	return result
}

func Str_split(s, sep string) []string {
	var result []string
	var tempword string
	n := len(sep)

	for i := 0; i < len(s); {
		if i <= len(s)-n && s[i:i+n] == sep {
			result = append(result, tempword)
			tempword = ""
			i += n
		} else {
			tempword += string(s[i])
			i++
		}
	}

	result = append(result, tempword)

	return result
}

func Substr(s string, toFind string) int {
	n := len(toFind)
	o := len(s)

	for i := 0; i < o; i++ {
		if i <= o-n && s[i:i+n] == toFind {
			return i
		}
	}

	return -1
}

func Str_comp(a, b string) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] < b[i] {
			return -1
		}

		if a[i] > b[i] {
			return 1
		}
	}

	if len(a) < len(b) {
		return -1
	} else if len(a) > len(b) {
		return 1
	}

	return 0
}

func Numb2Base(a int, base string) string {

	unique := unq(base)
	if !unique {
		return "NV"
	}

	var result string
	n := len(base)
	for a > 0 {
		digit := a % n
		result = string(base[digit]) + result
		a = a / n
	}

	return result
}

func unq(input string) bool {
	var arr [128]bool
	for _, index := range input {
		if arr[index] {
			return false
		}

		arr[index] = true
	}

	return true
}

type Circ struct {
	Radius    float32
	Diameter  float32
	Area      float32
	Perimeter float32
}

func Circl(r float32) Circ {
	d := 2 * r
	a := 3.14 * r * r
	p := 2 * 3.14 * r

	return Circ{Radius: r, Diameter: d, Area: a, Perimeter: p}
}

func Bubb(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}

func Str2Int(s string) int {
	var result int
	n := len(s)

	for i := 0; i < n; i++ {
		char := s[i]

		if char >= '0' || char <= '9' {
			result = result * 10
			result += int(char - '0')
		}
	}

	return result
}
