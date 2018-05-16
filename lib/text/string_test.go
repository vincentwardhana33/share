package text

import (
	"testing"

	"github.com/shuLhan/share/lib/test"
)

func TestStringJSONEscape(t *testing.T) {
	cases := []struct {
		in  string
		exp string
	}{{
		in:  "",
		exp: "",
	}, {
		in: `	this\ is
		//\"☺"`,
		exp: `\tthis\\ is\n\t\t\/\/\\\"☺\"`,
	}, {
		in: ` `,
		exp: `\u0002\b\f\u000E\u000F\u0010\u0014\u001E\u001F `,
	}}

	var got string

	for _, c := range cases {
		t.Log(c)

		got = StringJSONEscape(c.in)

		test.Assert(t, "", c.exp, got, true)
	}
}

func TestStringJSONUnescape(t *testing.T) {
	cases := []struct {
		in     string
		strict bool
		exp    string
		expErr string
	}{{
		in:  "",
		exp: "",
	}, {
		in: `\tthis\\ is\n\t\t\/\/\\\"☺\"`,
		exp: `	this\ is
		//\"☺"`,
	}, {
		in: `\u0002\b\f\u000E\u000F\u0010\u0014\u001E\u001F\u263A `,
		exp: `☺ `,
	}, {
		in:     `\uerror`,
		expErr: `strconv.ParseUint: parsing "erro": invalid syntax`,
	}, {
		in:  `\x`,
		exp: "x",
	}, {
		in:     `\x`,
		strict: true,
		expErr: `BytesJSONUnescape: invalid syntax at 1`,
	}}

	var (
		got string
		err error
	)

	for _, c := range cases {
		t.Log(c)

		got, err = StringJSONUnescape(c.in, c.strict)
		if err != nil {
			test.Assert(t, "err", c.expErr, err.Error(), true)
			continue
		}

		test.Assert(t, "value", c.exp, got, true)
	}

}
