## GO 

### Project create 

```bash

PS D:\sdev\222dis\gauth> go mod init github.com/222dis/gauth
go: creating new go.mod: module github.com/222dis/gauth
PS D:\sdev\222dis\gauth>
```
*github.com/222dis/gauth debe ser único


```bash
Tools environment: GOPATH=C:\Users\Lab_Redes\go
Installing 7 tools at C:\Users\Lab_Redes\go\bin in module mode.
  gotests
  gomodifytags
  impl
  goplay
  dlv
  staticcheck
  gopls


```

### Runing local  


```bash
PS D:\sdev\222dis\gauth>go install github.com/cosmtrek/air@latest

PS D:\sdev\222dis\gauth>air

```

http://localhost:3000/users

mongo-express 
http://localhost:7999/

### Runing docker  

```bash
PS D:\sdev\222dis\gauth>docker-compose up --build -d

PS D:\sdev\222dis\gauth> docker exec -it gauth_app bash
root@a5b44fbef5b1:/go/src/app# air

```

ToDo ...




