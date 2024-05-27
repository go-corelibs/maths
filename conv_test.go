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

func TestAs(t *testing.T) {


	Convey("Atoi", t, func() {
		So(Atoi("10"), ShouldEqual, 10)
		So(Atoi("nope"), ShouldEqual, math.MaxInt)
		So(Atoi("nope", 1010), ShouldEqual, 1010)
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
