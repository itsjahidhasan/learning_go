package conditions

import "fmt"

func LearnCondition (){
	var a = 10
    var b = 20

    if x,ok:=a,a<b;ok {
        fmt.Printf("%d is less than %v\n",x,ok)
    } else {
        fmt.Printf("%d is greater than or equal to %d\n", a, b)
    }
}
