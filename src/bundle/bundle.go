package bundle

import (
	"github.com/severinraez/cotenoer/inventory"
	"os/exec"
	"strings"
)


type Overview struct {
	Name string
	ActiveContainers int
}

func GetOverview(bundle inventory.Bundle) Overview {
	psCmd := exec.Command("docker-compose", "-f", bundle.ComposeFilePath, "ps", "-q")

	// output, err := psCmd.CombinedOutput()
	output, err := psCmd.Output()
	panicOnError(err)

	containerCount := len(strings.Split(string(output), "\n")) - 1

	return Overview{
		Name: bundle.Name,
		ActiveContainers: containerCount}
}

func panicOnError(err error) {
	if(err != nil) {
		panic(err)
	}
}
