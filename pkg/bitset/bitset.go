package bitset

const one rune = '1'

func uniform(sets []BitSet, tiebreaker func(a []BitSet, b []BitSet) []BitSet) BitSet {
	for i := 0; i < sets[0].BitLength(); i++ {
		var ons, offs []BitSet
		if len(sets) == 1 {
			break
		}

		for _, set := range sets {
			if set[i] {
				ons = append(ons, set)
			} else {
				offs = append(offs, set)
			}
		}

		sets = tiebreaker(ons, offs)
	}

	return sets[0]
}

// Uniform attempts to find the most uniform BitSet from the provided set of sets.
//
// This differs from MostCommon by attempting to find the most uniform BitSet of the given sets rather than constructing a new BitSet to hold the most common bit.
//
// For example, with a set of BitSets with the following bits set:
//
//		0b1101 0b1001 0b0000 0b0010 0b0011
//
// The BitSet 0b0010 would be returned, because it has the most in common with the rest of the BitSets.
// One can adjust the tie-breaking behaviour (which defaults to prefering 1s over 0s) by passing in a tie-breaking function in the optional argument.
func Uniform(sets []BitSet) BitSet {
	return uniform(sets, func(a []BitSet, b []BitSet) []BitSet {
		if len(a) >= len(b) {
			return a
		} else {
			return b
		}
	})
}

// UniformZeroes attempts to find the most uniform BitSet from the provided set of sets.
//
// This is similar to Uniform but prefers zeroes instead of ones, in the event of a tie.
func UniformZeroes(sets []BitSet) BitSet {
	return uniform(sets, func(a []BitSet, b []BitSet) []BitSet {
		if len(a) >= len(b) {
			return b
		} else {
			return a
		}
	})
}

// MostCommon returns a BitSet that represents the most common bits in the given slice of BitSets.
//
// For example, a slice containing BitSets 1001, 1100, 1011 and 0001 would return the BitSet 1001.
// If there is a tie for a particular bit position, MostCommon will prefer 1s to 0s.
//
// The size of the ultimate BitSet is based on the size of the first in the slice; each bitset should be aligned.
// The behaviour of a slice of unaligned bitsets is undefined.
func MostCommon(sets []BitSet) BitSet {
	bitLen := sets[0].BitLength()
	bs := NewBitSet(bitLen)
	for i := 0; i < bitLen; i++ {
		count := 0
		for _, set := range sets {
			if set[i] {
				count++
			}
		}

		if count >= len(sets)/2 {
			bs.Set(i)
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
