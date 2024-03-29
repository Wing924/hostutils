package hostutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPack(t *testing.T) {
	testPack(t, nil, nil)
	testPack(t, []string{}, []string{})
	testPack(t,
		[]string{},
		[]string{
			"#example101c.com",
			"#example103c.com",
			"#example104c.com",
		})
	testPack(t,
		[]string{"example[101-103]c.com"},
		[]string{"example101c.com   example102c.com\nexample103c.com"})
	testPack(t, []string{"www.example.com"}, []string{"www.example.com"})
	testPack(t,
		[]string{"example[101-105]c.com"},
		[]string{
			"example101c.com",
			"example102c.com",
			"example105c.com",
			"example104c.com",
			"example103c.com",
			"example103c.com",
		})
	testPack(t,
		[]string{"example-100-[101-105]c.com"},
		[]string{
			"example-100-101c.com",
			"example-100-102c.com",
			"example-100-105c.com",
			"example-100-104c.com",
			"example-100-103c.com",
			"example-100-103c.com",
		})
	testPack(t,
		[]string{"example[01-03]c.com"},
		[]string{
			"example01c.com",
			"example03c.com",
			"example02c.com",
		})
	testPack(t,
		[]string{"example[101-103,201]c.com"},
		[]string{
			"example101c.com",
			"example102c.com",
			"example103c.com",
			"example201c.com",
		})
	testPack(t,
		[]string{"example[101,103-105,201]c.com"},
		[]string{
			"example101c.com",
			"example103c.com",
			"example104c.com",
			"example105c.com",
			"example201c.com",
		})
	testPack(t,
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
	testPack(t,
		[]string{"example[101,103,105,201]c.com", "test[101-102]z.com"},
		[]string{
			"example101c.com",
			"example103c.com",
			"#example104c.com",
			"example105c.com",
			"example201c.com",
			"test101z.com",
			"test102z.com",
		})
	testPack(t,
		[]string{"example[01-02,102,0003]c.com"},
		[]string{
			"example01c.com",
			"example02c.com",
			"example102c.com",
			"example0003c.com",
		})
	testPack(t,
		[]string{"example[101-102]c.grp1.com"},
		[]string{
			"example101c.grp1.com",
			"example102c.grp1.com",
		})
	testPack(t,
		[]string{"example[101-102]"},
		[]string{
			"example101",
			"example102",
		})
}

func TestPackString(t *testing.T) {
	testPackString(t, []string{}, "")
	testPackString(t, []string{}, `#example101c.com
		#example102c.com
		#example103c.com`,
	)
	testPackString(t, []string{"example[101-103]c.com"}, "example101c.com   example102c.com\nexample103c.com")
	testPackString(t, []string{"www.example.com"}, "www.example.com")
	testPackString(t,
		[]string{"example[101-105]c.com"}, `
			example101c.com
			example102c.com
			example105c.com
			example104c.com
			example103c.com
			example103c.com`,
	)
	testPackString(t, []string{"example-100-[101-105]c.com"}, `
			example-100-101c.com
			example-100-102c.com
			example-100-105c.com
			example-100-104c.com
			example-100-103c.com
			example-100-103c.com`,
	)
}

func testPack(t *testing.T, expected []string, input []string) {
	t.Helper()
	assert.Equal(t, expected, Pack(input))
}

func testPackString(t *testing.T, expected []string, input string) {
	t.Helper()
	assert.Equal(t, expected, PackString(input))
}
