//go:build amd64 || arm64 || loong64 || mips64 || mips64le || ppc64 || ppc64le || riscv64 || s390x || wasm

package netpoll

func nativeUint(v int) uint64 {
	return uint64(v)
}
