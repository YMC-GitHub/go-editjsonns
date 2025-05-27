#!/bin/bash

# 帮助信息
show_help() {
    echo "用法: $0 [选项]"
    echo "选项:"
    echo "  -h, --help        显示帮助信息"
    echo "  -i, --init        初始化本地仓库"
    echo "  -s, --ssh         生成新的 SSH 密钥"
    echo "  -p, --step        执行特定步骤 (1-5):"
    echo "                    1: 创建 SSH 部署密钥"
    echo "                    2: 配置 SSH 配置文件"
    echo "                    3: 创建 GitHub 仓库"
    echo "                    4: 添加部署密钥"
    echo "                    5: 配置 GitHub Secrets"
    echo "  -u, --username    GitHub 用户名 (默认: ymc-github)"
    echo "  -r, --repo        仓库名称 (默认: go-editjsonns)"
    echo "  -e, --email       邮箱地址 (默认: ymc.github@gmail.com)"
    echo "  -d, --docker-user Docker Hub 用户名"
    echo "  -t, --docker-token Docker Hub 令牌"
    echo "示例:"
    echo "  $0                     # 基本设置"
    echo "  $0 -s                  # 包含 SSH 密钥生成"
    echo "  $0 -p 1               # 仅创建 SSH 部署密钥"
    echo "  $0 -p 5               # 仅配置 GitHub Secrets"
    echo "  $0 -u user -r repo     # 指定用户名和仓库"
}

# 默认配置变量
GITHUB_USERNAME="ymc-github"
REPO_NAME="go-editjsonns"
REPO_DESCRIPTION="A Go library for parsing and manipulating JSON namespace expressions"
# 格式化 SSH 密钥名称
format_key_name() {
    local name="gh_${1}_${2}"
    # 将 - 替换为 _
    name=$(echo "$name" | tr '-' '_')
    # 将多个连续的 _ 替换为单个 _
    name=$(echo "$name" | tr -s '_')
    echo "$name"
}
SSH_KEY_NAME=$(format_key_name "$GITHUB_USERNAME" "$REPO_NAME")
EMAIL="ymc.github@gmail.com"
INIT_REPO=false
DOCKER_USERNAME=""
DOCKER_TOKEN=""
GENERATE_SSH=false
SPECIFIC_STEP=0

# 检查 gh 是否已安装
if ! command -v gh &> /dev/null; then
    echo "错误: GitHub CLI (gh) 未安装"
    echo "请访问 https://cli.github.com/ 安装 gh 命令行工具"
    exit 1
fi

# 检查 gh 是否已登录
if ! gh auth status &> /dev/null; then
    echo "请先登录 GitHub CLI:"
    gh auth login
fi

# 解析命令行参数
while [[ $# -gt 0 ]]; do
    case $1 in
        -h|--help)
            show_help
            exit 0
            ;;
        -i|--init)
            INIT_REPO=true
            shift
            ;;
        -s|--ssh)
            GENERATE_SSH=true
            shift
            ;;
        -p|--step)
            SPECIFIC_STEP="$2"
            shift 2
            ;;
        -u|--username)
            GITHUB_USERNAME="$2"
            SSH_KEY_NAME=$(format_key_name "$2" "$REPO_NAME")
            shift 2
            ;;
        -r|--repo)
            REPO_NAME="$2"
            SSH_KEY_NAME=$(format_key_name "$GITHUB_USERNAME" "$2")
            shift 2
            ;;
        -e|--email)
            EMAIL="$2"
            shift 2
            ;;
        -d|--docker-user)
            DOCKER_USERNAME="$2"
            shift 2
            ;;
        -t|--docker-token)
            DOCKER_TOKEN="$2"
            shift 2
            ;;
        *)
            echo "未知选项: $1"
            show_help
            exit 1
            ;;
    esac
done

