# canvas-study-backend
(canvas-study)[https://github.com/b1018043/canvas-study]のバックエンド
 
## 利用手順
## docker
```bash
docker build -t imagename .
docker run -d -p 8080:8080 imagename
```
## dockerを利用しない場合
初めに以下のコマンドを実行
```bash
go mod download
```
### Makefile
```bash
make
./main
```
### go build
```bash
go build -o main main.go
./main
```
## 利用技術
- github.com/gin-gonic/gin
- github.com/gorilla/websocket
- gopkg.in/olahol/melody.v1
