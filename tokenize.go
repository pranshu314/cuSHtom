package main

func Split_line(text string, tokenizeon []rune) []string {
	resultSet := []string{}
	textAsRune := []rune(text)
	i := 0

	for len(textAsRune) > 0 {
		r := textAsRune[i]

		if RuneIndexOf(tokenizeon, r) > -1 {
			setItem := textAsRune[:i]
			resultSet = append(resultSet, string(removePad(setItem)))
			// resultSet = append(resultSet, string(textAsRune[i:i+1]))

			textAsRune = textAsRune[i+1:]
			i = 0
		}

		i++
	}

	return resultSet
}

func RuneIndexOf(r []rune, el rune) int {
	for i, e := range r {
		if el == e {
			return i
		}
	}

	return -1
}

func removePad(r []rune) []rune {
	if len(r) > 0 {
		if r[0] == ' ' {
			r = r[1:]
		}

		if r[len(r)-1] == ' ' {
			r = r[:len(r)-2]
		}
	}

	return r
}
