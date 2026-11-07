# OMC Community Plugin Registry

The official plugin registry for the Oh-My-CPA (OMC) and CLIProxyAPI (CPA) ecosystem.

This repository publishes a registry of vetted plugins and integrations for CPA and OMC, allowing users to discover and install extensions easily from the OMC console or directly in CPA.

## Registry URL

```
https://raw.githubusercontent.com/omc-community/omc-plugins/main/registry.json
```

## How to Register

### In Oh-My-CPA (OMC)

Oh-My-CPA connects to this source by default. You can manage source availability under **Settings → Plugins → Plugin Store**.

### In CLIProxyAPI (CPA)

Add the source to your `config.yaml`:

```yaml
plugins:
  enabled: true
  store-sources:
    - https://raw.githubusercontent.com/omc-community/omc-plugins/main/registry.json
```

## Plugin Entries

| ID | Name | Description | Author | Repository |
| --- | --- | --- | --- | --- |
| `omc-zcode-oauth` | OMC ZCode OAuth | Native ZCode provider with dual-portal OAuth login, model routing, and quota normalization. | JUNZE | [omc-community/omc-zcode-oauth](https://github.com/omc-community/omc-zcode-oauth) |
| `omc-qoder-oauth` | OMC Qoder OAuth | Native Qoder & Qoder CN OAuth provider with OpenAI-compatible execution and quota. | JUNZE | [omc-community/omc-qoder-oauth](https://github.com/omc-community/omc-qoder-oauth) |
| `omcpa-codex-candy-eval` | 降智测试 | Model performance evaluation plugin (candy reasoning, model fingerprints, pelican tests). | Yorick-Ryu / haowang02 | [Yorick-Ryu/cpa-plugin-codex-candy-eval](https://github.com/Yorick-Ryu/cpa-plugin-codex-candy-eval) |

## Contributing a Plugin

We welcome community plugins and extensions! To submit a plugin to this registry:

1. Fork this repository.
2. Add your plugin entry to `registry.json`.
3. Ensure every entry adheres to the requirements:
   - `id`: Unique identifier matching `^[A-Za-z0-9][A-Za-z0-9._-]{0,127}$`.
   - `name`: Human-readable display name.
   - `description`: Concise summary of what the plugin does.
   - `author`: Author or organization name.
   - `repository`: Exact `https://github.com/{owner}/{repo}` URL (no shorthand, no `.git` suffix).
   - `version`: Display fallback version matching `^[0-9][0-9A-Za-z.+-]*$` (no leading `v`).
   - `install`: `{"type": "github-release"}`.
   - `locales`: Optional localized descriptions (`zh-CN`, `zh-Hant`, `ms`, etc.).
   - `logo`: Optional direct HTTPS image URL or data URI.
4. Run `go test ./...` to verify schema validity.
5. Open a Pull Request.

## License

[MIT](LICENSE)
