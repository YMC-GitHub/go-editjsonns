#!/bin/bash

# 配置变量
GITHUB_USERNAME="ymc-github"
REPO_NAME="go-editjsonns"

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

# 检查 SSH 密钥
echo "=== 检查 SSH 密钥 ==="
if [ -f ~/.ssh/${SSH_KEY_NAME} ]; then
    echo "✅ SSH 密钥存在"
else
    echo "❌ SSH 密钥不存在"
fi

# 检查 SSH 配置
echo -e "\n=== 检查 SSH 配置 ==="
if grep -q "github.com-${REPO_NAME}" ~/.ssh/config; then
    echo "✅ SSH 配置已添加"
else
    echo "❌ SSH 配置未添加"
fi

# 测试 SSH 连接
echo -e "\n=== 测试 GitHub SSH 连接 ==="
ssh -T git@github.com 2>&1 | grep -q "successfully authenticated"
if [ $? -eq 0 ]; then
    echo "✅ SSH 连接成功"
else
    echo "❌ SSH 连接失败"
fi

# 检查远程仓库配置
echo -e "\n=== 检查远程仓库配置 ==="
if git remote -v | grep -q "origin.*github.com.*${REPO_NAME}"; then
    echo "✅ 远程仓库已配置"
else
    echo "❌ 远程仓库未配置"
fi

# 检查分支
echo -e "\n=== 检查分支配置 ==="
if git branch | grep -q "main"; then
    echo "✅ main 分支存在"
else
    echo "❌ main 分支不存在"
fi

if git branch | grep -q "develop"; then
    echo "✅ develop 分支存在"
else
    echo "❌ develop 分支不存在"
fi

# 检查必要文件
echo -e "\n=== 检查必要文件 ==="
files=(".github/workflows/test.yml" ".github/workflows/release.yml" "CONTRIBUTING.md" "LICENSE" "README.md")
for file in "${files[@]}"; do
    if [ -f "$file" ]; then
        echo "✅ $file 存在"
    else
        echo "❌ $file 不存在"
    fi
done

echo -e "\n如果发现任何问题，请参考 scripts/setup-github.sh 中的说明进行修复。" 