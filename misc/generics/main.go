package main

import "fmt"

type number interface {
	int | int8 | int16 | int32 | int64 | float32 | float64
}

// function
func sum[v number](num []v) v {
	var res v
	for _, x := range num {
		res += x
	}
	return res
}

func show[k comparable, v any](mp map[k]v) {
	for key, value := range mp {
		fmt.Println(key, ": ", value)
	}
}

// structure
type data[T number] struct {
	a, b T
}

func (d *data[T]) area() T {
	return d.a * d.b
}

func main() {
	a := []int{4, 5, 6}
	b := []float64{4.5, 6.3, 3.25}
	fmt.Println(sum(a))
	fmt.Println(sum(b))

	mp := map[string]int32{
		"abc": 45,
		"xyz": 90,
	}
	show(mp)

	d := data[float32]{
		a: 9.8,
		b: 8.1,
	}
	fmt.Println(d.area())
}
