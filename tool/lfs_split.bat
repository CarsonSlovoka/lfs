@echo off
echo src: %1
pause
:: 100KB 100000
:: 7M    7000000
lfs.exe -action=split -src=%1 -chunkSize=100000
