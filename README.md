# LFS

Large File Storage (LFS)

## 開發動機

因為discord最多只能上傳8MB

所以透過此方法來拆分與合併

> ⚠ 不鼓勵運用此方法來把discord當作儲藏室，該方法應僅運用在傳暫存檔，傳完也請把檔案刪除，避免增加不必要的負擔。

## build

```
git clone -o https://github.com/CarsonSlovoka/lfs.git
cd lfs/src
go build -ldflags "-s -w" -o lfs.exe
```

> 🧙 如果您不想要自己build，可以於release頁面下載，現成的檔案
>
> https://github.com/CarsonSlovoka/lfs/releases/download/v0.0.0/lfs-windows-amd64-v0.0.0.zip

## USAGE

- for split

```yaml
lfs.exe -action=split -src="xxx.exe"
lfs.exe -action=split -src="xxx.exe" -chunkSize=5000000   # 5MB
```

完成之後會於當前的工作路徑下[生成出output的資料夾](https://github.com/CarsonSlovoka/lfs/blob/7bb5ab26630a88695d6700f72e72374c11ec2505/src/main.go#L27)，請在裡面察看結果

for merge

```yaml
lfs.exe -action=merge -src="C://xxx/dir"
lfs.exe -action=merge -src="C://xxx/dir" -maxIdx=10 # 最多可以處理到文件編號 10.lfs
lfs.exe -action=merge -src="./output/" -maxIdx=10
```

> [🚀lfs實用的bat參考](tool)
