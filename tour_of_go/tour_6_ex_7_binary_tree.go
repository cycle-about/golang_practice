package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

/*
Exercise: Equivalent Binary Trees

There can be many different binary trees with the same sequence of values stored in it. 
For example, here are two binary trees storing the sequence 1, 1, 2, 3, 5, 8, 13. 

A function to check whether two binary trees store the same sequence is quite complex in most languages. 
We'll use Go's concurrency and channels to write a simple solution.

This example uses the tree package, which defines the type:

	type Tree struct {
	    Left  *Tree
	    Value int
	    Right *Tree
	}

/ 1. Implement the Walk function.
/ 2. Test the Walk function.
	a. The function tree.New(k) constructs a randomly-structured (but always sorted) binary tree holding the values k, 2k, 3k, ..., 10k.
	b. Create a new channel ch and kick off the walker:
		go Walk(tree.New(1), ch)
	c. Then read and print 10 values from the channel. It should be the numbers 1, 2, 3, ..., 10.
3. Implement the Same function using Walk to determine whether t1 and t2 store the same values.
4. Test the Same function.
	- Same(tree.New(1), tree.New(1)) should return true, and Same(tree.New(1), tree.New(2)) should return false.

The documentation for Tree can be found here (https://pkg.go.dev/golang.org/x/tour/tree?utm_source=godoc#Tree).

-------------------

Steps
	1. Watch video about binary tree as review -  https://www.youtube.com/watch?v=fAAZixBzIAI&t=69s
		Depth-first search: use stack
			Examples: in-order (left-root-right), pre-order (root-left-right), post-order (left-right-root)
		Breadth-first search: use queue
			Aka: level-order traversal
		Recursion: depth-first conducive to that because it uses a stack (like recursion uses, but breadth-first needing an ordered queue makes it not suited for recursion
		Trees in the video are NOT sorted
		In-order traversal of sorted tree yields the results in sorted order

		Standard algorithm for recursive in-order traversal
			
			def in_order_traversal(node):
			    if node is not None:
			        in_order_traversal(node.left)     # Traverse left subtree
			        print(node.value)                 # Visit current node
			        in_order_traversal(node.right)    # Traverse right subtree

	2. Create a tree.New and run various functions on it
	3. Start a version of walk that just adds values to a slice (array?)
*/

// Walk walks the tree t sending all values from the tree to the channel ch.
// uses in-order traversal
func Walk(t *tree.Tree, ch chan int) {
    if t != nil {
	    Walk(t.Left, ch)
	    ch <- t.Value
	    Walk(t.Right, ch)
	}
}

// Same determines whether the trees t1 and t2 contain the same values.
// func Same(t1, t2 *tree.Tree) bool

func Run_tour_6_ex_tree() {
	fmt.Println("balanced binary tree")
	ch := make(chan int, 10) // buffered channel, capacity 10
	
	go func() {
        Walk(tree.New(2), ch)
        close(ch) // Only close once, after Walk finishes
    }()

	for val := range ch {
    	fmt.Print(val, " ")
    }
}