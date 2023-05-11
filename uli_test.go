package uli

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetHex(t *testing.T) {

	a := NewUli("51bf608414ad5726a3c1bec098f77b1b54ffb2787f8d528a74c1d7fde6470ea4")
	b := "51bf608414ad5726a3c1bec098f77b1b54ffb2787f8d528a74c1d7fde6470ea4"

	assert.Equal(t, a.GetHex(), b, "The two numbers should be the same.")
}

func TestSetHex(t *testing.T) {

	a := NewUli("")
	a.SetHex("51bf608")
	b := []uint64{85718536}

	assert.Equal(t, a.bitsArr, b, "The two arrays should be the same.")
}

func TestOr(t *testing.T) {

	a := NewUli("403db")
	b := NewUli("51bf6")
	c := a.Or(b)
	d := NewUli("51bff")

	assert.Equal(t, d.GetHex(), c.GetHex(), "The two numbers should be the same.")
}

func TestAnd(t *testing.T) {

	a := NewUli("403db")
	b := NewUli("")
	c := a.And(b)
	d := NewUli("0")

	assert.Equal(t, d.GetHex(), c.GetHex(), "The two numbers should be the same.")
}

func TestXor(t *testing.T) {

	a := NewUli("51bf6")
	b := NewUli("403db")
	c := a.Xor(b)
	d := NewUli("1182D")

	assert.Equal(t, d.GetHex(), c.GetHex(), "The two numbers should be the same.")
}

func TestInv(t *testing.T) {

	a := NewUli("51bf6")
	c := a.Inv()
	b := NewUli("FFFFFFFFFFFAE409")

	assert.Equal(t, b.GetHex(), c.GetHex(), "The two numbers should be the same.")
}

func TestRightShift(t *testing.T) {

	a := NewUli("51bf6")
	c := a.RightShift(2)
	b := NewUli("146fd")

	assert.Equal(t, b.GetHex(), c.GetHex(), "The two numbers should be the same.")
}

func TestLeftShift(t *testing.T) {

	a := NewUli("aaaaaaaaaaaaaaaa")
	c := a.LeftShift(64)
	b := NewUli("aaaaaaaaaaaaaaaa0000000000000000")

	assert.Equal(t, b.GetHex(), c.GetHex(), "The two numbers should be the same.")
}

func TestAdd(t *testing.T) {
	a := NewUli("2")
	b := NewUli("ffffffffffffffff")
	b1 := NewUli("ffffffffffffffffffffffffffffffff")
	c := a.Add(b)
	c1 := b1.Add(a)
	d := NewUli("10000000000000001")
	d1 := NewUli("100000000000000000000000000000001")

	a1 := NewUli("0")
	b2 := NewUli("ffffffffffffffff")
	c2 := a1.Add(b2)
	d2 := NewUli("ffffffffffffffff")

	a2 := NewUli("ffffffffffffffff")
	b3 := NewUli("ffffffffffffffffffffffffffffffff")
	c3 := a2.Add(b3)
	d3 := NewUli("10000000000000000fffffffffffffffe")

	assert.Equal(t, d1.GetHex(), c1.GetHex(), "The two numbers should be the same.")
	assert.Equal(t, d.GetHex(), c.GetHex(), "The two numbers should be the same.")
	assert.Equal(t, d2.GetHex(), c2.GetHex(), "The two numbers should be the same.")
	assert.Equal(t, d3.GetHex(), c3.GetHex(), "The two numbers should be the same.")
}

func TestSub(t *testing.T) {
	a := NewUli("ffffffffffffffffffffffffffffffff0000000000000000")
	b := NewUli("ffffffffffffffff")
	c := a.Sub(b)
	d := NewUli("fffffffffffffffffffffffffffffffe0000000000000001")

	a1 := NewUli("ffffffffffffffffffffffffffffffff")
	b1 := NewUli("1")
	c1 := a1.Sub(b1)
	d1 := NewUli("fffffffffffffffffffffffffffffffe")

	a2 := NewUli("33ced2c76b26cae94e162c4c0d2c0ff7c13094b0185a3c122e732d5ba77efebc")
	b2 := NewUli("22e962951cb6cd2ce279ab0e2095825c141d48ef3ca9dabf253e38760b57fe03")
	c2 := a2.Sub(b2)
	d2 := NewUli("10e570324e6ffdbc6b9c813dec968d9bad134bc0dbb061530934f4e59c2700b9")

	assert.Equal(t, d.GetHex(), c.GetHex(), "The two numbers should be the same.")
	assert.Equal(t, d1.GetHex(), c1.GetHex(), "The two numbers should be the same.")
	assert.Equal(t, d2.GetHex(), c2.GetHex(), "The two numbers should be the same.")
}
