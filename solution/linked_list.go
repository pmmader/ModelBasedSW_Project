package main

import "fmt"

// Example linked list
type node[T any] struct {
	val  T
	next *node[T]
}

func mkNode[T any](v T) *node[T] {
	return &node[T]{val: v, next: nil}
}

func insert[T any](n *node[T], v T) {
	n.next = mkNode[T](v)
}

// n > 0
func mkList[T any](n int, v T) *node[T] {

	head := mkNode(v)
	current := head

	for i := 0; i < n-1; i++ {
		insert(current, v)
		current = current.next

	}
	return head

}

func len[T any](n *node[T]) int {
	i := 0

	for n != nil {
		i++
		n = n.next

	}

	return i

}

func testNode() {
	n1 := mkList[int](10, 1)
	fmt.Printf("\n%d", len(n1))

	n2 := mkList[bool](10, true)
	fmt.Printf("\n%d", len(n2))

}

// Example linked list translated by Monomorphization
type node_mono_int struct {
	val  int
	next *node_mono
}

func mkNode_mono_int(v int) *node_mono {
	return &node_mono_int{val: v, next: nil}
}

func insert_mono_int(n *node_mono_int, v int) {
	n.next = mkNode_mono_int(v)
}

// n > 0
func mkList_mono_int(n int, v int) *node_mono_int {

	head := mkNode_mono_int(v)
	current := head

	for i := 0; i < n-1; i++ {
		insert_mono_int(current, v)
		current = current.next

	}
	return head

}

func len_mono_int(n *node_mono_int) int {
	i := 0

	for n != nil {
		i++
		n = n.next

	}

	return i

}

func testNode_mono_int() {
	n1 := mkList_mono_int(10, 1)
	fmt.Printf("\n%d", len_mono_int(n1))

}



type node_mono_bool struct {
	val  bool
	next *node_mono
}

func mkNode_mono_bool(v int) *node_mono_bool {
	return &node_mono_boolval: v, next: nil}
}

func insert_mono_bool(n *node_mono_bool, v bool) {
	n.next = mkNode_mono_bool(v)
}

// n > 0
func mkList_mono_bool(n int, v bool) *node_mono_bool {

	head := mkNode_mono_bool(v)
	current := head

	for i := 0; i < n-1; i++ {
		insert_mono_bool(current, v)
		current = current.next

	}
	return head

}

func len_mono_bool(n *node_mono_bool) int {
	i := 0

	for n != nil {
		i++
		n = n.next

	}

	return i

}

func testNode_mono_bool() {
	n1 := mkList_mono_bool(10, 1)
	fmt.Printf("\n%d", len_mono_bool(n1))




// with Example linked list Generic translation
type node_gen struct {
	val  interface{}
	next *node_gen
}

func mkNode_gen(v interface{}) *node_gen {
	return &node_gen{val: v, next: nil}
}

func insert_gen(n *node_gen, v interface{}) {
	n.next = mkNode_gen(v)
}

// n > 0
func mkList_gen(n int, v interface{}) *node_gen {

	head := mkNode_gen(v)
	current := head

	for i := 0; i < n-1; i++ {
		insert_gen(current, v)
		current = current.next

	}
	return head

}

func len_gen(n *node_gen) int {
	i := 0

	for n != nil {
		i++
		n = n.next

	}

	return i

}

func testNode_gen() {
	n1 := mkList_gen(10, 1)
	fmt.Printf("\n%d", len_gen(n1))

	n2 := mkList_gen(10, true)
	fmt.Printf("\n%d", len_gen(n2))

}

// Example summing up numeric values
func sum[T int | float32](xs []T) T {
	var x T
	x = 0
	for _, v := range xs {

		x = x + v
	}

	return x

}

// Example summing up numeric values Monorization
func sum_mono_int(xs []int) int {
	var x int
	x = 0
	for _, v := range xs {

		x = x + v
	}

	return x

}



func sum_mono_float32(xs []float32) float32 {
	var x float32
	x = 0
	for _, v := range xs {

		x = x + v
	}

	return x
}


/*/Example summing up numeric values Generics
func sum_gen(xs []interface{}) interface{} {
	var x interface{}
	x = 0
	for _, v := range xs {

		x = x + v
	}

	return x

}*/

// Example swap
func swap[T any](x *T, y *T) {
	tmp := *x
	*x = *y
	*y = tmp
}

// Example swap Monorization
func swap_mono(x *int, y *int) {
	tmp := *x
	*x = *y
	*y = tmp
}

// Example swap Generics
func swap_gen(x *interface{}, y *interface{}) {
	tmp := *x
	*x = *y
	*y = tmp
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	a := 3
	b := 5
	fmt.Printf("Example Implementation")
	testNode()
	fmt.Printf("\nMonomorphization, only int because go does not allow functions with the same name but different parameters")
	//The duplication problem could be solved by changing the names of the functions, but that would somewhat defeat the purpuse of generic methodes. To circumvent that one would need to write a mapper function, or something similar.
	testNode_mono()
	fmt.Printf("\nGenerics")
	testNode_gen()
	fmt.Printf("\nsum example")
	fmt.Printf("\n%d", sum(nums))
	fmt.Printf("\nsum mono")
	fmt.Printf("\n%d", sum_mono(nums))
	fmt.Printf("\nsum generics is not possible, because interface{} does not have a definitaion for the +-operator")
	fmt.Printf("\nBefore swaps a=%d, b=%d", a, b)
	fmt.Printf("\nswap example")
	swap(&a, &b)
	fmt.Printf("\na=%d, b=%d", a, b)
	fmt.Printf("\nswap mono")
	swap_mono(&a, &b)
	fmt.Printf("\na=%d, b=%d", a, b)
	fmt.Printf("\nswap gen not possible, because *int does not implement *interface{}")
	/*var c *interface{} = &a;
	swap_gen(c, d)
	fmt.Printf("\na=%d, b=%d", a, b)*/
}
