package yaslug

func findChar(s rune) string {
	charList := map[string]string{
		"а": "a", "б": "b", "в": "v", "г": "g", "д": "d",
		"е": "e", "ё": "yo", "ж": "zh", "з": "z", "и": "i",
		"й": "i", "к": "k", "л": "l", "м": "m", "н": "n",
		"о": "o", "п": "p", "р": "r", "с": "c", "т": "t",
		"у": "u", "ф": "f", "х": "x", "ц": "c", "ч": "ch",
		"ш": "sh", "щ": "sch", "ъ": "", "ы": "y", "ь": "",
		"э": "e", "ю": "yu", "я": "ya",
	}
	ch := charList[string(s)]
	if ch == "" {
		return string(s)
	}
	return ch
}

func cyrToLat(in string) string {
	out := ""
	for _, char := range in {
		out = out + findChar(char)
	}
	return out
}
