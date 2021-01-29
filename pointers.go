package main

func pointers() {
	var x **int
	var y *int
	z := 3
	y = &z
	x = &y
}
