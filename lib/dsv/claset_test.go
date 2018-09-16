// Copyright 2015-2018, Shulhan <ms@kilabit.info>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package dsv

import (
	"testing"

	"github.com/shuLhan/share/lib/tabula"
	"github.com/shuLhan/share/lib/test"
)

func TestReaderWithClaset(t *testing.T) {
	fcfg := "testdata/claset.dsv"

	claset := tabula.Claset{}

	_, e := NewReader(fcfg, &claset)
	if e != nil {
		t.Fatal(e)
	}

	test.Assert(t, "", 3, claset.GetClassIndex(), true)

	claset.SetMajorityClass("regular")
	claset.SetMinorityClass("vandalism")

	clone := claset.Clone().(tabula.ClasetInterface)

	test.Assert(t, "", 3, clone.GetClassIndex(), true)
	test.Assert(t, "", "regular", clone.MajorityClass(), true)
	test.Assert(t, "", "vandalism", clone.MinorityClass(), true)
}
