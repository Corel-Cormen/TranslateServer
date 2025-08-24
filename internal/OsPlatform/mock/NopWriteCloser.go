package MockOsPlatformApi

import "bytes"

type NopWriteCloser struct {
	*bytes.Buffer
}

func (n *NopWriteCloser) Close() error {
	return nil
}
