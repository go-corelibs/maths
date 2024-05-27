// Copyright (c) 2023  The Go-Curses Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package maths

import (
	"fmt"
	"math"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type tString string

func TestAs(t *testing.T) {

	Convey("Atoi", t, func() {
		So(Atoi("10"), ShouldEqual, 10)
		So(Atoi("nope"), ShouldEqual, math.MaxInt)
		So(Atoi("nope", 1010), ShouldEqual, 1010)
	})

	Convey("ToNumber", t, func() {

		// inputs: int, uint, float, string, []byte
		invalid, ok := ToNumber[int](nil)
		So(ok, ShouldBeFalse)
		So(invalid, ShouldEqual, 0)

		i2f64, ok := ToNumber[float64](1010)
		So(ok, ShouldBeTrue)
		So(i2f64, ShouldEqual, 1010.0)

		u2f64, ok := ToNumber[float64](uint64(1010))
		So(ok, ShouldBeTrue)
		So(u2f64, ShouldEqual, 1010.0)

		f2i64, ok := ToNumber[int64](10.01)
		So(ok, ShouldBeTrue)
		So(f2i64, ShouldEqual, int64(10))

		s2u64, ok := ToNumber[uint64]("12345")
		So(ok, ShouldBeTrue)
		So(s2u64, ShouldEqual, uint64(12345))

		fs2u64, ok := ToNumber[uint64]("12.345")
		So(ok, ShouldBeTrue)
		So(fs2u64, ShouldEqual, uint64(12))

		b2f64, ok := ToNumber[float64]([]byte("10"))
		So(ok, ShouldBeTrue)
		So(b2f64, ShouldEqual, 10.0)

		ts2f64, ok := ToNumber[float64](tString("10"))
		So(ok, ShouldBeTrue)
		So(ts2f64, ShouldEqual, 10.0)

		nope, ok := ToNumber[int](struct{ string }{"nope"})
		So(ok, ShouldBeFalse)
		So(nope, ShouldEqual, 0)

		nope, ok = ToNumber[int](struct{ string }{"nope"}, 10)
		So(ok, ShouldBeFalse)
		So(nope, ShouldEqual, 10)

	})

	Convey("ToInt", t, func() {

		for idx, test := range []struct {
			input  interface{}
			def    []int
			output int
		}{
			{int(1), []int{0}, 1},
			{int8(1), []int{0}, 1},
			{int16(1), []int{0}, 1},
			{int32(1), []int{0}, 1},
			{int64(1), []int{0}, 1},
			{uint(1), []int{0}, 1},
			{uint8(1), []int{0}, 1},
			{uint16(1), []int{0}, 1},
			{uint32(1), []int{0}, 1},
			{uint64(1), []int{0}, 1},
			{float32(1), []int{0}, 1},
			{float64(1), []int{0}, 1},
			{"10", []int{0}, 10},
			{[]byte("10"), []int{0}, 10},
			{"nope", []int{1010}, 1010},
			{"nope", nil, math.MaxInt},
			{struct{ string }{"nope"}, nil, math.MaxInt},
		} {
			SoMsg(fmt.Sprintf("test #%d", idx), ToInt(test.input, test.def...), ShouldEqual, test.output)
		}

	})

	Convey("ToInt64", t, func() {

		for idx, test := range []struct {
			input  interface{}
			def    []int64
			output int64
		}{
			{int(1), []int64{0}, int64(1)},
			{int8(1), []int64{0}, int64(1)},
			{int16(1), []int64{0}, int64(1)},
			{int32(1), []int64{0}, int64(1)},
			{int64(1), []int64{0}, int64(1)},
			{uint(1), []int64{0}, int64(1)},
			{uint8(1), []int64{0}, int64(1)},
			{uint16(1), []int64{0}, int64(1)},
			{uint32(1), []int64{0}, int64(1)},
			{uint64(1), []int64{0}, int64(1)},
			{float32(1), []int64{0}, int64(1)},
			{float64(1), []int64{0}, int64(1)},
			{"10", []int64{0}, int64(10)},
			{[]byte("10"), []int64{0}, int64(10)},
			{"nope", []int64{1010}, 1010},
			{"nope", nil, math.MaxInt64},
			{struct{ string }{"nope"}, nil, math.MaxInt64},
		} {
			SoMsg(fmt.Sprintf("test #%d", idx), ToInt64(test.input, test.def...), ShouldEqual, test.output)
		}

	})

	Convey("ToUint", t, func() {

		for idx, test := range []struct {
			input  interface{}
			def    []uint
			output uint
		}{
			{int(1), []uint{0}, uint(1)},
			{int8(1), []uint{0}, uint(1)},
			{int16(1), []uint{0}, uint(1)},
			{int32(1), []uint{0}, uint(1)},
			{int64(1), []uint{0}, uint(1)},
			{uint(1), []uint{0}, uint(1)},
			{uint8(1), []uint{0}, uint(1)},
			{uint16(1), []uint{0}, uint(1)},
			{uint32(1), []uint{0}, uint(1)},
			{uint64(1), []uint{0}, uint(1)},
			{float32(1), []uint{0}, uint(1)},
			{float64(1), []uint{0}, uint(1)},
			{"10", []uint{0}, uint(10)},
			{[]byte("10"), []uint{0}, uint(10)},
			{"nope", []uint{1010}, 1010},
			{"nope", nil, math.MaxUint},
			{struct{ string }{"nope"}, nil, math.MaxUint},
		} {
			SoMsg(fmt.Sprintf("test #%d", idx), ToUint(test.input, test.def...), ShouldEqual, test.output)
		}

	})

	Convey("ToUint64", t, func() {

		for idx, test := range []struct {
			input  interface{}
			def    []uint64
			output uint64
		}{
			{int(1), []uint64{0}, uint64(1)},
			{int8(1), []uint64{0}, uint64(1)},
			{int16(1), []uint64{0}, uint64(1)},
			{int32(1), []uint64{0}, uint64(1)},
			{int64(1), []uint64{0}, uint64(1)},
			{uint(1), []uint64{0}, uint64(1)},
			{uint8(1), []uint64{0}, uint64(1)},
			{uint16(1), []uint64{0}, uint64(1)},
			{uint32(1), []uint64{0}, uint64(1)},
			{uint64(1), []uint64{0}, uint64(1)},
			{float32(1), []uint64{0}, uint64(1)},
			{float64(1), []uint64{0}, uint64(1)},
			{"10", []uint64{0}, uint64(10)},
			{[]byte("10"), []uint64{0}, uint64(10)},
			{"nope", []uint64{1010}, 1010},
			{"nope", nil, math.MaxUint64},
			{struct{ string }{"nope"}, nil, math.MaxUint64},
		} {
			SoMsg(fmt.Sprintf("test #%d", idx), ToUint64(test.input, test.def...), ShouldEqual, test.output)
		}
	})

	Convey("ToFloat64", t, func() {

		for idx, test := range []struct {
			input  interface{}
			def    []float64
			output float64
		}{
			{int(1), []float64{0}, float64(1)},
			{int8(1), []float64{0}, float64(1)},
			{int16(1), []float64{0}, float64(1)},
			{int32(1), []float64{0}, float64(1)},
			{int64(1), []float64{0}, float64(1)},
			{uint(1), []float64{0}, float64(1)},
			{uint8(1), []float64{0}, float64(1)},
			{uint16(1), []float64{0}, float64(1)},
			{uint32(1), []float64{0}, float64(1)},
			{uint64(1), []float64{0}, float64(1)},
			{float32(1), []float64{0}, float64(1)},
			{float64(1), []float64{0}, float64(1)},
			{"10", []float64{0}, float64(10)},
			{[]byte("10"), []float64{0}, float64(10)},
			{"nope", []float64{1010}, 1010},
			{"nope", nil, math.MaxFloat64},
			{struct{ string }{"nope"}, nil, math.MaxFloat64},
		} {
			SoMsg(fmt.Sprintf("test #%d", idx), ToFloat64(test.input, test.def...), ShouldEqual, test.output)
		}

	})

}
