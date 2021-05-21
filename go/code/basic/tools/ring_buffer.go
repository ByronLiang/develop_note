package tools

import (
	"errors"
)

var (
	ErrInitSize = errors.New("initial size must be great than zero")
	ErrIsEmpty = errors.New("ring buffer is empty")
)

type RingBuffer struct {
	buf         []interface{}
	initialSize int
	size        int
	r           int // read pointer
	w           int // write pointer
}

func NewRingBuffer(initialSize int) (*RingBuffer, error) {
	if initialSize <= 0 {
		return nil, ErrInitSize
	}
	// initial size must >= 2
	if initialSize == 1 {
		initialSize = 2
	}

	return &RingBuffer{
		buf:         make([]interface{}, initialSize),
		initialSize: initialSize,
		size:        initialSize,
	}, nil
}

func (r *RingBuffer) Read() (interface{}, error) {
	if r.IsEmpty() {
		return nil, ErrIsEmpty
	}

	v := r.buf[r.r]
	// 指向未读下标
	r.r++
	if r.r == r.size {
		r.r = 0
	}
	return v, nil
}

func (r *RingBuffer) Pop() (interface{}, error) {
	return r.Read()
}

// 当前读下标的读值
func (r *RingBuffer) Peak() (interface{}, error) {
	if r.IsEmpty() {
		return nil, ErrIsEmpty
	}
	v := r.buf[r.r]
	return v, nil
}

func (r *RingBuffer) ReadAll() ([]interface{}, error) {
	if r.IsEmpty() {
		return nil, ErrIsEmpty
	}
	if r.r > r.w {
		total := r.size - r.r + r.w
		// 写节点已经覆盖读节点处
		container := make([]interface{}, total)
		copy(container[0:], r.buf[r.r:])
		copy(container[r.size-r.r:], r.buf[:r.w])
		r.Reset()
		return container, nil
	}
	rest := r.w - r.r
	if rest == 0 {
		return nil, nil
	}
	all := make([]interface{}, rest)
	copy(all, r.buf[r.r:r.w])
	r.r = 0
	r.w = 0
	return all, nil
}

func (r *RingBuffer) Write(v interface{}) {
	r.buf[r.w] = v
	r.w++

	if r.w == r.size {
		r.w = 0
	}

	if r.w == r.r { // full
		r.grow()
	}
}

func (r *RingBuffer) grow() {
	var size int
	if r.size < 1024 {
		size = r.size * 2
	} else {
		size = r.size + r.size/4
	}

	buf := make([]interface{}, size)

	copy(buf[0:], r.buf[r.r:])
	copy(buf[r.size-r.r:], r.buf[0:r.r])

	r.r = 0
	r.w = r.size
	r.size = size
	r.buf = buf
}

func (r *RingBuffer) IsEmpty() bool {
	return r.r == r.w
}

// Capacity returns the size of the underlying buffer.
func (r *RingBuffer) Capacity() int {
	return r.size
}

func (r *RingBuffer) Len() int {
	if r.r == r.w {
		return 0
	}

	if r.w > r.r {
		return r.w - r.r
	}

	return r.size - r.r + r.w
}

func (r *RingBuffer) Reset() {
	r.r = 0
	r.w = 0
	r.size = r.initialSize
	r.buf = make([]interface{}, r.initialSize)
}
