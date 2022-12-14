package AsciiArt

import (
	"fmt"
	"strings"
)

func AsciiArt(rawinput, style string) (string, error) {
	if err := ReadTemplates(style); err != nil {
		return "", err
	}

	input := strings.Split(rawinput, "\r\n")
	removeNewline(&input)

	var outputString string

	for _, s := range input {
		if s == "" {
			outputString = outputString + "\n"
			continue
		}

		asciiLine, err := PrintInput(s)
		if err != nil {
			return "", err
		}

		outputString = outputString + asciiLine
	}

	return outputString, nil
}

var (
	Store [128][8]string // Переменная для хранения символов из файла
)

// Выводит данную строку на консоль символами из файла
func PrintInput(s string) (result string, err error) {

	for i := 0; i < 8; i++ {
		var tmp string

		for _, r := range s {
			if r < 0 || r > 127 || Store[int(r)][0] == "" {
				err = fmt.Errorf("a character %c is not available", r)
				return "", err
			}

			tmp = tmp + Store[int(r)][i]
		}

		result = result + tmp + "\n"
	}

	return result, nil
}

// Когда в вводной строке нет слов, а только '\n' либо вообще ничего, создается лишняя пустая строка, из-за которой на консоль выводится лишшняя новая линия. Эта функция убирает эту строку.
// "\n" -> "", ""
// "\nHello" -> "", "Hello"
func removeNewline(input *[]string) {
	nowords := true
	for _, s := range *input {
		if len(s) > 0 {
			nowords = false
			break
		}
	}
	if nowords {
		*input = (*input)[1:]
	}
}
