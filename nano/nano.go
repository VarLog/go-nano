// Copyright 2016 Maxim Fedorenko. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nano

import (
	"fmt"
	"math"
	"math/rand"
)

const (
	Anisotropy        = 4000.
	Saturation        = 800.
	Damping           = (2. * Anisotropy) / Saturation
	Radius            = 20.e-7
	GyromagneticRatio = 1.76e+7
)

func Calculate(fieldStrengthStart, dt, epsillon float64) (result *Vector, iterCount int) {
	rand.Seed(42) // For debug

	volume := (4. / 3.) * math.Pi * math.Pow(Radius, 3.)
	t := 0.
	n := rand.Float64()

	fmt.Printf("Nano %v!\n", n)
	fmt.Printf("Volume %v!\n", volume)
	fmt.Printf("Time %v\n!", t)

	v := NewVectorRand()
	return v, 0
}

type Vector struct {
	x, y, z float64
}

func NewVectorRand() *Vector {
	v := &Vector{rand.Float64(), rand.Float64(), rand.Float64()}
	v.Normalize()
	return v
}

func (v *Vector) Mod() float64 {
	return math.Sqrt(math.Pow(v.x, 2.) + math.Pow(v.y, 2.) + math.Pow(v.z, 2.))
}

func (v *Vector) Normalize() {
	mod := v.Mod()
	v.x /= mod
	v.y /= mod
	v.z /= mod
}