# 创建 SSH 部署密钥
create_ssh_key() {
    echo "=== 1. 生成 SSH 部署密钥 ==="
    if [ -f ~/.ssh/${SSH_KEY_NAME} ]; then
        echo "SSH 密钥 ${SSH_KEY_NAME} 已存在，将使用现有密钥"
    else
        echo "创建新的 SSH 密钥..."
        ssh-keygen -t ed25519 -C "$EMAIL" -f ~/.ssh/${SSH_KEY_NAME} -N ""
    fi
}

# 配置 SSH 配置文件
configure_ssh() {
    echo "=== 2. 配置 SSH 配置文件 ==="
        cat > ~/.ssh/config << EOF
Host github.com
    HostName github.com
    User git
    IdentityFile ~/.ssh/${SSH_KEY_NAME}
EOF
    cat ~/.ssh/config
}

# 创建 GitHub 仓库
create_github_repo() {
    echo "=== 3. 创建 GitHub 仓库 ==="
    # git init
    gh repo create "${GITHUB_USERNAME}/${REPO_NAME}" --public --description "${REPO_DESCRIPTION}"
}

# 添加部署密钥
add_deploy_key() {
    echo "=== 4. 添加部署密钥 ==="
    # git init
    # git remote add origin git@github.com:${GITHUB_USERNAME}/${REPO_NAME}.git
    gh repo deploy-key add ~/.ssh/${SSH_KEY_NAME}.pub --title "${SSH_KEY_NAME}" --allow-write
}

# 配置 GitHub Secrets
configure_secrets() {
    echo "=== 5. 配置 GitHub Secrets ==="
    local repo="${GITHUB_USERNAME}/${REPO_NAME}"
    if [ -f "secret/gh.token.auto.md" ]; then
        echo "从文件加载 GitHub Secrets..."
        gh secret set --repo "$repo" -f secret/gh.token.auto.md
    else
        if [ -n "$DOCKER_USERNAME" ] && [ -n "$DOCKER_TOKEN" ]; then
            echo "使用命令行参数配置 GitHub Secrets..."
            gh secret set --repo "$repo" DOCKERHUB_USERNAME -b"${DOCKER_USERNAME}"
            gh secret set --repo "$repo" DOCKERHUB_TOKEN -b"${DOCKER_TOKEN}"
        else
            echo "错误: 未找到 secret/gh.token.auto.md 文件，且未提供 Docker Hub 凭据"
            echo "请提供以下任一选项："
            echo "1. 创建 secret/gh.token.auto.md 文件"
            echo "2. 使用 -d 和 -t 选项提供 Docker Hub 凭据"
            exit 1
        fi
    fi
}

# 根据特定步骤执行
if [ "$SPECIFIC_STEP" != "0" ]; then
    case $SPECIFIC_STEP in
        1)
            create_ssh_key
            ;;
        2)
            configure_ssh
            ;;
        3)
            create_github_repo
            ;;
        4)
            add_deploy_key
            ;;
        5)
            configure_secrets
            ;;
        *)
            echo "错误: 无效的步骤编号 ${SPECIFIC_STEP}"
            show_help
            exit 1
            ;;
    esac
    exit 0
fi

# 正常流程执行
if [ "$GENERATE_SSH" = true ]; then
    create_ssh_key
    configure_ssh
fi

create_github_repo

if [ "$GENERATE_SSH" = true ]; then
    add_deploy_key
fi

if [ -n "$DOCKER_USERNAME" ] && [ -n "$DOCKER_TOKEN" ]; then
    configure_secrets
fi

if [ "$INIT_REPO" = true ]; then
    echo "=== 6. 初始化本地仓库 ==="
    git init
    git add .
    git commit -m "feat: initial commit"

    echo "=== 7. 创建并推送分支 ==="
    git checkout -b main
    git push -u origin main

    git checkout -b develop
    git push -u origin develop
fi


echo "仓库地址：https://github.com/${GITHUB_USERNAME}/${REPO_NAME}" 