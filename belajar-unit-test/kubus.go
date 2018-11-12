package main

import (
	"math"
)

// struct kubus
type Kubus struct {
	Sisi float64
}

// method Volume milik kubus
func (k Kubus) Volume() float64 {
	return math.Pow(k.Sisi, 3)
}

// method luas milik kubus
func (k Kubus) Luas() float64 {
	return math.Pow(k.Sisi, 2) * 6
}

// method keliling milik kubus
func (k Kubus) Keliling() float64 {
	return k.Sisi * 12
}
func main() {

}
