package main

import (
	"encoding/binary"
	"sync"
)

func Solution(f func(uint8) uint8, inp uint32) uint32 {
	in := make([]byte, 4)
	binary.LittleEndian.PutUint32(in, inp)
	out := make([]byte, 4)
	var wg sync.WaitGroup
	for i := range in {
		wg.Add(1)
		go func(in uint8, idx int) {
			defer wg.Done()
			out[idx] = f(in)
		}(in[i], i)
	}
	wg.Wait()
	return binary.LittleEndian.Uint32(out)
}
