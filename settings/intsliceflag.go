// Copyright 2018 Google Inc. All Rights Reserved.
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

// Package settings provides a central interface to webborer settings.
package settings

import (
	"fmt"
	"strconv"
	"strings"
)

// IntSliceFlag is a flag.Value that takes a comma-separated string and turns
// it into a slice of ints.
type IntSliceFlag []int

func (f *IntSliceFlag) String() string {
	if f == nil {
		return ""
	}
	tmpslice := []string{}
	for _, v := range *f {
		tmpslice = append(tmpslice, strconv.Itoa(v))
	}
	return strings.Join(tmpslice, ",")
}

func (f *IntSliceFlag) Set(value string) error {
	for _, v := range strings.Split(value, ",") {
		if i, err := strconv.Atoi(strings.TrimSpace(v)); err == nil {
			*f = append(*f, i)
		} else {
			return fmt.Errorf("Unable to parse %s as int.", v)
		}
	}
	return nil
}
