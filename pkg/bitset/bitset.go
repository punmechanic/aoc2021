package bitset

const one rune = '1'

// MostCommon returns a BitSet that represents the most common bits in the given slice of BitSets.
// For example, a slice containing BitSets 1001, 1100, 1011 and 0001 would return the BitSet 1001.
// The size of the ultimate BitSet is based on the size of the first in the slice; each bitset should be aligned.
// The behaviour of a slice of unaligned bitsets is undefined.
func MostCommon(sets []BitSet) BitSet {
	bitLen := sets[0].BitLength()
	oneCounts := make([]int, len(sets))
	for _, set := range sets {
		for i := 0; i < bitLen; i++ {
			if set[i] {
				oneCounts[i]++
			}
		}
	}

	bs := NewBitSet(bitLen)
	half := len(sets) / 2
	for idx, count := range oneCounts {
		if count > half {
			bs.Set(idx)
		}
	}

	return bs
}

// FromString reads a BitSet from a string by converting each character to a number - 1s and 0s are the only correct input.
// Any other character will be ignored.
func FromString(str string) BitSet {
	bs := NewBitSet(len(str))
	for idx, char := range str {
		if char == one {
			bs.Set(idx)
		}
	}

	return bs
}

type BitSet []bool

func NewBitSet(bits int) BitSet {
	return make(BitSet, bits)
}

func (b BitSet) BitLength() int {
	return len([]bool(b))
}

func (b BitSet) Set(idx int) {
	if idx > b.BitLength() || idx < 0 {
		panic("index out of range")
	}

	b[idx] = true
}

func (b BitSet) Unset(idx int) {
	if idx > b.BitLength() || idx < 0 {
		panic("index out of range")
	}

	b[idx] = false
}

func (b BitSet) Negate() BitSet {
	bs := NewBitSet(b.BitLength())
	for idx, bit := range b {
		bs[idx] = !bit
	}

	return bs
}

func (b BitSet) Uint() uint {
	var o uint
	shift := b.BitLength() - 1
	for _, bit := range b {
		var b uint8
		if bit == true {
			b = 1
		}

		shifted := uint(b) << shift
		o |= shifted
		shift--
	}

	return o
}
