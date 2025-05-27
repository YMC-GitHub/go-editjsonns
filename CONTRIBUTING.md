# Contributing Guide

## Development Workflow

### 1. Branch Management

- `main`: 主分支，保持稳定
- `develop`: 开发分支，用于集成功能
- `feature/*`: 功能分支，如 `feature/new-parser`
- `fix/*`: 修复分支，如 `fix/array-index-bug`
- `release/*`: 发布分支，如 `release/v1.0.0`

### 2. Version Management

We follow [Semantic Versioning](https://semver.org/):

- MAJOR version (x.0.0) - 不兼容的 API 更改
- MINOR version (0.x.0) - 向后兼容的功能添加
- PATCH version (0.0.x) - 向后兼容的错误修复

### 3. Development Process

1. 创建新分支：
   ```bash
   # 功能分支
   git checkout -b feature/new-feature develop
   
   # 修复分支
   git checkout -b fix/bug-description develop
   ```

2. 提交更改：
   ```bash
   # 使用规范的提交消息
   git commit -m "feat: add new feature"
   git commit -m "fix: resolve array index issue"
   ```

3. 更新 CHANGELOG.md：
   - 在 [Unreleased] 部分添加更改记录
   - 遵循 "Added", "Changed", "Deprecated", "Removed", "Fixed", "Security" 分类

4. 合并到开发分支：
   ```bash
   git checkout develop
   git merge --no-ff feature/new-feature
   ```

### 4. Release Process

1. 从 develop 创建发布分支：
   ```bash
   git checkout -b release/v1.0.0 develop
   ```

2. 更新版本号：
   - 更新 CHANGELOG.md
   - 确保所有测试通过

3. 合并到主分支并标记：
   ```bash
   git checkout main
   git merge --no-ff release/v1.0.0
   git tag -a v1.0.0 -m "Release version 1.0.0"
   git push origin v1.0.0
   ```

4. 合并回 develop：
   ```bash
   git checkout develop
   git merge --no-ff release/v1.0.0
   ```

### 5. Commit Message Guidelines

遵循 [Conventional Commits](https://www.conventionalcommits.org/)：

```
<type>(<scope>): <description>

[optional body]

[optional footer]
```

类型：
- feat: 新功能
- fix: 错误修复
- docs: 文档更改
- style: 格式调整
- refactor: 代码重构
- test: 测试相关
- chore: 构建过程或辅助工具变动

示例：
```bash
feat(parser): add support for nested arrays
fix(validator): handle empty string input
docs(readme): update installation instructions
```

### 6. Code Review Process

1. 创建 Pull Request 时：
   - 提供清晰的描述
   - 包含测试用例
   - 更新相关文档
   - 确保 CI 测试通过

2. Review 检查项：
   - 代码质量和风格
   - 测试覆盖率
   - 文档完整性
   - 性能影响

### 7. Release Checklist

- [ ] 所有测试通过
- [ ] CHANGELOG.md 已更新
- [ ] 文档已更新
- [ ] API 兼容性已检查
- [ ] 性能测试已完成
- [ ] 代码审查已完成 