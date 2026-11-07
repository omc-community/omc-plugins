package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type pluginEntry struct {
	ID          string                     `json:"id"`
	Name        string                     `json:"name"`
	Description string                     `json:"description"`
	Author      string                     `json:"author"`
	Version     string                     `json:"version,omitempty"`
	Repository  string                     `json:"repository"`
	Logo        string                     `json:"logo,omitempty"`
	Homepage    string                     `json:"homepage,omitempty"`
	License     string                     `json:"license,omitempty"`
	Tags        []string                   `json:"tags,omitempty"`
	Locales     map[string]json.RawMessage `json:"locales,omitempty"`
	Install     json.RawMessage            `json:"install,omitempty"`
}

type registryDoc struct {
	SchemaVersion int           `json:"schema_version"`
	Plugins       []pluginEntry `json:"plugins"`
}

func main() {
	if err := generate(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func generate() error {
	officialBytes, err := os.ReadFile("official.json")
	if err != nil {
		return fmt.Errorf("read official.json: %w", err)
	}
	communityBytes, err := os.ReadFile("community.json")
	if err != nil {
		return fmt.Errorf("read community.json: %w", err)
	}

	var officialDoc registryDoc
	if err := json.Unmarshal(officialBytes, &officialDoc); err != nil {
		return fmt.Errorf("decode official.json: %w", err)
	}

	var communityDoc registryDoc
	if err := json.Unmarshal(communityBytes, &communityDoc); err != nil {
		return fmt.Errorf("decode community.json: %w", err)
	}

	combined := registryDoc{
		SchemaVersion: 2,
		Plugins:       make([]pluginEntry, 0, len(officialDoc.Plugins)+len(communityDoc.Plugins)),
	}

	// Invariant: Official plugins are appended first and must ALWAYS appear
	// above community plugins in registry.json.
	seen := make(map[string]bool)
	for _, p := range officialDoc.Plugins {
		if p.ID != "" {
			seen[p.ID] = true
		}
		combined.Plugins = append(combined.Plugins, p)
	}

	for _, p := range communityDoc.Plugins {
		if p.ID != "" && seen[p.ID] {
			return fmt.Errorf("duplicate plugin id between official and community: %s", p.ID)
		}
		if p.ID != "" {
			seen[p.ID] = true
		}
		combined.Plugins = append(combined.Plugins, p)
	}

	out, err := json.MarshalIndent(combined, "", "  ")
	if err != nil {
		return fmt.Errorf("marshal registry.json: %w", err)
	}
	out = append(out, '\n')

	if err := os.WriteFile("registry.json", out, 0o644); err != nil {
		return fmt.Errorf("write registry.json: %w", err)
	}
	fmt.Printf("Successfully generated registry.json with %d plugins (%d official, %d community)\n",
		len(combined.Plugins), len(officialDoc.Plugins), len(communityDoc.Plugins))
	return nil
}
