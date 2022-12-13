package yaslug

import "strings"

// * перевести стороку в нижний регистр
// * функцией translite перевести русские символы в английский аналог
// * функцией purifyChar заменить все недопустимые символы на пробелы
// * разбить строку на слова по пробелам
// * соединить слова обратно через дефис.

func Url(src string) string {
	res := strings.ToLower(src)
	res = strings.Trim(res, "-")
	res = cyrToLat(res)
	res = strings.Map(purifyChar, res)
	words := strings.Fields(res)
	return strings.Join(words, "-")
}

func purifyChar(r rune) rune {
	const validChars string = "abcdefghijklmnopqrstuvwxyz01234567890- "
	if strings.IndexRune(validChars, r) == -1 {
		return ' '
	}
	return r
}
