package uli

import (
	"fmt"
	"strconv"
)

type Uli struct {
	bitsArr []uint64
}

func NewUli(inputNum string) *Uli {
	var len int = len(inputNum)
	bitsArr := make([]uint64, len)

	bitsArr = Div2Blocks(inputNum)
	num := Uli{bitsArr: bitsArr}

	return &num
}

func (num Uli) GetHex() string {
	len := len(num.bitsArr)
	res := ""

	for i := 0; i < len; i++ {
		res += fmt.Sprintf("%016x", num.bitsArr[i])
	}

	return res
}

func (num *Uli) SetHex(inputNum string) {
	blocksArr := Div2Blocks(inputNum)
	num.bitsArr = blocksArr
}

func (num *Uli) Or(secondNum *Uli) *Uli {
	var longerLen int
	var result = NewUli("")

	if len(num.bitsArr) > len(secondNum.bitsArr) {
		longerLen = len(num.bitsArr)
		secondNum.AddZeros(longerLen - len(secondNum.bitsArr))
	} else {
		longerLen = len(secondNum.bitsArr)
		num.AddZeros(longerLen - len(num.bitsArr))
	}

	result.bitsArr = make([]uint64, longerLen)

	for i := 0; i < longerLen; i++ {
		result.bitsArr[i] = num.bitsArr[i] | secondNum.bitsArr[i]
	}

	return result
}

func (num *Uli) And(secondNum *Uli) *Uli {
	var longerLen int
	var result = NewUli("")

	if len(num.bitsArr) > len(secondNum.bitsArr) {
		longerLen = len(num.bitsArr)
		secondNum.AddZeros(longerLen - len(secondNum.bitsArr))
	} else {
		longerLen = len(secondNum.bitsArr)
		num.AddZeros(longerLen - len(num.bitsArr))
	}

	result.bitsArr = make([]uint64, longerLen)

	for i := 0; i < longerLen; i++ {
		result.bitsArr[i] = num.bitsArr[i] & secondNum.bitsArr[i]
	}

	return result
}

func (num *Uli) Xor(secondNum *Uli) *Uli {
	var longerLen int
	result := NewUli("")

	if len(num.bitsArr) > len(secondNum.bitsArr) {
		longerLen = len(num.bitsArr)
		secondNum.AddZeros(longerLen - len(secondNum.bitsArr))
	} else {
		longerLen = len(secondNum.bitsArr)
		num.AddZeros(longerLen - len(num.bitsArr))
	}

	result.bitsArr = make([]uint64, longerLen)

	for i := 0; i < longerLen; i++ {
		result.bitsArr[i] = num.bitsArr[i] ^ secondNum.bitsArr[i]
	}

	return result
}

func (num *Uli) Inv() *Uli {
	result := NewUli("")
	len := len(num.bitsArr)
	result.bitsArr = make([]uint64, len)

	for i := 0; i < len; i++ {
		result.bitsArr[i] = ^num.bitsArr[i]
	}

	return result
}

func (num *Uli) RightShift(shift int) *Uli {
	var lostBits uint64
	var younglostBits uint64
	result := NewUli("")
	intShift := shift / 64
	modeShift := shift % 64

	l := len(num.bitsArr)
	copybitsArr := num.bitsArr

	copybitsArr = copybitsArr[:l-intShift]

	var power uint64 = uint64(1 << modeShift)
	var lostbitsShift uint64 = uint64(64 - modeShift)
	result.bitsArr = make([]uint64, len(copybitsArr))

	for i := 0; i < len(copybitsArr); i++ {
		younglostBits = copybitsArr[i] % power
		result.bitsArr[i] = copybitsArr[i]>>uint64(modeShift) + lostBits<<lostbitsShift
		lostBits = younglostBits
	}

	return result
}

