package interation

import "fmt"

func Repeat(character string, repeatTimes int) string {
	var repeated string
	for i := 0; i < repeatTimes; i++ {
		repeated += character + fmt.Sprintf("%d", i)
	}
	return repeated
}

func main() {
	word := "teste"
	repeatTime := 100
	fmt.Println(Repeat(word, repeatTime))
}
