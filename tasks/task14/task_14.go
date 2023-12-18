package task14

import (
	"fmt"
)

func RecognizeType(unknown interface{}) {
	switch unknown.(type) {
	case int:
		fmt.Println("int value")
	case bool:
		fmt.Println("bool value")
	case string:
		fmt.Println("string value")
	case chan int:
		fmt.Println("chan int value")
	}
}
