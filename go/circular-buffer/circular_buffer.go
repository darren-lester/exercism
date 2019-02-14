package circular

import "fmt"

// Buffer is a circular buffer
type Buffer struct {
	length       int
	value        []byte
	cursor       int
	oldestCursor int
}

// NewBuffer creates a buffer of a defined size
func NewBuffer(size int) *Buffer {
	return &Buffer{0, make([]byte, size), 0, 0}
}

// ReadByte reads the oldest value written to the buffer
func (buffer *Buffer) ReadByte() (byte, error) {
	if buffer.length == 0 {
		return 0, fmt.Errorf("buffer is empty")
	}

	oldestValue := buffer.value[buffer.oldestCursor]
	buffer.value[buffer.oldestCursor] = 0
	buffer.length--
	shiftCursor(buffer, &buffer.oldestCursor)

	return oldestValue, nil
}

// WriteByte writes a value to the buffer
func (buffer *Buffer) WriteByte(c byte) error {
	if buffer.Full() {
		return fmt.Errorf("buffer is full")
	}

	buffer.value[buffer.cursor] = c
	shiftCursor(buffer, &buffer.cursor)
	buffer.length++

	return nil
}

// Overwrite writes a value to the buffer.
// If the buffer is full it overwrites the oldest value
// written to the buffer.
func (buffer *Buffer) Overwrite(c byte) {
	if buffer.Full() {
		buffer.value[buffer.oldestCursor] = c
		shiftCursor(buffer, &buffer.oldestCursor)
	} else {
		buffer.WriteByte(c)
	}
}

// Reset resets the buffer to its initial state
func (buffer *Buffer) Reset() {
	for i := range buffer.value {
		buffer.value[i] = 0
	}
	buffer.length = 0
}

// Full checks whether the buffer is full
func (buffer *Buffer) Full() bool {
	return buffer.length == cap(buffer.value)
}

func shiftCursor(buffer *Buffer, cursor *int) {
	*cursor = (*cursor + 1) % cap(buffer.value)
}
