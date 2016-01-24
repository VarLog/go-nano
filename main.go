// Copyright 2016 Maxim Fedorenko. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"fmt"
	"github.com/varlog/go-nano/nano"
)

const (
	epsillon = 1.e-8
	dt       = 1.e-3
	H        = 1.
)

func main() {
	fieldStrengthRange := [2]float64{1000, 1500}
	fmt.Println("Hello!")

	res, iterCount := nano.Calculate(fieldStrengthRange[0], dt, epsillon)
	mod := res.Mod()
	fmt.Printf("%v, %v, %v\n", res, mod, iterCount)
}
