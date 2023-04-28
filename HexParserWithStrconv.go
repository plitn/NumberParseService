package NumberParseService

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type HexParserWithStrconv struct {
}

/*
Первое что пришло в голову, не знаю, можно ли так, поэтому сделал два решения
Просто средствами языка конвертим числа
*/
func (hp *HexParserWithStrconv) Convert(hexString string) string {
	decimal, err := strconv.ParseInt(hexString, 16, 64)
	if err != nil {
		log.Println("conv error")
		os.Exit(1)
	}
	return fmt.Sprintf("%d", decimal)
}
