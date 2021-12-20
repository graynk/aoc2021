package main

import "fmt"

func iterate(enhancement string, image Image, steps int) int {
	first := rune(enhancement[0])
	last := rune(enhancement[len(enhancement)-1])
	infinity := '.'
	for i := 0; i < steps; i++ {
		image = image.Enhance(enhancement, infinity)
		if infinity == '.' {
			infinity = first
		} else {
			infinity = last
		}
	}
	return image.CountLit()
}

func main() {
	enhancement, image := FromFile("./day20/input.txt")
	fmt.Println(iterate(enhancement, image, 2))
	fmt.Println(iterate(enhancement, image, 50))
}
