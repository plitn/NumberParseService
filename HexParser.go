package NumberParseService

type HexConverter interface {
	Convert(hexString string) (decString string)
}
