package main

import (
"fmt"
)

var (
        stack []string 
	stackTop = -1
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
	push("A")
	push("B")
	push("C")
	pop()
	pop()
	pop()
}
