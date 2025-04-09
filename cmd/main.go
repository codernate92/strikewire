package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"Strikewire/internal/fetcher"
	"Strikewire/internal/mitre"
)

func main() {
	actor := flag.String("actor", "", "Name of threat actor (e.g., APT28, Lapsus$)")
	out := flag.String("out", "reports", "Output directory")
	flag.Parse()

	if *actor == "" {
		fmt.Println("Usage: strikewire -actor [NAME] -out [DIR]")
		os.Exit(1)
	}

	// Create output dir if missing
	os.MkdirAll(*out, os.ModePerm)

	// Load raw profile (struct) from local map
	profile, ok := fetcher.Profiles[*actor]
	if !ok {
		log.Fatalf("❌ No threat profile found for actor: %s", *actor)
	}

	// Load MITRE TTPs
	ttpMap, err := mitre.LoadTTPs("data/mitre_ttp.json")
	if err != nil {
		log.Fatalf("❌ Failed to load MITRE TTP data: %v", err)
	}

	// JSON Report
	jsonPath := filepath.Join(*out, fmt.Sprintf("%s.json", *actor))
	jsonString, err := fetcher.FetchThreatIntel(*actor)
	if err != nil {
		log.Fatalf("❌ Error generating JSON: %v", err)
	}
	os.WriteFile(jsonPath, []byte(jsonString), 0644)

	// Markdown Report
	mdPath := filepath.Join(*out, fmt.Sprintf("%s.md", *actor))
	err = fetcher.GenerateMarkdownReport(*actor, profile, ttpMap, mdPath)
	if err != nil {
		log.Fatalf("❌ Failed to write markdown: %v", err)
	}

	fmt.Printf("✅ STRIKEWIRE reports saved to:\n- %s\n- %s\n", jsonPath, mdPath)
}
