package hostutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnpack(t *testing.T) {
	testUnpack(t, nil, nil)
	testUnpack(t, []string{}, []string{})
	testUnpack(t, []string{"aa[123*2]b"}, []string{"aa[123*2]b"})
	testUnpack(t, []string{""}, []string{})
	testUnpack(t,
		[]string{"example[999-1001]c.com"},
		[]string{
			"example0999c.com",
			"example1000c.com",
			"example1001c.com",
		})
	testUnpack(t,
		[]string{"example[1001-999]c.com"},
		[]string{
			"example0999c.com",
			"example1000c.com",
			"example1001c.com",
		})
	testUnpack(t,
		[]string{"example[101-103]c.com"},
		[]string{
			"example101c.com",
			"example102c.com",
			"example103c.com",
		})
	testUnpack(t,
		[]string{"example[1001,101-102]c.com"},
		[]string{
			"example1001c.com",
			"example101c.com",
			"example102c.com",
		})
	testUnpack(t,
		[]string{"example[103-101]c.com"},
		[]string{
			"example101c.com",
			"example102c.com",
			"example103c.com",
		})
	testUnpack(t,
		[]string{"example[1-2][01-02]c.com"},
		[]string{
			"example101c.com",
			"example102c.com",
			"example201c.com",
			"example202c.com",
		})
	testUnpack(t, []string{"www.example.com"}, []string{"www.example.com"})
	testUnpack(t,
		[]string{"example[101-105]c.com"},
		[]string{
			"example101c.com",
			"example102c.com",
			"example103c.com",
			"example104c.com",
			"example105c.com",
		})
	testUnpack(t,
		[]string{"example-100-[101-105]c.com"},
		[]string{
			"example-100-101c.com",
			"example-100-102c.com",
			"example-100-103c.com",
			"example-100-104c.com",
			"example-100-105c.com",
		})
	testUnpack(t,
		[]string{"example[01-03]c.com"},
		[]string{
			"example01c.com",
			"example02c.com",
			"example03c.com",
		})
	testUnpack(t,
		[]string{"example[101-103,201]c.com"},
		[]string{
			"example101c.com",
			"example102c.com",
			"example103c.com",
			"example201c.com",
		})
	testUnpack(t,
		[]string{"example[101,103-105,201]c.com"},
		[]string{
			"example101c.com",
			"example103c.com",
			"example104c.com",
			"example105c.com",
			"example201c.com",
		})
	testUnpack(t,
		[]string{"example[101,103-105,201]c.com", "test[101-102]z.com"},
		[]string{
			"example101c.com",
			"example103c.com",
			"example104c.com",
			"example105c.com",
			"example201c.com",
			"test101z.com",
			"test102z.com",
		})
	testUnpack(t,
		[]string{"example[101,103,105,201]c.com", "test[101-102]z.com"},
		[]string{
			"example101c.com",
			"example103c.com",
			"example105c.com",
			"example201c.com",
			"test101z.com",
			"test102z.com",
		})
	testUnpack(t,
		[]string{"example[01-02,102,0003]c.com"},
		[]string{
			"example0003c.com",
			"example01c.com",
			"example02c.com",
			"example102c.com",
		})
}

func TestUnpackString(t *testing.T) {
	testUnpackString(t, "", []string{})
	testUnpackString(t, "aa[123*2]b", []string{"aa[123*2]b"})
	testUnpackString(t, "example[101,103,105,201]c.com test[101-102]z.com",
		[]string{
			"example101c.com",
			"example103c.com",
			"example105c.com",
			"example201c.com",
			"test101z.com",
			"test102z.com",
		})
	testUnpackString(t, `
		example[101,103,105,201]c.com
		test[101-102]z.com`,
		[]string{
			"example101c.com",
			"example103c.com",
			"example105c.com",
			"example201c.com",
			"test101z.com",
			"test102z.com",
		})

}

func testUnpack(t *testing.T, input []string, expected []string) {
	assert.Equal(t, expected, Unpack(input))
}

func testUnpackString(t *testing.T, input string, expected []string) {
	assert.Equal(t, expected, UnpackString(input))
}
