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
[CLIProxyAPI](https://github.com/router-for-me/CLIProxyAPI) (CPA) 生态的官方插件软件源索引（`registry.json`）。

收录经过审核与测试的官方提供商扩展、辅助工具及社区贡献插件，支持在 Oh-My-CPA 控制台及 CPA 中一键发现并安装。

---

## 软件源地址

```
https://raw.githubusercontent.com/omc-community/omc-plugins/main/registry.json
```

---

## 注册与使用方法

### 在 Oh-My-CPA 中使用

Oh-My-CPA 默认集成本软件源。可在控制台的 **设置 → 插件 → 插件商店** 中直接查看与管理该源的状态。

### 在 CLIProxyAPI (CPA) 中使用

将该地址加入 CPA 的 `config.yaml` 配置文件中：

```yaml
plugins:
  enabled: true
  store-sources:
    - https://raw.githubusercontent.com/omc-community/omc-plugins/main/registry.json
```

---

## 当前收录插件

| 标识 (ID) | 名称 | 描述 | 作者 | 仓库地址 |
| --- | --- | --- | --- | --- |
| `omc-zcode-oauth` | **OMC ZCode OAuth** | 支持 Z.ai / BigModel 双入口 OAuth 登录、模型发现、额度路由与配额规范化。 | JUNZE | [omc-community/omc-zcode-oauth](https://github.com/omc-community/omc-zcode-oauth) |
| `omc-qoder-oauth` | **OMC Qoder OAuth** | 支持 Qoder 与 Qoder CN 双入口 OAuth 登录、OpenAI 兼容调用及配额展示。 | JUNZE | [omc-community/omc-qoder-oauth](https://github.com/omc-community/omc-qoder-oauth) |
| `omcpa-codex-candy-eval` | **降智测试** | 糖果推理题、模型指纹与鹈鹕绘图测试，支持批量评测与历史记录。 | Yorick-Ryu / haowang02 | [Yorick-Ryu/cpa-plugin-codex-candy-eval](https://github.com/Yorick-Ryu/cpa-plugin-codex-candy-eval) |

---

## 提交社区插件

欢迎为 OMC 社区生态贡献插件！提交流程如下：

1. **Fork** 本仓库。
2. 在 `registry.json` 中添加你的插件条目。
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
4. 本地运行测试验证：
   ```bash
   go test -v ./...
   ```
5. 提交 **Pull Request**。

---

## 许可证

[MIT](LICENSE)
