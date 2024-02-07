package validpalindrome

import "strings"

func isPalindromeWithGo122(s string) bool {
	var (
		zero byte = 48
		nine byte = 57
		A    byte = 65
		Z    byte = 90
		a    byte = 97
		z    byte = 122 
	)

	var builder strings.Builder
	
	for i := range len(s) {
			b := s[i]
			if zero <= b && b <= nine {
				builder.WriteByte(b)
				continue
			}
			
			if A <= b && b <= Z {
				builder.WriteByte(b)
				continue
			}

			if a <= b && b <= z {
				builder.WriteByte(b)
			}
	}

	str := builder.String()

	size := len(str)
	for i := range size/2 {
		if str[i] != str[size-i-1] {
			return false
		} 
	}

	return true
}

func isPalindromeSecond(s string) bool {
	var (
		zero byte = 48
		nine byte = 57
		A    byte = 65
		Z    byte = 90
		a    byte = 97
		z    byte = 122

        upperLowerDiff byte = 32 
	)

	var builder strings.Builder
	
	for i:=0; i<len(s); i++ {
			b := s[i]
			if zero <= b && b <= nine {
				builder.WriteByte(b)
				continue
			}
			
			if A <= b && b <= Z {
				builder.WriteByte(b+upperLowerDiff)
				continue
			}

			if a <= b && b <= z {
				builder.WriteByte(b)
			}
	}

	str := builder.String()

	size := len(str)
    for i:=0; i < size/2; i++ {
        if str[i] != str[size-i-1] {
			return false
		} 
	}

	return true
}

// isPalindromeOne is designed to run very fast (Runtime: 100% 0ms, Memory: 51.49% 3.15MB)
func isPalindromeOne(s string) bool {
	var upperLowerDiff byte = 32 
	var str []byte
	
	for i:=0; i<len(s); i++ {
			b := s[i]
			if ('0' <= b && b <= '9') || ('a' <= b && b <= 'z') {
				str = append(str, b)
				continue
			}
			
			if 'A' <= b && b <= 'Z' {
				str = append(str, b+upperLowerDiff)
				continue
			}
	}

	size := len(str)
    for i:=0; i < size/2; i++ {
        if str[i] != str[size-i-1] {
			return false
		} 
	}

	return true
}

// isPalindrome is designed to be memory efficient (Runtime: 64.28% 4ms, Memory: 92.17% 2.65MB)
func isPalindrome(s string) bool {
    i, j := 0, len(s)-1
	for i <= j {
		for i < len(s) && isNonAlphanumeric(s[i]) {
			i++
		} 
        if i >= len(s) {
            return true
        }

		for j >= 0 && isNonAlphanumeric(s[j]) {
			j--
		}

        if j < 0 {
            return true
        }

		if toLower(s[i]) != toLower(s[j]) {
			return false
		}

		i++
        j--
	}

	return true
}

func isNonAlphanumeric(b byte) bool {
	if b < '0' || ('9' < b && b < 'A') ||  ('Z' < b && b < 'a') || 'z' < b {
		return true		
	}
	return false
}


func toLower(b byte) byte {
	if 'A' <= b && b <= 'Z' {
		var diff byte = 32
		return b + diff
	}

	return b
}