<p align="center">
  <img src="assets/logo.png" width="96" height="96" alt="OMC Plugins logo">
</p>

<h1 align="center">OMC 社区插件源</h1>

<p align="center">
  <a href="https://github.com/WizisCool/oh-my-cpa">
    <picture>
      <source media="(prefers-color-scheme: dark)" srcset="https://raw.githubusercontent.com/WizisCool/oh-my-cpa/main/web/src/assets/brand/omc-wordmark-dark.svg">
      <img src="https://raw.githubusercontent.com/WizisCool/oh-my-cpa/main/web/src/assets/brand/omc-wordmark-light.svg" alt="Oh-My-CPA" height="22">
    </picture>
  </a>
</p>

<p align="center">
  <a href="https://github.com/WizisCool/oh-my-cpa">Oh-My-CPA</a> 与
  <a href="https://github.com/router-for-me/CLIProxyAPI">CLIProxyAPI</a> 的官方插件软件源仓库
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

本项目发布 [Oh-My-CPA](https://github.com/WizisCool/oh-my-cpa) (OMC) 与
[CLIProxyAPI](https://github.com/router-for-me/CLIProxyAPI) (CPA) 生态的 2 个独立软件源列表（以及 1 个聚合总索引）：
- **`official.json`**：官方团队直接维护并校验的核心提供商与集成插件
- **`community.json`**：社区开发者贡献与维护的扩展插件
- **`registry.json`**：包含官方维护与社区贡献的聚合总索引

---

## 软件源订阅地址

您可以根据部署与安全需求选择订阅对应的源：

### 1. 官方维护源 (`official.json`)

仅包含官方团队维护的核心集成插件：
```
https://raw.githubusercontent.com/omc-community/omc-plugins/main/official.json
```

| 标识 (ID) | 名称 | 描述 | 作者 | 仓库地址 |
| --- | --- | --- | --- | --- |
| `omc-zcode-oauth` | **OMC ZCode OAuth** | 支持 Z.ai / BigModel 双入口 OAuth 登录、模型发现、额度路由与配额规范化。 | JUNZE | [omc-community/omc-zcode-oauth](https://github.com/omc-community/omc-zcode-oauth) |
| `omc-qoder-oauth` | **OMC Qoder OAuth** | 支持 Qoder 与 Qoder CN 双入口 OAuth 登录、OpenAI 兼容调用及配额展示。 | JUNZE | [omc-community/omc-qoder-oauth](https://github.com/omc-community/omc-qoder-oauth) |
| `omc-workbuddy-oauth` | **OMC WorkBuddy OAuth** | 支持 WorkBuddy（中国版）与 WorkBuddy AI（国际版）双入口 OAuth 登录、OpenAI 兼容调用及配额展示。 | JUNZE | [omc-community/omc-workbuddy-oauth](https://github.com/omc-community/omc-workbuddy-oauth) |

### 2. 社区贡献源 (`community.json`)

包含社区贡献者提交的扩展与测试工具：
```
https://raw.githubusercontent.com/omc-community/omc-plugins/main/community.json
```

| 标识 (ID) | 名称 | 描述 | 作者 | 仓库地址 |
| --- | --- | --- | --- | --- |
| `omcpa-codex-candy-eval` | **降智测试** | 糖果推理题、模型指纹与鹈鹕绘图测试，支持批量评测与历史记录。 | Yorick-Ryu / haowang02 | [Yorick-Ryu/cpa-plugin-codex-candy-eval](https://github.com/Yorick-Ryu/cpa-plugin-codex-candy-eval) |

### 3. 聚合总源 (`registry.json`)

一键订阅包含所有官方及社区插件：
```
https://raw.githubusercontent.com/omc-community/omc-plugins/main/registry.json
```

---

## 注册与使用方法

### 在 Oh-My-CPA 中使用

Oh-My-CPA 默认集成聚合总源。您也可以在控制台的 **设置 → 插件 → 插件商店** 中单独添加或开关 `official.json` 与 `community.json`。

### 在 CLIProxyAPI (CPA) 中使用

将需要的源地址加入 CPA 的 `config.yaml` 配置文件中：

```yaml
plugins:
  enabled: true
  store-sources:
    - https://raw.githubusercontent.com/omc-community/omc-plugins/main/official.json
    - https://raw.githubusercontent.com/omc-community/omc-plugins/main/community.json
```

---

## 提交社区插件

欢迎为 OMC 社区生态贡献插件！提交流程如下：

1. **Fork** 本仓库。
2. 在 **`community.json`**（官方团队核心扩展则编辑 `official.json`）中添加你的插件条目。
   > **注意**：聚合总源 `registry.json` 由 CI 自动生成（且**官方维护插件永远置顶排列在社区贡献插件上方**），无需且请勿手动编辑 `registry.json`。
3. 确保字段满足 CPA/OMC 校验规范：
   - `id`: 唯一标识符，满足 `^[A-Za-z0-9][A-Za-z0-9._-]{0,127}$`。
   - `name`: 插件展示名称。
   - `description`: 插件功能简述。
   - `author`: 插件作者或组织。
   - `repository`: 严格的 GitHub 仓库完整链接 `https://github.com/{owner}/{repo}`（不支持简写，不可带 `.git` 后缀）。
   - `version`: 语义化版本号，不带前导 `v`。
   - `install`: 指定 `{"type": "github-release"}`。
   - `locales`: 可选的多语言描述映射（如 `zh-CN`, `zh-Hant`, `ms`）。
   - `logo`: 可选的直接 HTTPS 图片地址。
4. 本地运行测试验证（或通过 `go run ./scripts/generate.go` 本地生成预览）：
   ```bash
   go test -v ./...
   ```
5. 提交 **Pull Request**。

---

## 许可证

[MIT](LICENSE)
