# pdm
## zsh
curl -sSL https://pdm-project.org/install-pdm.py | python - -p /d/SDK/pdm
setx -M PATH "$PATH:/d/SDK/pdm/bin" 不要在zsh中执行,因为有额外的变量
## powershell
irm https://pdm-project.org/install-pdm.py | py - -p D:/SDK/pdm

`setx /M PATH "$env:PATH;D:\SDK\pdm\bin" //powershell 无法识别$PATH 和 %PATH%`
## cmd
`setx /M PATH "%PATH%;D:\SDK\pdm\bin"`

### 最重要的一步
然后可以放弃了去下载windows python了 https://www.python.org/downloads/windows/
curl -sSL https://pdm-project.org/install-pdm.py | python - -p /d/sdk/pdm

pdm config install.cache on
pdm config cache_dir /d/sdk/pdm/Cache
pdm config global_project.path /d/sdk/pdm/global-project
pdm config log_dir /d/sdk/pdm/Logs
pdm config python.install_root /d/sdk/pdm/python
pdm config venv.location /d/sdk/pdm/venvs

# uv
curl -LsSf https://astral.sh/uv/install.sh | sh

uv sync