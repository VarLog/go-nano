// Copyright 2016 Maxim Fedorenko. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	v "github.com/spate/vectormath"
	"github.com/varlog/go-nano/nano"
	"math/rand"
)

const (
	epsillon = 1.e-4
	dt       = 1.e-8
	dH       = 1.
)

func main() {
	fieldStrengthRange := [2]float32{1000, 1500}

	rand.Seed(420) // For debug

	field := &v.Vector3{}
	v.V3MakeFromElems(field, rand.Float32(), rand.Float32(), rand.Float32())
	v.V3ScalarMul(field, field, fieldStrengthRange[0])

	fmt.Printf("Field Strength %v\n", field)
	res, iterCount := nano.Calculate(field, dt, epsillon)

	fmt.Printf("Magnetization %v, Count of the iterations: %v\n", res, iterCount)

}
