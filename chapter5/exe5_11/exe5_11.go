// 来源：http://lovenoodles.cn/2021/12/16/221
package main

import (
	"fmt"
	"os"
	"strings"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},
	// uncomment line below to introduce a dependency cycle.
	//"intro to programming": {"data structures"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	order, err := topoSort(prereqs)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	for i, course := range order {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func index(s string, slice []string) (int, error) {
	for i, v := range slice {
		if s == v {
			return i, nil
		}
	}
	return 0, fmt.Errorf("not found")
}

func topoSort(m map[string][]string) (order []string, err error) {
	resolved := make(map[string]bool)
	var visitAll func([]string, []string)

	visitAll = func(items []string, parents []string) {
		for _, v := range items {
			vResolved, seen := resolved[v] //当前课程已经存在且resolved标识为false，说明当前课程在本次深度搜索中出现重复，即图中出现环
			if !vResolved && seen {
				start, _ := index(v, parents) // Ignore error since v has to be in parents.
				err = fmt.Errorf("cycle: %s", strings.Join(append(parents[start:], v), " -> "))
			}
			if !seen {
				resolved[v] = false
				visitAll(m[v], append(parents, v))
				resolved[v] = true
				order = append(order, v)
			}
		}
	}

	for k := range m {
		if err != nil {
			return nil, err
		}
		visitAll([]string{k}, nil)
	}
	return order, nil
}
