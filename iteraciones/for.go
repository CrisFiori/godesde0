package iteraciones

import (
	"fmt"
)

func Interar(){
	
	for i := 0; i < 100; i+=5 {
		if i == 50{
			break
		}
		fmt.Println(i)
	}
}