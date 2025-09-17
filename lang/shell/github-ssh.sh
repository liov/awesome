ssh-keygen -t ed25519 -C "lby.i@qq.com"
cat ~/.ssh/liov.pub

# Ubuntu/Debian
sudo apt-get install gnupg

# CentOS/RHEL
sudo yum install gnupg

# macOS (Homebrew)
brew install gnupg

gpg --full-generate-key
# 快速生成
gpg --generate-key

# 列出公钥
gpg --list-keys

# 列出私钥
gpg --list-secret-keys

# 设置全局签名
git config --global user.signingkey <KEY-ID>
git config --global commit.gpgsign true

# 或针对单个仓库
git config user.signingkey <KEY-ID>
git config commit.gpgsign true
# 查找 Key ID
gpg --list-secret-keys --keyid-format LONG

# 导出到文件
gpg --armor --export <KEY-ID> > public-key.asc

# 直接显示
gpg --armor --export <KEY-ID>
# 导出私钥（备份）
gpg --export-secret-keys --armor <KEY-ID> > private-key.asc
# 导入
gpg --import public-key.asc
gpg --import private-key.asc
# 获取公钥
gpg --armor --export <KEY-ID> | pbcopy  # macOS
gpg --armor --export <KEY-ID> | xclip   # Linux