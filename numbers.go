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
	"strconv"
)

// IntegerLen returns the number of digits in the generic Integers given
func IntegerLen[V Integers](v V) (length int) {
	if v == 0 {
		length = 1
		return
	}
	for v != 0 {
		v /= 10
		length++
	}
	return
}

// DecimalLen returns the number of digits in the generic Decimal given
func DecimalLen[V Decimal](v V) (length int) {
	length = len(strconv.FormatFloat(float64(v), 'g', -1, 64))
	return
}
