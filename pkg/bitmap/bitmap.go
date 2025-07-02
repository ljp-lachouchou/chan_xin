package bitmap

type Bitmap struct {
	bits []byte
	size int
}

func NewBitmap(size int) *Bitmap {
	if size == 0 {
		size = 256
	}
	return &Bitmap{
		bits: make([]byte, size),
		size: size * 8,
	}
}
func hash(id string) int {
	// 使用BKDR哈希算法
	seed := 131313 // 31 131 1313 13131 131313, etc
	hash := 0
	for _, c := range id {
		hash = hash*seed + int(c)
	}
	return hash & 0x7FFFFFFF
}

func (b *Bitmap) Set(id string) {
	idHash := hash(id)
	bitIndex := idHash % b.size
	byteIndex := bitIndex / 8
	bit := bitIndex % 8
	b.bits[byteIndex] |= 1 << bit
}
func (b *Bitmap) IsSet(id string) bool {
	idHash := hash(id)
	bitIndex := idHash % b.size
	byteIndex := bitIndex / 8
	bit := bitIndex % 8
	return b.bits[byteIndex]&(1<<bit) != 0
}
func (b *Bitmap) Export() []byte {
	return b.bits
}
func Load(bits []byte) *Bitmap {
	if len(bits) == 0 {
		return NewBitmap(0)
	}
	return &Bitmap{
		bits: bits,
		size: len(bits) * 8,
	}
}
