package fetcher

import (
	"encoding/json"
	"fmt"
)

func FetchThreatIntel(profileName string) (string, error) {
	profile, exists := Profiles[profileName]
	if !exists {
		return "", fmt.Errorf("profile %s not found", profileName)
	}
	reportBytes, err := json.MarshalIndent(profile, "", "  ")
	if err != nil {
		return "", err
	}

	return string(reportBytes), nil
}