func (num *Uli) LeftShift(shift int) *Uli {
	var lostBits uint64
	var oldlostBits uint64
	result := NewUli("")
	intShift := shift / 64
	modeShift := shift % 64

	copybitsArr := num.bitsArr

	for j := 0; j < intShift; j++ {
		copybitsArr = append(copybitsArr, uint64(0))
		fmt.Println(copybitsArr)
	}

	if modeShift == 0 {
		result.bitsArr = copybitsArr
		return result
	}

	var lostbitsShift uint64 = uint64(64 - modeShift)
	result.bitsArr = make([]uint64, intShift+len(copybitsArr))

	for i := len(copybitsArr) - 1; i >= 0; i-- {
		oldlostBits = copybitsArr[i] >> lostbitsShift
		result.bitsArr[i] = copybitsArr[i]<<uint64(modeShift) + lostBits
		lostBits = oldlostBits
	}

	if oldlostBits != 0 {
		var biggerResult = make([]uint64, 1+len(result.bitsArr))
		biggerResult[0] = lostBits
		for t := 0; t < len(result.bitsArr); t++ {
			biggerResult[t+1] = result.bitsArr[t]
		}
		result.bitsArr = biggerResult
	}

	return result
}

func (num *Uli) Add(secondNum *Uli) *Uli {

	if AllZeros(num.bitsArr) {
		return secondNum
	} else if AllZeros(secondNum.bitsArr) {
		return num
	}

	cf := uint64(0)
	l := 0

	const maxValue uint64 = 18446744073709551615

	if len(num.bitsArr) < len(secondNum.bitsArr) {
		l = len(secondNum.bitsArr)
	} else {
		l = len(num.bitsArr)
	}

	result := NewUli("")
	result.bitsArr = make([]uint64, l+1)

	for i := 0; i < l || cf != 0; i++ {
		var value1 uint64
		var value2 uint64

		if i < len(secondNum.bitsArr) {
			value1 = secondNum.bitsArr[len(secondNum.bitsArr)-i-1]
		} else {
			value1 = 0
		}

		if i < len(num.bitsArr) {
			value2 = num.bitsArr[len(num.bitsArr)-i-1]
		} else {
			value2 = 0
		}

		if cf == 1 {
			if value1 >= maxValue-value2 || value2 >= maxValue-value1 {
				if value2 > value1 {
					result.bitsArr[l-i] = (value1 - 1) - (maxValue - value2) + 1
				} else {
					result.bitsArr[l-i] = (value2 - 1) - (maxValue - value1) + 1
				}
				cf = 1
			} else {
				result.bitsArr[l-i] = value1 + value2 + 1
				cf = 0
			}
		} else if cf == 0 {
			if value1 >= maxValue-value2 || value2 >= maxValue-value1 {
				if value2 > value1 {
					result.bitsArr[l-i] = (value1 - 1) - (maxValue - value2)
				} else {
					result.bitsArr[l-i] = (value2 - 1) - (maxValue - value1)
				}
				cf = 1
			} else {
				result.bitsArr[l-i] = value1 + value2
				cf = 0
			}
		}
	}

	return result
}

func (num *Uli) AddZeros(zerosCount int) {
	for j := 0; j < zerosCount; j++ {
		num.bitsArr = append(num.bitsArr, 0)
	}
}

func Div2Blocks(strNum string) []uint64 {
	var len int = len(strNum)
	var blockCount int
	var block string
	var count int
	const blockLen = 16

	if len%blockLen != 0 {
		blockCount = len/blockLen + 1
	} else {
		blockCount = len / blockLen
	}

	for t := 0; t < blockCount*blockLen-len; t++ {
		strNum = "0" + strNum
	}

	blocksArr := make([]uint64, blockCount)

	for i := 0; i < blockCount; i++ {
		for j := count; j < blockLen+count; j++ {
			block += string(strNum[j])
		}

		count += 16
		blocksArr[i], _ = strconv.ParseUint(block, 16, 64)
		block = ""
	}

	return blocksArr
}

func AllZeros(array []uint64) bool {
	l := len(array)
	tmp := make([]uint64, l)
	flag := true

	for i := 0; i < l; i++ {
		if array[i] != tmp[i] {
			flag = false
		}
	}

	return flag
}
