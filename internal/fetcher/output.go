package fetcher

import (
	"fmt"
	"os"
	"strings"

	"Strikewire/internal/mitre"
)

func GenerateMarkdownReport(actor string, profile ThreatProfile, ttpMap map[string]mitre.TTP, filename string) error {
	var sb strings.Builder

	sb.WriteString(fmt.Sprintf("# STRIKEWIRE Report: %s\n\n", actor))

	sb.WriteString("## 🧬 Malware\n")
	for _, m := range profile.Malware {
		sb.WriteString(fmt.Sprintf("- %s\n", m))
	}
	sb.WriteString("\n")

	sb.WriteString("## 🌐 Infrastructure\n")
	for _, infra := range profile.Infrastructure {
		sb.WriteString(fmt.Sprintf("- `%s`\n", infra))
	}
	sb.WriteString("\n")

	sb.WriteString("## 🎯 Tactics, Techniques, and Procedures (MITRE ATT&CK)\n")
	for _, ttpID := range profile.TTPs {
		if ttp, exists := ttpMap[ttpID]; exists {
			sb.WriteString(fmt.Sprintf("- **%s** (%s): %s\n", ttp.Name, ttp.ID, ttp.Description))
		} else {
			sb.WriteString(fmt.Sprintf("- %s: _No description found_\n", ttpID))
		}
	}
	sb.WriteString("\n")

	sb.WriteString("## 📄 Intelligence Reports\n")
	for _, r := range profile.Reports {
		sb.WriteString(fmt.Sprintf("- [%s](%s)\n", r, r))
	}

	// Write markdown to file
	err := os.WriteFile(filename, []byte(sb.String()), 0644)
	if err != nil {
		return err
	}

	return nil
}
