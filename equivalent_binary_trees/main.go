package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	// Walk left->current->right to send numbers in order.
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		defer close(c1)
		Walk(t1, c1)
	}()
	go func() {
		defer close(c2)
		Walk(t2, c2)
	}()

	for {
		x1, c1open := <-c1
		x2, c2open := <-c2

		// if one closed before the other then they are different lengths, so not equal
		if c1open != c2open {
			return false
		}

		// ints are sent/received in order, so as soon as the ints in same position don't match then tree is not equal
		if x1 != x2 {
			return false
		}

		// if both closed then we're done checking and they must be equal
		if !c1open && !c2open {
			break
		}
	}

	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
	fmt.Println(Same(tree.New(2), tree.New(2)))

	// test different length trees
	t1 := &tree.Tree{
		Left:  &tree.Tree{Value: -2},
		Value: -1,
	}
	t2 := &tree.Tree{
		Left:  &tree.Tree{Value: -2},
		Value: -1,
		Right: &tree.Tree{Value: 0},
	}
	fmt.Println(Same(t1, t2))
}
