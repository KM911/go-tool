' 获取一个路径参数 然后cd 到该路径

' 1. 获取参数
set objArgs = WScript.Arguments
' 2. 输出参数
Set shell = CreateObject("WScript.Shell")
shell.Sendkeys "cd " & objArgs(0)
shell.Sendkeys "{ENTER}"

