// Copyright (c) 2024  The Go-CoreLibs Authors
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
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestNumbers(t *testing.T) {

	Convey("IntegerLen", t, func() {

		for idx, test := range []struct {
			input  int
			length int
		}{
			{0, 1},
			{1, 1},
			{10, 2},
			{34, 2},
			{1977, 4},
		} {
			SoMsg(fmt.Sprintf("%d: %d == %d", idx, test.input, test.length), IntegerLen(test.input), ShouldEqual, test.length)
		}

	})

	Convey("DecimalLen", t, func() {

		for idx, test := range []struct {
			input  float64
			length int
		}{
			{0, 1},
			{1.0, 1},
			{10.01, 5},
			{34.02, 5},
			{1977.101, 8},
		} {
			SoMsg(fmt.Sprintf("%d: %f == %d", idx, test.input, test.length), DecimalLen(test.input), ShouldEqual, test.length)
		}

	})

}
