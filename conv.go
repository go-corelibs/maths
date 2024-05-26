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
	"strconv"
)

// The As* generic functions all recast their given Number to the standard go
// types as indicated in the names of the generic functions, ie: AsInt recasts
// to `int` and AsFloat64 recasts to `float64`

func AsInt[V Number](v V) int {
	return int(v)
}

func AsUint[V Number](v V) uint {
	return uint(v)
}

func AsInt64[V Number](v V) int64 {
	return int64(v)
}

func AsUint64[V Number](v V) uint64 {
	return uint64(v)
}

func AsFloat32[V Number](v V) float32 {
	return float32(v)
}

func AsFloat64[V Number](v V) float64 {
	return float64(v)
}

// Atoi is a wrapper around strconv.Atoi with the given value converted to a
// string first using fmt.Sprintf with a "%v" replacement
//
// def is an optional default value if the strconv.Atoi call returns an error,
// only the first def value is ever used and if there are no def values
// provided, math.MaxInt is returned
func Atoi(v interface{}, def ...int) (number int) {
	s := fmt.Sprintf("%v", v)
	if value, err := strconv.Atoi(s); err == nil {
		return value
	}
	if len(def) > 0 {
		return def[0]
	}
	return math.MaxInt
}

// The To* functions all accept an interface{} value and based on a standard
// go type switch, recasts int*, uint* and float* values to the type indicated
// in the function's name. For string and []byte values, strconv functions are
// used to convert the values into the specific type for the To* function.
// ie: ToInt converts to `int` and ToFloat64 converts to `float64`

func ToInt(v interface{}, def ...int) int {
	var i int
	var err error
	switch t := v.(type) {
	case int:
		return t
	case int8:
		return int(t)
	case int16:
		return int(t)
	case int32:
		return int(t)
	case int64:
		return int(t)
	case uint:
		return int(t)
	case uint8:
		return int(t)
	case uint16:
		return int(t)
	case uint32:
		return int(t)
	case uint64:
		return int(t)
	case float32:
		return int(t)
	case float64:
		return int(t)
	case string:
		i, err = strconv.Atoi(t)
	case []byte:
		i, err = strconv.Atoi(string(t))
	default:
	}
	return toRV(int(i), math.MaxInt, def, err)
}

func ToInt64(v interface{}, def ...int64) int64 {
	var i int64
	var err error
	switch t := v.(type) {
	case int:
		return int64(t)
	case int8:
		return int64(t)
	case int16:
		return int64(t)
	case int32:
		return int64(t)
	case int64:
		return t
	case uint:
		return int64(t)
	case uint8:
		return int64(t)
	case uint16:
		return int64(t)
	case uint32:
		return int64(t)
	case uint64:
		return int64(t)
	case float32:
		return int64(t)
	case float64:
		return int64(t)
	case string:
		i, err = strconv.ParseInt(t, 10, 64)
	case []byte:
		i, err = strconv.ParseInt(string(t), 10, 64)
	default:
	}
	return toRV(int64(i), math.MaxInt64, def, err)
}

func ToUint(v interface{}, def ...uint) uint {
	var i uint64
	var err error
	switch t := v.(type) {
	case int:
		return uint(t)
	case int8:
		return uint(t)
	case int16:
		return uint(t)
	case int32:
		return uint(t)
	case int64:
		return uint(t)
	case uint:
		return t
	case uint8:
		return uint(t)
	case uint16:
		return uint(t)
	case uint32:
		return uint(t)
	case uint64:
		return uint(t)
	case float32:
		return uint(t)
	case float64:
		return uint(t)
	case string:
		i, err = strconv.ParseUint(t, 10, 64)
	case []byte:
		i, err = strconv.ParseUint(string(t), 10, 64)
	default:
	}
	return toRV(uint(i), math.MaxUint, def, err)
}

func ToUint64(v interface{}, def ...uint64) uint64 {
	var i uint64
	var err error
	switch t := v.(type) {
	case int:
		return uint64(t)
	case int8:
		return uint64(t)
	case int16:
		return uint64(t)
	case int32:
		return uint64(t)
	case int64:
		return uint64(t)
	case uint:
		return uint64(t)
	case uint8:
		return uint64(t)
	case uint16:
		return uint64(t)
	case uint32:
		return uint64(t)
	case uint64:
		return t
	case float32:
		return uint64(t)
	case float64:
		return uint64(t)
	case string:
		i, err = strconv.ParseUint(t, 10, 64)
	case []byte:
		i, err = strconv.ParseUint(string(t), 10, 64)
	default:
	}
	return toRV(i, math.MaxUint64, def, err)
}

func ToFloat64(v interface{}, def ...float64) float64 {
	var f float64
	var err error
	switch t := v.(type) {
	case int:
		return float64(t)
	case int8:
		return float64(t)
	case int16:
		return float64(t)
	case int32:
		return float64(t)
	case int64:
		return float64(t)
	case uint:
		return float64(t)
	case uint8:
		return float64(t)
	case uint16:
		return float64(t)
	case uint32:
		return float64(t)
	case uint64:
		return float64(t)
	case float32:
		return float64(t)
	case float64:
		return t
	case string:
		f, err = strconv.ParseFloat(t, 64)
	case []byte:
		f, err = strconv.ParseFloat(string(t), 64)
	default:
	}
	return toRV(f, math.MaxFloat64, def, err)
}

func toRV[V Number](v, m V, def []V, err error) (value V) {
	if err == nil {
		return v
	} else if len(def) > 0 {
		return def[0]
	}
	return m
}
