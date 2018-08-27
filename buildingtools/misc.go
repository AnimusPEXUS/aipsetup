package buildingtools

import "strings"

func FilterAutotoolsConfigOptions(
	args []string,
	exact []string,
	prefixes []string,
) ([]string, error) {

filtering:
	for i := len(args) - 1; i != -1; i -= 1 {

		for _, j := range exact {
			if args[i] == j {
				args = append(args[:i], args[i+1:]...)
				continue filtering
			}
		}

		for _, j := range prefixes {
			if strings.HasPrefix(args[i], j) {
				args = append(args[:i], args[i+1:]...)
				continue filtering
			}
		}
	}

	return args, nil
}
