package hostutils

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestPack(t *testing.T) {
	cases := []struct {
		input, output []string
	}{
		{nil, nil},
		{[]string{""}, nil},
		{[]string{"abc101z"}, []string{"abc101z"}},
		{[]string{"abc101z", "abc102z"}, []string{"abc[101-102]z"}},
		{[]string{"abc103z", "abc102z", "abc101z"}, []string{"abc[101-103]z"}},
		{[]string{"abc101z", "abc102z", "abc104z"}, []string{"abc[101-102,104]z"}},
		{[]string{"abc101z", "abc102z", "abc104"}, []string{"abc104", "abc[101-102]z"}},
		{[]string{"abc101z-02", "abc102z-02", "abc102z-03"}, []string{"abc[101-102]z-02", "abc102z-03"}},
		{[]string{"abc101z-01", "abc102z-02", "abc102z-03"}, []string{"abc101z-01", "abc102z-[02-03]"}},
		{[]string{"abc101z-01p", "abc101z-01q", "abc102z-01p", "abc102z-01q"}, []string{"abc[101-102]z-01p", "abc[101-102]z-01q"}},
	}
	for _, test := range cases {
		test := test
		t.Run(strings.Join(test.output, ";"), func(t *testing.T) {
			actual := Pack(test.input)
			assert.EqualValues(t, test.output, actual)
		})
	}
}

func TestParseHost(t *testing.T) {
	cases := []struct {
		input    string
		expected host
	}{
		{"", host{
			NonDigits: []string{""},
			Digits:    nil,
		}},
		{"a", host{
			NonDigits: []string{"a"},
			Digits:    nil,
		}},
		{"abc", host{
			NonDigits: []string{"abc"},
			Digits:    nil,
		}},
		{"1", host{
			NonDigits: []string{""},
			Digits:    []digit{{1, 1}},
		}},
		{"01", host{
			NonDigits: []string{""},
			Digits:    []digit{{1, 2}},
		}},
		{"a0", host{
			NonDigits: []string{"a"},
			Digits:    []digit{{0, 1}},
		}},
		{"abc101", host{
			NonDigits: []string{"abc"},
			Digits:    []digit{{101, 3}},
		}},
		{"abc011", host{
			NonDigits: []string{"abc"},
			Digits:    []digit{{11, 3}},
		}},
		{"abc001", host{
			NonDigits: []string{"abc"},
			Digits:    []digit{{1, 3}},
		}},
		{"abc001def", host{
			NonDigits: []string{"abc", "def"},
			Digits:    []digit{{1, 3}},
		}},
		{"abc001def2", host{
			NonDigits: []string{"abc", "def"},
			Digits:    []digit{{1, 3}, {2, 1}},
		}},
	}
	for _, test := range cases {
		test := test
		t.Run(test.input, func(t *testing.T) {
			actual := parseHost(test.input)
			assert.EqualValues(t, test.expected, *actual)
			assert.Equal(t, test.input, actual.String())
		})
	}
}

func TestMergeHost(t *testing.T) {
	cases := []struct {
		a, b, result string
	}{
		{"a1", "a2", "a[1-2]"},
		{"b1", "a2", ""},
		{"a1a", "a2", ""},
		{"a01", "a02", "a[01-02]"},
		{"a01", "a03", "a[01,03]"},
		{"a01x", "a03x", "a[01,03]x"},
		{"a01-02", "a02-02", "a[01-02]-02"},
		{"a01-02", "a01-03", "a01-[02-03]"},
		{"a01-02", "a03-04", ""},
		{"a01", "a001", "a[01,001]"},
	}

	for _, test := range cases {
		test := test
		t.Run(test.result, func(t *testing.T) {
			h1 := parseHost(test.a)
			h2 := parseHost(test.b)
			g := mergeHost(h1, h2)
			if test.result != "" {
				require.NotNil(t, g)
				assert.Equal(t, test.result, g.String())
			} else {
				require.Nil(t, g)
			}
		})
	}
}

func TestHostGroup_AppendHost(t *testing.T) {
	{
		g := &hostGroup{
			host:       *parseHost("abc101z"),
			RangeIndex: 0,
			Conds: []cond{
				{101, 102, 3},
			},
		}
		ok := g.AppendHost(parseHost("ab103z"))
		assert.False(t, ok)
	}
	{
		g := &hostGroup{
			host:       *parseHost("abc101z"),
			RangeIndex: 0,
			Conds: []cond{
				{101, 102, 3},
			},
		}
		ok := g.AppendHost(parseHost("abc103z1"))
		assert.False(t, ok)
	}
	{
		g := &hostGroup{
			host:       *parseHost("abc101z1"),
			RangeIndex: 0,
			Conds: []cond{
				{101, 102, 3},
			},
		}
		ok := g.AppendHost(parseHost("abc103z5"))
		assert.False(t, ok)
	}
	{
		g := &hostGroup{
			host:       *parseHost("abc101z"),
			RangeIndex: 0,
			Conds: []cond{
				{101, 102, 3},
			},
		}
		ok := g.AppendHost(parseHost("abc103z"))
		assert.True(t, ok)
		assert.Len(t, g.Conds, 1)
		assert.Equal(t, cond{101, 103, 3}, g.Conds[0])
	}
	{
		g := &hostGroup{
			host:       *parseHost("abc101z"),
			RangeIndex: 0,
			Conds: []cond{
				{101, 102, 3},
			},
		}
		ok := g.AppendHost(parseHost("abc104z"))
		assert.True(t, ok)
		assert.Len(t, g.Conds, 2)
		assert.Equal(t, cond{101, 102, 3}, g.Conds[0])
		assert.Equal(t, cond{104, 104, 3}, g.Conds[1])
	}
}
