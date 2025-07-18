package gomock

import "io"

// CountBytes return the total number of bytes read through a reader
func CountBytes(reader io.Reader) (total int, err error) {
	var (
		n int

		buffer = make([]byte, 1024)
	)
	for err == nil {
		n, err = reader.Read(buffer)
		total += n
	}
	if err == io.EOF {
		err = nil
	}

	// Variables are optional in this case. Could be replaced with just `return`
	return total, err
}
