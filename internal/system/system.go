package system

import (
	"os/exec"
)

func FindExecutablePath(paths []string) (string, error) {
	for _, path := range paths {
		if _, err := exec.LookPath(path); err == nil {
			return path, nil
		}
	}
	return "", exec.ErrNotFound
}
