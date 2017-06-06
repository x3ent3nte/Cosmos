package bit

import (
)

func IsBitOneAt(value int, pos uint) bool {
	return (value & (1 << pos)) != 0
}

func SetBitOneAt(value int, pos uint) int {
	return value | (1 << pos)
}

func SetBitZeroAt(value int, pos uint) int {
	return value & (^(1 << pos))
}

func ToggleBitAt(value int, pos uint) int {
	return value ^ (1 << pos)
}