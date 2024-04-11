package main

import (
	"fmt"
	"math"
)

func main() {
	b, l := 4.0, 4.0
	for i := 1; i <= 16; i++ {
		z := float64(i)
		//Ir := Newmark1935(b, l, z)
		_ = Boussinesq1883(b, l, z)
		//fmt.Printf("Newmark1935 z:%v I:%v\n", z, Ir)
	}
	_ = Boussinesq1883(b, l, 0.0000001)

}

func Boussinesq1883(b, l, z float64) float64 {
	//from Terzaghi Theoretical Soil mechnaics 1943
	m := b / z
	n := l / z

	m2 := m * m
	n2 := n * n
	mn2 := m * m * n * n
	Dmn2 := 2.0 * m * n

	part0 := math.Sqrt(m2 + n2 + 1.0)

	part1 := (Dmn2 * part0) / (m2 + n2 + mn2 + 1.0)
	part2 := (m2 + n2 + 2.0) / (m2 + n2 + 1.0)

	part3 := (Dmn2 * part0) / (m2 + n2 + 1.0 - mn2)
	part4 := math.Atan(part3)

	enclosedPart := (part1 * part2) + part4
	I_sigma := enclosedPart / (4.0 * math.Pi)
	fmt.Printf("Boussinesq m:n %.3f:%.3f   Is:%.5f\n", m, n, I_sigma)
	return I_sigma
}

func Newmark1935(B, L, Z float64) float64 {
	b2 := B * B
	l2 := L * L
	z2 := Z * Z
	bl2 := b2 * l2
	particle0 := B * L * Z
	particle1 := (B * B) + (L * L) + (Z * Z)
	particle2 := (B * B) + (L * L) + (2 * Z * Z)
	particle3 := math.Sqrt(particle1)

	part1 := 2 * (particle0 * particle2) / (particle3 * (bl2 + (z2 * particle1)))
	//fmt.Printf("Newmark1935 part1 :%v\n", part1)
	part2 := math.Asin(2 * (particle0 * particle3) / ((z2 * particle1) + bl2))
	//fmt.Printf("Newmark1935 part2 :%v\n", part2)
	//fmt.Printf("Newmark1935 InvPi :%v\n", Inv(math.Pi))
	m := part1 + part2
	//fmt.Printf("Newmark1935 m :%v\n", m)
	result := Inv(math.Pi) * m * 0.25
	//fmt.Printf("Newmark1935 result :%v\n", result)
	return result
}

func Inv(i float64) float64 {
	return 1 / i
}
