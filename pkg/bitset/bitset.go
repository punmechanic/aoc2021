package bitset

type BitSet []bool

func NewBitSet(bits int) BitSet {
	return make(BitSet, bits)
}

func (b BitSet) BitLength() int {
	return len([]bool(b)) - 1
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

func (b BitSet) Uint() uint {
	var o uint
	shift := b.BitLength()
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
