package wavefront

import (
	"fmt"
	"time"

	"github.com/masters-of-cats/cfbench/bench"
)

func BuildWavefrontOutput(hostName string, phases bench.Phases) string {
	timeOfTest := time.Now().Unix()
	result := ""
	for _, phase := range phases {
		if !phase.IsValid() {
			continue
		}
		result += fmt.Sprintf("cfbench.%s %d %d source=%s\n", phase.ShortName, int64(phase.Duration()), timeOfTest, hostName)
	}
	return result
}
