package decrypt

func Nimbus(str string) string {
	decryptstr := ""
	for _, c := range str {
		asciicode := int(c)
		character := string(asciicode - 3)
		decryptstr += character
	}
	return decryptstr

}
