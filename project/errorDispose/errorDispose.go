package errorDispose

import "fmt"

func ErrorPrint(error error, errorText string) {
	if error != nil {
		fmt.Println(error, errorText)
		return
	}
}
