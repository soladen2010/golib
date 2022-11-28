package strs

// replace multi spaces to one space
func TrimMultiSpace(s string) string {
	for Contains(s, "  ") {
		s = ReplaceAll(s, "  ", " ")
	}
	return s
}

// SplitLast by xp, get the last string of the splitted slice
func SplitLast(s, sep string) string {
	strList := genRightSplit(s, sep, 0, 2)
	length := len(strList)
	if length > 0 {
		return strList[length-1]
	} else {
		return ""
	}
}

// SplitFirst by xp, get the fisrt string of the splitted slice
func SplitFirst(s, sep string) string {
	return genSplit(s, sep, 0, 2)[0]
}

// Reverse by xp, reverse the string
func Reverse(s string) string {
	length := len(s)
	if length > 1 {
		t := make([]byte, length)
		for i := 0; i < length; i++ {
			t[length-i-1] = s[i]
		}
		return string(t)
	} else {
		return s
	}
}

// Generic right split: splits after each instance of sep from right,
// including sepSave bytes of sep in the subarrays.
func genRightSplit(s, sep string, sepSave, n int) []string {
	if n == 0 {
		return nil
	}
	if sep == "" {
		return explode(s, n)
	}
	if n < 0 {
		n = Count(s, sep) + 1
	}

	a := make([]string, n)
	i := n - 1
	for i > 0 {
		m := LastIndex(s, sep)
		if m < 0 {
			break
		}
		a[i] = s[m+len(sep)-sepSave:]
		s = s[:m]
		i--
	}
	a[i] = s
	return a[i:]
}

// RsplitN slices s into substrings separated by sep from right and returns a slice of
// the substrings between those separators.
//
// The count determines the number of substrings to return:
//   n > 0: at most n substrings; the first substring will be the unsplit remainder.
//   n == 0: the result is nil (zero substrings)
//   n < 0: all substrings
func RsplitN(s, sep string, n int) []string { return genRightSplit(s, sep, 0, n) }
