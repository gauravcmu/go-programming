package main

import (
"fmt"
)

var (
        stack []string 
	stackTop = -1

	graph map[string][]string
	visited map[string] bool 
)

func pop() string {
    //pop top element on stack
    top := stack[stackTop]
    stack = append(stack[:stackTop])
    stackTop -= 1
    fmt.Printf("Pop - top %v\n", top)
    return top
}

func push(element string)  {
    //push element to top of stack
    stack = append(stack, element)
    stackTop = stackTop +1 
    fmt.Printf("stack top %v\n", stackTop)
}

func main() {
	fmt.Printf("Hello World\n")
	graph = map[string][]string{"A":{"B", "C"}, "B": {"A", "C", "D"}, "C": {"A", "B", "E"}, "D":{"B"}, "E": {"B", "C"}}
	visited = map[string]bool{"A":false, "B":false, "C":false, "D":false, "E":false}
	DFS(graph, "A")
}

func DFS (graph map[string][]string, v string) {
	push(v) 

	for stackTop != -1 {
		u := pop()	
		if visited[u] != true {
			visited[u] = true 
			fmt.Printf("-----Visiting %v ------\n", u)
			for _, value := range graph[u] {
				if visited[value] != true {
				fmt.Printf("pushing %v\n", value)
				push(value)
				}
			}
		}
	}
}


