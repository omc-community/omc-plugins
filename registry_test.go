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

func TestRegistryIsValid(t *testing.T) {
	raw, err := os.ReadFile("registry.json")
	if err != nil {
		t.Fatalf("failed to read registry.json: %v", err)
	}
	doc, err := validateRegistry(raw)
	if err != nil {
		t.Fatalf("registry.json validation failed: %v", err)
	}
	if len(doc.Plugins) == 0 {
		t.Fatalf("expected plugins list to be non-empty")
	}
}
