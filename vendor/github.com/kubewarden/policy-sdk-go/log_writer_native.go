//go:build !wasi
// +build !wasi

// note well: we have to use the tinygo wasi target, because the wasm one is
// meant to be used inside of the browser

package sdk

import (
	"fmt"
)

func (k *KubewardenLogWriter) Write(p []byte) (n int, err error) {
	n, err = k.buffer.Write(p)
	line, _ := k.buffer.ReadBytes('\n')
	fmt.Printf("NATIVE: |%s|\n", string(line))
	return
}
