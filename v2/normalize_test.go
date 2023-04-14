package hostutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNormalize(t *testing.T) {
	testNormalize(t, nil, nil)
	testNormalize(t, []string{}, nil)
	testNormalize(t, []string{"ww[101]"}, []string{"ww101"})
	testNormalize(t, []string{"ww[101-103]", "ww[104]"}, []string{"ww[101-104]"})
	testNormalize(t, []string{"a[101,103,105]", "a[102,104,107]"}, []string{"a[101-105,107]"})
	testNormalize(t, []string{"[101-103]", "104"}, []string{"[101-104]"})
}

func TestNormalizeString(t *testing.T) {
	testNormalizeString(t, "", nil)
	testNormalizeString(t, "ww[101]", []string{"ww101"})
	testNormalizeString(t, "ww[101-103] ww[104]", []string{"ww[101-104]"})
	testNormalizeString(t, "a[101,103,105] a[102,104,107]", []string{"a[101-105,107]"})
	testNormalizeString(t, "[101-103] 104", []string{"[101-104]"})
}

func testNormalize(t *testing.T, input []string, expected []string) {
	assert.Equal(t, expected, Normalize(input))
}

func testNormalizeString(t *testing.T, input string, expected []string) {
	assert.Equal(t, expected, NormalizeString(input))
}
