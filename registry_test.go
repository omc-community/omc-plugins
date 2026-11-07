package main

import (
	"encoding/json"
	"fmt"
	"net/url"
	"os"
	"regexp"
	"strings"
	"testing"
)

type registryDocument struct {
	SchemaVersion int              `json:"schema_version"`
	Plugins       []registryPlugin `json:"plugins"`
}

type registryPlugin struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Author      string `json:"author"`
	Version     string `json:"version"`
	Repository  string `json:"repository"`
	Install     struct {
		Type string `json:"type"`
	} `json:"install"`
}

const (
	installTypeGitHubRelease = "github-release"
	installTypeDirect        = "direct"
)

var (
	pluginIDPattern      = regexp.MustCompile(`^[A-Za-z0-9][A-Za-z0-9._-]{0,127}$`)
	pluginVersionPattern = regexp.MustCompile(`^[0-9][0-9A-Za-z.+-]*$`)
)

func validateRegistry(raw []byte) (registryDocument, error) {
	var document registryDocument
	if errDecode := json.Unmarshal(raw, &document); errDecode != nil {
		return document, fmt.Errorf("decode registry: %w", errDecode)
	}
	if document.SchemaVersion != 1 && document.SchemaVersion != 2 {
		return document, fmt.Errorf("unsupported schema_version %d", document.SchemaVersion)
	}
	seen := make(map[string]struct{}, len(document.Plugins))
	for index, plugin := range document.Plugins {
		if document.SchemaVersion == 1 && pluginInstallType(plugin) == installTypeDirect {
			return document, fmt.Errorf("plugins[%d]: direct install requires schema_version 2", index)
		}
		if errPlugin := validatePlugin(plugin); errPlugin != nil {
			return document, fmt.Errorf("plugins[%d]: %w", index, errPlugin)
		}
		id := strings.TrimSpace(plugin.ID)
		if _, exists := seen[id]; exists {
			return document, fmt.Errorf("plugins[%d]: duplicate plugin id %q", index, id)
		}
		seen[id] = struct{}{}
	}
	return document, nil
}

func pluginInstallType(plugin registryPlugin) string {
	if installType := strings.ToLower(strings.TrimSpace(plugin.Install.Type)); installType != "" {
		return installType
	}
	return installTypeGitHubRelease
}

func validatePlugin(plugin registryPlugin) error {
	installType := pluginInstallType(plugin)
	required := map[string]string{
		"id": plugin.ID, "name": plugin.Name, "description": plugin.Description, "author": plugin.Author,
	}
	if installType == installTypeGitHubRelease {
		required["repository"] = plugin.Repository
	}
	for field, value := range required {
		if strings.TrimSpace(value) == "" {
			return fmt.Errorf("missing required field %s", field)
		}
	}
	if !pluginIDPattern.MatchString(strings.TrimSpace(plugin.ID)) {
		return fmt.Errorf("invalid id %q", plugin.ID)
	}
	if version := strings.TrimSpace(plugin.Version); version != "" {
		if strings.HasPrefix(version, "v") || !pluginVersionPattern.MatchString(version) {
			return fmt.Errorf("invalid version %q", plugin.Version)
		}
	}
	switch installType {
	case installTypeGitHubRelease:
		return validateRepository(plugin.Repository)
	case installTypeDirect:
		return nil
	default:
		return fmt.Errorf("unsupported install type %q", plugin.Install.Type)
	}
}

func validateRepository(repository string) error {
	parsed, errParse := url.Parse(strings.TrimSpace(repository))
	if errParse != nil {
		return fmt.Errorf("invalid repository URL: %w", errParse)
	}
	if parsed.Scheme != "https" || parsed.Host != "github.com" || parsed.RawQuery != "" || parsed.Fragment != "" {
		return fmt.Errorf("repository must be https://github.com/{owner}/{repo}")
	}
	segments := strings.Split(strings.Trim(parsed.EscapedPath(), "/"), "/")
	if len(segments) != 2 || segments[0] == "" || segments[1] == "" {
		return fmt.Errorf("repository must be https://github.com/{owner}/{repo}")
	}
	repo, errUnescape := url.PathUnescape(segments[1])
	if errUnescape != nil {
		return fmt.Errorf("invalid repository name: %w", errUnescape)
	}
	if strings.HasSuffix(repo, ".git") {
		return fmt.Errorf("repository must be https://github.com/{owner}/{repo}")
	}
	return nil
}

func TestRegistriesAreValid(t *testing.T) {
	files := []string{"official.json", "community.json", "registry.json"}
	for _, file := range files {
		t.Run(file, func(t *testing.T) {
			raw, err := os.ReadFile(file)
			if err != nil {
				t.Fatalf("failed to read %s: %v", file, err)
			}
			doc, err := validateRegistry(raw)
			if err != nil {
				t.Fatalf("%s validation failed: %v", file, err)
			}
			if len(doc.Plugins) == 0 {
				t.Fatalf("%s: expected plugins list to be non-empty", file)
			}
		})
	}
}

func TestRegistryMatchesGenerated(t *testing.T) {
	officialBytes, err := os.ReadFile("official.json")
	if err != nil {
		t.Fatalf("read official.json: %v", err)
	}
	communityBytes, err := os.ReadFile("community.json")
	if err != nil {
		t.Fatalf("read community.json: %v", err)
	}
	registryBytes, err := os.ReadFile("registry.json")
	if err != nil {
		t.Fatalf("read registry.json: %v", err)
	}

	var officialDoc registryDocument
	if err := json.Unmarshal(officialBytes, &officialDoc); err != nil {
		t.Fatalf("decode official.json: %v", err)
	}
	var communityDoc registryDocument
	if err := json.Unmarshal(communityBytes, &communityDoc); err != nil {
		t.Fatalf("decode community.json: %v", err)
	}
	var registryDoc registryDocument
	if err := json.Unmarshal(registryBytes, &registryDoc); err != nil {
		t.Fatalf("decode registry.json: %v", err)
	}

	expectedCount := len(officialDoc.Plugins) + len(communityDoc.Plugins)
	if len(registryDoc.Plugins) != expectedCount {
		t.Fatalf("registry.json has %d plugins, want %d (official %d + community %d)",
			len(registryDoc.Plugins), expectedCount, len(officialDoc.Plugins), len(communityDoc.Plugins))
	}

	officialIDs := make(map[string]bool)
	for _, p := range officialDoc.Plugins {
		officialIDs[p.ID] = true
	}
	communityIDs := make(map[string]bool)
	for _, p := range communityDoc.Plugins {
		if officialIDs[p.ID] {
			t.Fatalf("plugin ID %q exists in both official.json and community.json", p.ID)
		}
		communityIDs[p.ID] = true
	}

	// Invariant: Official plugins must ALWAYS appear above community plugins.
	for i, officialPlugin := range officialDoc.Plugins {
		if registryDoc.Plugins[i].ID != officialPlugin.ID {
			t.Fatalf("registry.json[%d] = %q, want official plugin %q (official plugins must always appear above community plugins)",
				i, registryDoc.Plugins[i].ID, officialPlugin.ID)
		}
	}

	offset := len(officialDoc.Plugins)
	for j, communityPlugin := range communityDoc.Plugins {
		idx := offset + j
		if registryDoc.Plugins[idx].ID != communityPlugin.ID {
			t.Fatalf("registry.json[%d] = %q, want community plugin %q",
				idx, registryDoc.Plugins[idx].ID, communityPlugin.ID)
		}
	}
}
