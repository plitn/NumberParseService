package NumberParseService

import (
	"fmt"
	"log"
	"os"
)

type HexParserCustom struct {
}

/*
Написал конвертацию ручками, тут просто перебираем строку,
каждый символ переводим в 10систему счисления и добавляем к числу ответу
*/
func (hp *HexParserCustom) Convert(hexString string) string {
	var (
		base, decimal int64
	)
	base = 1
	hexMapping := map[rune]int64{
		'0': 0, '1': 1, '2': 2,
		'3': 3, '4': 4, '5': 5,
		'6': 6, '7': 7, '8': 8,
		'9': 9, 'A': 10, 'B': 11,
		'C': 12, 'D': 13, 'E': 14, 'F': 15,
	}
	for _, digit := range hexString {
		value, err := hexMapping[rune(digit)]
		if !err {
			log.Println("parsing error")
			os.Exit(1)
		}
		decimal += value * base
		base *= 16
	}
	return fmt.Sprintf("%d", decimal)
}
