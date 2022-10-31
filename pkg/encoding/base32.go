package encoding

const numLetters = 26
const mask byte = 0b11111 << 3

func Map5BitBlockToValue(bits byte) rune {
	bits = bits & 0b11111

	if bits < numLetters {
		bits += byte('A')
	} else {
		bits += byte('2') - numLetters
	}

	return rune(bits)
}

func Map5ByteBlockInto5BitGroups(block []byte) [8]byte {
	var out [8]byte

	currByte := 0
	currBit := 0
	for i := range out {
		// get the relevant bits out of the current byte
		var currByteContents byte = 0
		if currByte < len(block) {
			currByteContents = block[currByte]
		}
		var newval byte = (currByteContents & (mask >> currBit))

		// get remaining bits from next byte
		if currBit >= 3 {
			newval <<= currBit - 3
			currByte++
			currByteContents = 0
			if currByte < len(block) {
				currByteContents = block[currByte]
			}
			var remainder byte = (currByteContents & (mask << (8 - currBit)))
			remainder >>= (11 - currBit)

			newval |= remainder
		} else {
			newval >>= 3 - currBit
		}

		out[i] = newval

		currBit = (currBit + 5) % 8
	}

	return out
}
