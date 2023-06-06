package main

import (
	"fmt"
	"sort"

	"github.com/maja42/goval"
)

func main() {
	strategies := []strategygroup{
		{
			id:     1,
			name:   "level",
			weight: 0.1,
			strategies: []strategy{
				{
					id:       1,
					pid:      1,
					name:     "level",
					rule:     "level < 3 && time >100",
					weight:   0.1,
					variable: map[string]interface{}{"level": 1, "time": 50},
					delay:    2,
				},
				{
					id:       2,
					pid:      1,
					name:     "level",
					rule:     "3<level && level<5",
					weight:   0.2,
					variable: map[string]interface{}{"level": 2},
					delay:    10,
				},
			},
		},
		{
			id:     2,
			name:   "time",
			weight: 0.1,
			strategies: []strategy{
				{
					id:       1,
					pid:      2,
					name:     "time",
					rule:     "time<=7",
					weight:   0.1,
					variable: map[string]interface{}{"time": 7},
					delay:    48,
				},
			},
		},
	}
	sort.Slice(strategies, func(i, j int) bool {
		return strategies[i].weight > strategies[j].weight
	})
	eval := goval.NewEvaluator()
	delay := 0
	for _, v := range strategies {
		sort.Slice(v.strategies, func(i, j int) bool {
			return v.strategies[i].weight > v.strategies[j].weight
		})
		for _, v2 := range v.strategies {
			fmt.Println(v2.rule, v2.variable)
			r, err := eval.Evaluate(v2.rule, v2.variable, nil)
			if err != nil {
				panic(err)
			}
			fmt.Println(v2.name, r)
			if v, ok := r.(bool); !v && ok {
				delay += int(v2.delay)
			}
		}
	}
	fmt.Println(delay)
}

type strategygroup struct {
	id         int
	name       string
	weight     float32
	strategies []strategy
}

type strategy struct {
	id       int
	pid      int
	name     string
	rule     string
	weight   float32
	variable map[string]interface{}
	delay    float32
}
