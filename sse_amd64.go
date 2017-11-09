// +build amd64

package highway

import (
	"github.com/intel-go/cpuid"
)

func useSSE() bool {
	return cpuid.HasFeature(cpuid.SSE4_1)
}
