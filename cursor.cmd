@echo off
chcp 65001 > nul
setlocal enabledelayedexpansion

REM 检查是否以管理员权限运行
net session >nul 2>&1
if %errorlevel% neq 0 (
    echo 正在请求管理员权限...
    powershell -Command "Start-Process '%~f0' -Verb RunAs"
    exit /b
)

REM 管理员权限下执行 PowerShell 命令
echo 正在执行重置cursor命令...
powershell -NoProfile -ExecutionPolicy Bypass -Command "irm https://aizaozao.com/accelerate.php/https://raw.githubusercontent.com/yuaotian/go-cursor-help/refs/heads/master/scripts/run/cursor_win_id_modifier.ps1 | iex"

