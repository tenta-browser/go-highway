// +build !amd64

package highway

func hashSSE(keys, init0, init1 *Lanes, p []byte) uint64 {
	return 0
}
