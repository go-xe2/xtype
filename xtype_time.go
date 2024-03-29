// Copyright 2017 gf Author(https://github.com/gogf/gf). All Rights Reserved.
//
// This Source Code Form is subject to the terms of the MIT License.
// If a copy of the MIT was not distributed with this file,
// You can obtain one at https://github.com/gogf/gf.

package xtype

import (
	"github.com/go-xe2/xtype/xtime"
	"time"
)

// Time converts <i> to time.Time.
func Time(i interface{}, format ...string) time.Time {
	if t := XTime(i, format...); t != nil {
		return t.Time
	}
	return time.Time{}
}

// Duration converts <i> to time.Duration.
// If <i> is string, then it uses time.ParseDuration to convert it.
// If <i> is numeric, then it converts <i> as nanoseconds.
func Duration(i interface{}) time.Duration {
	s := String(i)
	if !IsNumeric(s) {
		d, _ := time.ParseDuration(s)
		return d
	}
	return time.Duration(Int64(i))
}

// GTime converts <i> to *gtime.Time.
// The parameter <format> can be used to specify the format of <i>.
// If no <format> given, it converts <i> using gtime.NewFromTimeStamp if <i> is numeric,
// or using gtime.StrToTime if <i> is string.
func XTime(i interface{}, format ...string) *xtime.Time {
	s := String(i)
	if len(s) == 0 {
		return xtime.New()
	}
	// Priority conversion using given format.
	if len(format) > 0 {
		t, _ := xtime.StrToTimeFormat(s, format[0])
		return t
	}
	if IsNumeric(s) {
		return xtime.NewFromTimeStamp(Int64(s))
	} else {
		t, _ := xtime.StrToTime(s)
		return t
	}
}
