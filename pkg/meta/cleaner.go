package meta

import (
	"io"

	"vimagination.zapto.org/dos2unix"
)

type clean struct {
	r           io.Reader
	lastWasChar bool
}

func (c *clean) Read(p []byte) (int, error) {
	n, err := c.r.Read(p)
	p = p[:n]
	for i := 0; i < len(p); i++ {
		if p[i] == '\n' {
			if c.lastWasChar {
				c.lastWasChar = false
			} else {
				p = append(p[:i], p[i+1:]...)
				i--
			}
		} else {
			c.lastWasChar = true
		}
	}
	if err != nil {
		return len(p), err
	}
	return len(p), nil
}

func cleaner(r io.Reader) io.Reader {
	return &clean{
		r: dos2unix.DOS2Unix(r),
	}
}
