# LFS

Large File Storage (LFS)

## 開發動機

因為discord最多只能上傳8MB

所以透過此方法來拆分與合併

> ❗注意❗ 不鼓勵運用此方法來把discord當作儲藏室，該方法應僅運用在傳暫存檔，傳完也請把檔案刪除，避免增加不必要的負擔。

## build

```
go build -ldflags "-s -w"
```

## USAGE

- for split

```yaml
go run main.go -action=split -src="xxx.exe"
go run main.go -action=split -src="xxx.exe" -chunkSize=5000000   # 5MB
```

for merge

```yaml
go run main.go -action=merge -src="C://xxx/dir"
go run main.go -action=merge -src="C://xxx/dir" -maxIdx=10 # 最多可以處理到文件編號 10.lfs
go run main.go -action=merge -src="./output/" -maxIdx=10
```
