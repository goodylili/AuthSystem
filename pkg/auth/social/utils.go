package social

import (
	"os"
)

func getEnVars() []string {
	envVars := []string{"GOOGLE_CLIENT_ID", "GOOGLE_CLIENT_SECRET", "GITHUB_CLIENT_ID", "GITHUB_CLIENT_SECRET"}
	for _, envVar := range envVars {
		if os.Getenv(envVar) == "" {
			return nil

		}
	}

	return envVars

}
