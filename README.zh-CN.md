# OMC 社区插件源 (OMC Community Plugin Registry)

Oh-My-CPA (OMC) 与 CLIProxyAPI (CPA) 生态的官方插件软件源仓库。

本项目维护可供 Oh-My-CPA 与 CLIProxyAPI 注册使用的插件软件源索引，方便用户在 OMC 控制台商店或 CPA 中一键发现并安装扩展。

## 软件源地址

```
https://raw.githubusercontent.com/omc-community/omc-plugins/main/registry.json
```

## 注册与使用方法

### 在 Oh-My-CPA 中使用

Oh-My-CPA 默认集成此软件源。可在控制台的 **设置 → 插件 → 插件商店** 中查看与管理该源的状态。

### 在 CLIProxyAPI (CPA) 中使用

将该地址加入 CPA 的 `config.yaml` 配置文件中：

```yaml
plugins:
  enabled: true
  store-sources:
    - https://raw.githubusercontent.com/omc-community/omc-plugins/main/registry.json
```

## 当前收录插件

| 标识 (ID) | 名称 | 描述 | 作者 | 仓库地址 |
| --- | --- | --- | --- | --- |
| `omc-zcode-oauth` | OMC ZCode OAuth | 支持 Z.ai / BigModel 双入口 OAuth 登录、模型发现、额度路由与配额规范化。 | JUNZE | [omc-community/omc-zcode-oauth](https://github.com/omc-community/omc-zcode-oauth) |
| `omc-qoder-oauth` | OMC Qoder OAuth | 支持 Qoder 与 Qoder CN 双入口 OAuth 登录、OpenAI 兼容调用及配额展示。 | JUNZE | [omc-community/omc-qoder-oauth](https://github.com/omc-community/omc-qoder-oauth) |
| `omcpa-codex-candy-eval` | 降智测试 | 糖果推理题、模型指纹与鹈鹕绘图测试，支持批量评测与历史记录。 | Yorick-Ryu / haowang02 | [Yorick-Ryu/cpa-plugin-codex-candy-eval](https://github.com/Yorick-Ryu/cpa-plugin-codex-candy-eval) |

## 提交社区插件

欢迎为 OMC 社区生态贡献插件！提交流程如下：

1. Fork 本仓库。
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
4. 本地运行 `go test ./...` 确保校验通过。
5. 提交 Pull Request。

## 许可证

[MIT](LICENSE)
