package xtests

import (
	"os"

	"github.com/cockroachdb/errors"
)

func SetupASDFDataDir(path string) (func(), error) {
	return envSetter(map[string]string{
		"ASDF_DATA_DIR": path,
	})
}

func envSetter(envs map[string]string) (func(), error) {
	originals := make(map[string]string, len(envs))

	for name, value := range envs {
		if val, ok := os.LookupEnv(name); ok {
			originals[name] = val
		}
		if err := os.Setenv(name, value); err != nil {
			return nil, errors.Wrapf(err, "failed to set %s", name)
		}
	}

	return func() {
		for name := range envs {
			if original, ok := originals[name]; ok {
				_ = os.Setenv(name, original)
			} else {
				_ = os.Unsetenv(name)
			}
		}
	}, nil
}
