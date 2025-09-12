package assert

import (
	"bufio"
	"bytes"
)

func NewReaderWriter() (*bufio.Reader, *bufio.Writer) {
	buffer := bytes.NewBuffer(nil)
	return bufio.NewReader(buffer),
		bufio.NewWriter(buffer)
}
