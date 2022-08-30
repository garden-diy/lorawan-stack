// Copyright © 2022 The Things Network Foundation, The Things Industries B.V.
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

// Package normalizedpayload implements functionality for parsing and validating normalized payload.
package normalizedpayload

import (
	"fmt"
	"time"

	pbtypes "github.com/gogo/protobuf/types"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"golang.org/x/exp/constraints"
)

// Air is an air measurement.
type Air struct {
	Temperature      *float64
	RelativeHumidity *float64
	Pressure         *float64
}

// Wind is a wind measurement.
type Wind struct {
	Speed     *float64
	Direction *float64
}

// Measurement is a measurement.
type Measurement struct {
	Time *time.Time
	Air  *Air
	Wind *Wind
}

var (
	errFieldType    = errors.DefineInvalidArgument("field_type", "invalid field type of `{path}`")
	errFieldMinimum = errors.DefineInvalidArgument(
		"field_minimum",
		"`{path}` should be equal or greater than `{minimum}`",
	)
	//nolint:unused
	errFieldExclusiveMinimum = errors.DefineInvalidArgument(
		"field_exclusive_minimum",
		"`{path}` should be greater than `{minimum}`",
	)
	errFieldMaximum = errors.DefineInvalidArgument(
		"field_maximum",
		"`{path}` should be equal or less than `{maximum}`",
	)
	errFieldExclusiveMaximum = errors.DefineInvalidArgument(
		"field_exclusive_maximum",
		"`{path}` should be less than `{maximum}`",
	)
	errUnknownField = errors.DefineInvalidArgument("unknown_field", "unknown field `{path}`")
)

type fieldParser func(dst *Measurement, src *pbtypes.Value, path string) error

// object validates that the path is a structure and sets the target to an empty value.
func object[T any](selector func(*Measurement) **T) fieldParser {
	return func(dst *Measurement, src *pbtypes.Value, path string) error {
		_, ok := src.Kind.(*pbtypes.Value_StructValue)
		if !ok {
			return errFieldType.WithAttributes("path", path)
		}
		*selector(dst) = new(T)
		return nil
	}
}

type fieldValidator[T any] func(v T, path string) error

func validate[T any](val T, validators []fieldValidator[T], path string) error {
	for _, v := range validators {
		if err := v(val, path); err != nil {
			return err
		}
	}
	return nil
}

// parseTime parses and validates the time. The input value must be RFC3339.
func parseTime(selector func(dst *Measurement) **time.Time, vals ...fieldValidator[time.Time]) fieldParser {
	return func(dst *Measurement, src *pbtypes.Value, path string) error {
		val, ok := src.Kind.(*pbtypes.Value_StringValue)
		if !ok {
			return errFieldType.WithAttributes("path", path)
		}
		t, err := time.Parse(time.RFC3339Nano, val.StringValue)
		if err != nil {
			return err
		}
		if err := validate(t, vals, path); err != nil {
			return err
		}
		*selector(dst) = &t
		return nil
	}
}

// parseNumber parses and validates a number.
func parseNumber(selector func(dst *Measurement) **float64, vals ...fieldValidator[float64]) fieldParser {
	return func(dst *Measurement, src *pbtypes.Value, path string) error {
		val, ok := src.Kind.(*pbtypes.Value_NumberValue)
		if !ok {
			return errFieldType.WithAttributes("path", path)
		}
		n := val.NumberValue
		if err := validate(n, vals, path); err != nil {
			return err
		}
		*selector(dst) = &n
		return nil
	}
}

// minimum returns a field validator that checks the inclusive minimum.
func minimum[T constraints.Ordered](min T) fieldValidator[T] {
	return func(v T, path string) error {
		if v < min {
			return errFieldMinimum.WithAttributes(
				"path", path,
				"minimum", min,
			)
		}
		return nil
	}
}

// exclusiveMinimum returns a field validator that checks the exclusive minimum.
//
//nolint:unused,deadcode
func exclusiveMinimum[T constraints.Ordered](min T) fieldValidator[T] {
	return func(v T, path string) error {
		if v <= min {
			return errFieldExclusiveMinimum.WithAttributes(
				"path", path,
				"minimum", min,
			)
		}
		return nil
	}
}

// maximum returns a field validator that checks the inclusive maximum.
func maximum[T constraints.Ordered](max T) fieldValidator[T] {
	return func(v T, path string) error {
		if v > max {
			return errFieldMaximum.WithAttributes(
				"path", path,
				"maximum", max,
			)
		}
		return nil
	}
}

// exclusiveMaximum returns a field validator that checks the exclusive maximum.
func exclusiveMaximum[T constraints.Ordered](max T) fieldValidator[T] {
	return func(v T, path string) error {
		if v >= max {
			return errFieldExclusiveMaximum.WithAttributes(
				"path", path,
				"maximum", max,
			)
		}
		return nil
	}
}

var fieldParsers = map[string]fieldParser{
	"time": parseTime(
		func(dst *Measurement) **time.Time {
			return &dst.Time
		},
	),
	"air": object(
		func(dst *Measurement) **Air {
			return &dst.Air
		},
	),
	"air.temperature": parseNumber(
		func(dst *Measurement) **float64 {
			return &dst.Air.Temperature
		},
		minimum(-273.15),
	),
	"air.relativeHumidity": parseNumber(
		func(dst *Measurement) **float64 {
			return &dst.Air.RelativeHumidity
		},
		minimum(0.0),
		maximum(100.0),
	),
	"air.pressure": parseNumber(
		func(dst *Measurement) **float64 {
			return &dst.Air.Pressure
		},
		minimum(900.0),
		maximum(1100.0),
	),
	"wind": object(
		func(dst *Measurement) **Wind {
			return &dst.Wind
		},
	),
	"wind.speed": parseNumber(
		func(dst *Measurement) **float64 {
			return &dst.Wind.Speed
		},
		minimum(0.0),
	),
	"wind.direction": parseNumber(
		func(dst *Measurement) **float64 {
			return &dst.Wind.Direction
		},
		minimum(0.0),
		exclusiveMaximum(360.0),
	),
}

// Parse parses and validates the measurements.
func Parse(measurements []*pbtypes.Struct) ([]Measurement, error) {
	res := make([]Measurement, len(measurements))
	for i, src := range measurements {
		err := parse(&res[i], src, "")
		if err != nil {
			return nil, err
		}
	}
	return res, nil
}

func parse(dst *Measurement, src *pbtypes.Struct, prefix string) error {
	for k, v := range src.GetFields() {
		path := fmt.Sprintf("%s%s", prefix, k)
		parser, ok := fieldParsers[path]
		if !ok {
			return errUnknownField.WithAttributes("path", path)
		}
		if err := parser(dst, v, path); err != nil {
			return err
		}
		if s, ok := v.Kind.(*pbtypes.Value_StructValue); ok {
			if err := parse(dst, s.StructValue, path+"."); err != nil {
				return err
			}
		}
	}
	return nil
}
