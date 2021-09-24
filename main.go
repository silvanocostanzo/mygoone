package main

import (
	"fmt"

	calcv1 "github.com/silvanocostanzo/go_math/calc"
	"github.com/silvanocostanzo/go_math/geometry"
	calcv2 "github.com/silvanocostanzo/go_math/v2/calc"
)

func main() {
	v := geometry.CubeVolume(3)
	fmt.Println(v)

	s := calcv2.Add(2, 3, 4, 5)
	fmt.Println(s)

	s1 := calcv1.Add(2, 3)
	fmt.Println(s1)
}
