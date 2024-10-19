package encrypt

func Nimbus(str string) string {
	encrypstr := ""
	for _, c := range str {
		asciiCode := int(c)
		character := string(asciiCode + 3)
		encrypstr += character
	}
	return encrypstr
}
