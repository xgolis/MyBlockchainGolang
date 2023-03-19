package thirdphase

import "hash/fnv"

type ByteArrayWrapper struct {
	Contents []byte
}

func (b *ByteArrayWrapper) ByteArrayWrapper(c []byte) {
	b.Contents = []byte{}
	for i := 0; i < len(c); i++ {
		b.Contents[i] = c[i]
	}
}

func (b *ByteArrayWrapper) equals(other ByteArrayWrapper) bool {
	c := other.Contents
	if b.Contents == nil {
		if c == nil {
			return true
		} else {
			return false
		}
	} else {
		if len(b.Contents) != len(c) {
			return false
		}
		for i := 0; i < len(c); i++ {
			if b.Contents[i] != c[i] {
				return false
			}
		}
		return true
	}

}

func (b *ByteArrayWrapper) hashCode() int {
	h := fnv.New32a()
	h.Write(b.Contents)
	return int(h.Sum32())
}
