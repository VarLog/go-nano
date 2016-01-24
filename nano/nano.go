// Copyright 2016 Maxim Fedorenko. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package nano

import (
	"fmt"
	v "github.com/spate/vectormath"
	"math"
	"math/rand"
)

const (
	Anisotropy        = 4000.
	Saturation        = 800.
	Damping           = (20. * Anisotropy) / Saturation
	Radius            = 20.e-7
	GyromagneticRatio = 1.76e+7
)

func Calculate(field *v.Vector3, dt, epsillon float32) (magnetization *v.Vector3, iterCount int) {

	//volume := (4. / 3.) * math.Pi * math.Pow(Radius, 3.)

	//anisotropyAxis := &v.Vector3{}
	//v.V3MakeFromElems(anisotropyAxis, rand.Float32(), rand.Float32(), rand.Float32())

	magnetization = &v.Vector3{}
	v.V3MakeFromElems(magnetization, rand.Float32(), rand.Float32(), rand.Float32())

	iterCount = 0

	for {
		v.V3Normalize(magnetization, magnetization)

		//fmt.Printf("magnetization %v\n", magnetization)

		fieldEffectiveAnisotropy := &v.Vector3{}
		fieldEffectiveAnisotropy.X = -Damping * magnetization.X
		fieldEffectiveAnisotropy.Y = -Damping * magnetization.Y
		fieldEffectiveAnisotropy.Z = 0.

		fieldEffective := &v.Vector3{}
		v.V3Add(fieldEffective, fieldEffectiveAnisotropy, field)

		fieldR := &v.Vector3{}
		{
			vec := &v.Vector3{}
			v.V3Cross(vec, fieldEffective, magnetization)
			vec.X *= Damping
			vec.Y *= Damping
			vec.Z *= Damping
			v.V3Sub(fieldR, fieldEffective, vec)
		}

		fieldM := fieldR.Length()

		v.V3Normalize(fieldR, fieldR)

		dte := (GyromagneticRatio * fieldM * dt) / (1 + float32(math.Pow(Damping, 2.)))

		res := &v.Vector3{}

		{
			vec1 := &v.Vector3{}
			v.V3Copy(vec1, fieldR)

			dot := v.V3Dot(fieldR, magnetization)
			vec1.X *= dot
			vec1.Y *= dot
			vec1.Z *= dot

			dte2 := float32(math.Pow(float64(dte), 2.))

			vec1.X *= dte2
			vec1.Y *= dte2
			vec1.Z *= dte2

			vec2 := &v.Vector3{}
			v.V3Cross(vec2, fieldR, magnetization)

			vec2.X *= dte
			vec2.Y *= dte
			vec2.Z *= dte

			v.V3Add(res, magnetization, vec1)
			v.V3Add(res, res, vec2)

			c := 1. / (1. + (dte2 * fieldR.LengthSqr()))

			res.X *= c
			res.Y *= c
			res.Z *= c
		}

		{
			v.V3Normalize(res, res)

			vec := &v.Vector3{}
			v.V3Cross(vec, res, magnetization)
			dot := v.V3Dot(res, magnetization)

			fmt.Printf("res - magnetization %v\n", vec.Length())
			fmt.Printf("res - magnetization %v\n", dot)
		}
		magnetization = res
		iterCount++

		{
			v.V3Normalize(fieldEffective, fieldEffective)
			vec := &v.Vector3{}
			v.V3Cross(vec, magnetization, fieldEffective)

			fmt.Printf("Diff %v\n", vec.Length())
			if vec.Length() <= epsillon {
				break
			}
		}
	}

	return

}
