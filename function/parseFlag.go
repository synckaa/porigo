package function

import (
	"strconv"
	"strings"
)

func ParseFlag(port string) []int {

	if strings.Contains(port, "-") {
		parts := strings.Split(port, "-")
		if len(parts) != 2 {
			return nil
		}

		start, err := strconv.Atoi(parts[0])
		if err != nil {
			return nil
		}
		end, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil
		}

		return []int{start, end}

	}
	intPort, err := strconv.Atoi(port)
	if err != nil {
		return nil
	}

	return []int{intPort, intPort}
}
