package mitre

import (
	"encoding/json"
	"os"
)

type TTP struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func LoadTTPs(filename string) (map[string]TTP, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var ttpList []TTP
	err = json.Unmarshal(data, &ttpList)
	if err != nil {
		return nil, err
	}

	ttpMap := make(map[string]TTP)
	for _, ttp := range ttpList {
		ttpMap[ttp.ID] = ttp
	}
	return ttpMap, nil
}
