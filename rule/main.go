package main

import (
	"fmt"

	"github.com/maja42/goval"
)

func main() {
	eval := goval.NewEvaluator()
	result, err := eval.Evaluate(`42 > 21`, nil, nil) // Returns <true, nil>
	fmt.Println(result, err)
	m := map[string]interface{}{
		"price":   2000,
		"contain": "1",
	}
	result, err = eval.Evaluate(`price > 200 && contain in ["1", "2", "3"]`, m, nil)
	fmt.Println(result, err)
}
