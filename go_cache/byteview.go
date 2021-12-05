package go_cache

// ByteView 表示缓存值类型
type ByteView struct {
	b []byte
}

// Len return the view's length
func (v ByteView) Len() int {
	return len(v.b)
}

// ByteSlice return a copy of the data as a byte slice
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

// String return the data as a string
func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
