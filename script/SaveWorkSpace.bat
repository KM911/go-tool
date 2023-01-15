@REM for /f "delims=," %%a in ('
@REM     tasklist /fi "pid eq %1" /nh /fo:csv
@REM ') do echo %%~a

@REM 创建一个data.txt文件 将其写入
@echo off
setlocal enabledelayedexpansion
set "data="

@REM 我们的参数是一个或者多个进程ID 例如 test.bat 9116 9604
for %%a in (%*) do (
    @REM 通过tasklist命令获取进程信息
    for /f "delims=," %%b in ('tasklist /fi "pid eq %%a" /nh /fo:csv') do (
        @REM 将进程信息写入data.txt文件
        set "data=!data!%%b,"
    )
)

@REM 将data.txt文件写入
echo !data! > script\data.txt

