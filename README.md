<p align="center">
  <img src="assets/logo.png" width="96" height="96" alt="OMC Plugins logo">
</p>

<h1 align="center">OMC Community Plugin Registry</h1>

<p align="center">
  <a href="https://github.com/WizisCool/oh-my-cpa">
    <picture>
      <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/WizisCool/oh-my-cpa/main/web/src/assets/brand/omc-wordmark-dark.svg">
      <img src="https://raw.githubusercontent.com/WizisCool/oh-my-cpa/main/web/src/assets/brand/omc-wordmark-light.svg" alt="Oh-My-CPA" height="22">
    </picture>
  </a>
</p>

<p align="center">
  Unified plugin registry source for
  <a href="https://github.com/WizisCool/oh-my-cpa">Oh-My-CPA</a> and
  <a href="https://github.com/router-for-me/CLIProxyAPI">CLIProxyAPI</a>
</p>

<p align="center">
  <a href="https://github.com/omc-community/omc-plugins/actions/workflows/validate.yml"><img src="https://github.com/omc-community/omc-plugins/actions/workflows/validate.yml/badge.svg" alt="CI"></a>
  <img src="https://img.shields.io/badge/CLIProxyAPI-v8+-4f46e5" alt="CLIProxyAPI v8+">
  <img src="https://img.shields.io/badge/Oh--My--CPA-plugin--registry-0284c7" alt="Oh-My-CPA plugin-registry">
  <img src="https://img.shields.io/badge/license-MIT-blue" alt="MIT license">
</p>

<p align="center">
  <a href="README.md">English</a> · <a href="README.zh-CN.md">简体中文</a>
</p>

This repository publishes the official plugin registry document (`registry.json`) for the
[Oh-My-CPA](https://github.com/WizisCool/oh-my-cpa) (OMC) and
[CLIProxyAPI](https://github.com/router-for-me/CLIProxyAPI) (CPA) ecosystem.

It provides a curated, vetted software source for official provider extensions, companion
tools, and community plugins, enabling zero-configuration plugin discovery and one-click
installation inside Oh-My-CPA and CPA.

---

## Registry URL

```
https://raw.githubusercontent.com/omc-community/omc-plugins/main/registry.json
```

---

## Configuration & Usage

### Oh-My-CPA (OMC)

Oh-My-CPA includes this registry by default as its official ecosystem source.
You can view, enable, or disable it under **Settings → Plugins → Plugin Store**.

### CLIProxyAPI (CPA)

Add the raw registry URL to the `plugins.store-sources` list in your `config.yaml`:

```yaml
plugins:
  enabled: true
  store-sources:
    - https://raw.githubusercontent.com/omc-community/omc-plugins/main/registry.json
```

---

## Catalogued Plugins

| Plugin ID | Name | Description | Author | Repository |
| --- | --- | --- | --- | --- |
| `omc-zcode-oauth` | **OMC ZCode OAuth** | Native ZCode provider with dual-portal Z.ai / BigModel OAuth, model discovery, and quota normalization. | JUNZE | [omc-community/omc-zcode-oauth](https://github.com/omc-community/omc-zcode-oauth) |
| `omc-qoder-oauth` | **OMC Qoder OAuth** | Native Qoder & Qoder CN provider with dual-portal OAuth, OpenAI-compatible execution, and quota normalization. | JUNZE | [omc-community/omc-qoder-oauth](https://github.com/omc-community/omc-qoder-oauth) |
| `omcpa-codex-candy-eval` | **降智测试** | Model performance evaluation plugin (candy reasoning, model fingerprints, and pelican rendering). | Yorick-Ryu / haowang02 | [Yorick-Ryu/cpa-plugin-codex-candy-eval](https://github.com/Yorick-Ryu/cpa-plugin-codex-candy-eval) |

---

## Contributing a Plugin

We welcome extensions, OAuth providers, and tools for the Oh-My-CPA ecosystem! To add a plugin to this registry:

1. **Fork** this repository.
2. Edit `registry.json` and add your entry under the `plugins` array.
3. Validate your entry against CPA's registry requirements:
   - `id`: Unique identifier matching `^[A-Za-z0-9][A-Za-z0-9._-]{0,127}$`.
   - `name`: Clean human-readable display name.
   - `description`: Single clear sentence summarizing capability.
   - `author`: Maintainer name or organization.
   - `repository`: Full GitHub HTTPS address `https://github.com/{owner}/{repo}` (no `.git` suffix, no shorthand).
   - `version`: Display fallback version matching `^[0-9][0-9A-Za-z.+-]*$` (no leading `v`).
   - `install`: `{"type": "github-release"}`.
   - `locales`: Optional localized descriptions (`zh-CN`, `zh-Hant`, `ms`, etc.).
   - `logo`: Direct HTTPS URL to artwork or SVG/PNG asset.
4. Run validation tests:
   ```bash
   go test -v ./...
   ```
5. Submit a **Pull Request**.

---

## License

[MIT](LICENSE)
