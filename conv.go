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
	"reflect"
	"strconv"
)

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
	return defRv(math.MaxInt, def)
}

// ToNumber is a generic function for detecting the arbitrary value (v) type
// and converting the value (by recasting or by strconv parsing) to another
// Number type, using reflection to determine the correct means of conversion
//
// If the value given cannot be transformed into the requested Number type the
// "ok" return value will be false. In these cases, if value returned will be
// either the first def value given or zero
//
// Examples:
//
//	i, ok := ToNumber[int](struct{string}{"nope"}, 10)
//	// ok == false; i == int(10)
//
//	i, ok := ToNumber[int](struct{string}{"nope"})
//	// ok == false; i == int(0)
//
//	type Thing string
//	v := Thing("10")
//	i, ok := ToNumber[int](v)
//	// ok == true; i == int(10)
func ToNumber[V Number](v interface{}, def ...V) (value V, ok bool) {

	var sok bool
	var str string

	rv := reflect.ValueOf(v)
	if !rv.IsValid() {
		return 0, false
	}

	switch rv.Type().Kind() {

	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return V(rv.Int()), true

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return V(rv.Uint()), true

	case reflect.Float32, reflect.Float64:
		return V(rv.Float()), true

	case reflect.String:
		sok = true
		str = fmt.Sprintf("%v", rv.Interface())

	case reflect.Slice:

		switch rv.Type().Elem().Kind() {
		case reflect.Uint8:
			sok = true
			str = string(rv.Bytes())
		}

	}

	if sok {
		if i, err := strconv.Atoi(str); err == nil {
			return V(i), true
		} else if f, err := strconv.ParseFloat(str, 64); err == nil {
			return V(f), true
		}
	}

	if len(def) > 0 {
		return def[0], false
	}

	return
}

// ToInt is a convenience wrapper around ToNumber with the primary difference
// that the "counld not convert and no default" case returns math.MaxInt
// instead of zero
func ToInt(v interface{}, def ...int) int {
	if i, ok := ToNumber[int](v); ok {
		return i
	}
	return defRv(math.MaxInt, def)
}

// ToInt64 is a convenience wrapper around ToNumber with the primary difference
// that the "counld not convert and no default" case returns math.MaxInt64
// instead of zero
func ToInt64(v interface{}, def ...int64) int64 {
	if i, ok := ToNumber[int64](v); ok {
		return i
	}
	return defRv(math.MaxInt64, def)
}

// ToUint is a convenience wrapper around ToNumber with the primary difference
// that the "counld not convert and no default" case returns math.MaxUint
// instead of zero
func ToUint(v interface{}, def ...uint) uint {
	if i, ok := ToNumber[uint](v); ok {
		return i
	}
	return defRv(math.MaxUint, def)
}

// ToUint64 is a convenience wrapper around ToNumber with the primary
// difference that the "counld not convert and no default" case returns
// math.MaxUint64 instead of zero
func ToUint64(v interface{}, def ...uint64) uint64 {
	if i, ok := ToNumber[uint64](v); ok {
		return i
	}
	return defRv(math.MaxUint64, def)
}

// ToFloat64 is a convenience wrapper around ToNumber with the primary
// difference that the "counld not convert and no default" case returns
// math.MaxFloat64 instead of zero
func ToFloat64(v interface{}, def ...float64) float64 {
	if i, ok := ToNumber[float64](v); ok {
		return i
	}
	return defRv(math.MaxFloat64, def)
}

func defRv[V Number](m V, def []V) (value V) {
	if len(def) > 0 {
		return def[0]
	}
	return m
}
