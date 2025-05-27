# 脚本使用指南

本目录包含两个用于 GitHub 仓库设置和配置检查的脚本：`setup-github.sh` 和 `check-github-config.sh`。

## 在 Windows PowerShell 中运行

在 Windows 系统中，有以下几种方式可以运行这些 Bash 脚本：

### 方式 1：使用 Git Bash（推荐）

1. 安装 [Git for Windows](https://gitforwindows.org/)
2. 右键点击脚本文件，选择 "Git Bash Here"
3. 在打开的 Git Bash 窗口中运行：
```bash
./setup-github.sh [选项]
# 或
./check-github-config.sh
```

### 方式 2：使用 WSL（Windows Subsystem for Linux）

1. 安装并启用 WSL（Windows Subsystem for Linux）
2. 在 PowerShell 中使用以下命令：
```powershell
wsl bash ./scripts/setup-github.sh [选项]
# 或
wsl bash ./scripts/check-github-config.sh
```

### 方式 3：直接在 PowerShell 中运行

如果已安装 Git for Windows，可以在 PowerShell 中直接使用：
```powershell
bash ./scripts/setup-github.sh [选项]
# 或
bash ./scripts/check-github-config.sh
```

注意：确保脚本具有执行权限。如果遇到权限问题，可以在 Git Bash 中运行：
```bash
chmod +x ./scripts/*.sh
```

## setup-github.sh

用于自动化设置 GitHub 仓库、SSH 密钥和相关配置的脚本。

### 前置条件

1. 已安装 GitHub CLI (`gh`)
2. 已登录 GitHub CLI
3. Bash shell 环境

### 基本用法

```bash
./scripts/setup-github.sh [选项]
```

### 可用选项

- `-h, --help`        显示帮助信息
- `-i, --init`        初始化本地仓库
- `-s, --ssh`         生成新的 SSH 密钥
- `-p, --step`        执行特定步骤 (1-5):
  1. 创建 SSH 部署密钥
  2. 配置 SSH 配置文件
  3. 创建 GitHub 仓库
  4. 添加部署密钥
  5. 配置 GitHub Secrets
- `-u, --username`    GitHub 用户名 (默认: ymc-github)
- `-r, --repo`        仓库名称 (默认: go-editjsonns)
- `-e, --email`       邮箱地址 (默认: ymc.github@gmail.com)
- `-d, --docker-user` Docker Hub 用户名
- `-t, --docker-token` Docker Hub 令牌

### GitHub Secrets 配置

配置 GitHub Secrets 有两种方式：

1. 使用 secrets 文件（推荐）：
   - 创建 `secret/gh.token.auto.md` 文件
   - 文件格式示例：
   ```
   DOCKERHUB_USERNAME=your-username
   DOCKERHUB_TOKEN=your-token
   # 可以添加其他 secrets
   ```
   - 运行命令：
   ```bash
   ./scripts/setup-github.sh -p 5
   ```

2. 使用命令行参数：
   ```bash
   ./scripts/setup-github.sh -d your-dockerhub-user -t your-dockerhub-token
   ```

注意：secrets 文件应该添加到 `.gitignore` 中以避免泄露敏感信息。

### 使用示例

1. 基本设置（不包含 SSH）：
```bash
./scripts/setup-github.sh
```

2. 包含 SSH 密钥生成：
```bash
./scripts/setup-github.sh -s
```

3. 执行特定步骤（例如：只创建 SSH 密钥）：
```bash
./scripts/setup-github.sh -p 1
```

4. 指定用户名和仓库：
```bash
./scripts/setup-github.sh -u your-username -r your-repo
```

5. 配置 Docker Hub 凭据：
```bash
./scripts/setup-github.sh -d your-dockerhub-user -t your-dockerhub-token
```

### SSH 密钥命名规则

SSH 密钥文件会按照以下格式命名：
- 格式：`gh_用户名_仓库名`
- 所有连字符 `-` 会被替换为下划线 `_`
- 多个连续的下划线会被压缩为一个
- 示例：用户 "ymc-github" 的仓库 "go-editjsonns" 生成的密钥名为 "gh_ymc_github_go_editjsonns"

## check-github-config.sh

用于检查 GitHub 仓库配置是否正确的脚本。

### 基本用法

```bash
./scripts/check-github-config.sh
```

### 检查项目

脚本会检查以下配置：

1. SSH 密钥
   - 检查密钥文件是否存在
   - 使用与 setup-github.sh 相同的命名规则

2. SSH 配置
   - 检查 ~/.ssh/config 中的配置是否正确

3. GitHub SSH 连接
   - 测试与 GitHub 的 SSH 连接是否正常

4. 远程仓库配置
   - 检查是否正确配置了 origin 远程仓库

5. 分支配置
   - 检查 main 分支是否存在
   - 检查 develop 分支是否存在

6. 必要文件
   - .github/workflows/test.yml
   - .github/workflows/release.yml
   - CONTRIBUTING.md
   - LICENSE
   - README.md

### 输出说明

- ✅ 表示检查项通过
- ❌ 表示检查项失败

如果发现任何配置问题，可以使用 `setup-github.sh` 进行修复。 


```powershell
bash ./scripts/setup-github.sh -p 1
bash ./scripts/setup-github.sh -p 2
bash ./scripts/setup-github.sh -p 3
bash ./scripts/setup-github.sh -p 4
bash ./scripts/setup-github.sh -p 5

sh -c "ls -a ./ | grep \."

git add .gitignore .gitattributes; git commit -m "build(core): init for git"
git add .editorconfig; git commit -m "build(core): init style for text editor"
git add .dockerignore  Dockerfile; git commit -m "build(core): dev in win.wsl.docker.alpine"
git add .github/workflows; git commit -m "build(core): build in github workflow"

git add CONTRIBUTING.md; git commit -m "docs(core): set note for contributing"
git add .github/RELEASE_TEMPLATE.md ; git commit -m "docs(core): set note for release_template"

git add scripts ; git commit -m "build(core): set scripts for opv this repo"
git add scripts/README.md ; git commit -m "docs(core): set note opv.this.repo"


git add LICENSE; git commit -m "docs(core): set license for this repo"

git add pkg/jsonns; git commit -m "build(core): code jsonns in workspace"

git add go.mod; git commit -m "build(core): set this project"
git add CHANGELOG.md; git commit -m "docs(core): set note for changelog"
git add README.md; git commit -m "docs(core): set readme for this repo"

git add scripts ; git commit -m "build(core): put scripts for ssh.config.gh"

```