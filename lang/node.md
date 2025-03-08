# nvm
curl -o- https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.1/install.sh | bash

## windows
https://github.com/coreybutler/nvm-windows/releases
NVM_HOME=D:\sdk\nvm
NVM_SYMLINK=D:\sdk\nodejs
nvm install latest 64
nvm use latest
# pnpm
corepack enable pnpm

pnpm config set registry https://registry.npmmirror.com/

pnpm config set registry https://registry.npmjs.org

pnpm config set global-bin-dir D:\sdk\pnpm\bin
pnpm config set cache-dir D:\sdk\pnpm\cache
pnpm config set state-dir D:\sdk\pnpm\state
pnpm config set global-dir D:\sdk\pnpm\bin\global