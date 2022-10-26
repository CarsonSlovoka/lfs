@echo off
echo src: %1
pause
lfs.exe -action=merge -src=%1 -maxIdx=10
pause