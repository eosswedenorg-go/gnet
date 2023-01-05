//go:build 386 || arm || mips || mipsle

package netpoll

func nativeUint(v int) uint32 {
	return uint32(v)
}
