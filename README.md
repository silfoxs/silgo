# silgo
``` bash
      ___                       ___       ___           ___     
     /\  \          ___        /\__\     /\  \         /\  \    
    /::\  \        /\  \      /:/  /    /::\  \       /::\  \   
   /:/\ \  \       \:\  \    /:/  /    /:/\:\  \     /:/\:\  \  
  _\:\~\ \  \      /::\__\  /:/  /    /:/  \:\  \   /:/  \:\  \ 
 /\ \:\ \ \__\  __/:/\/__/ /:/__/    /:/__/_\:\__\ /:/__/ \:\__\
 \:\ \:\ \/__/ /\/:/  /    \:\  \    \:\  /\ \/__/ \:\  \ /:/  /
  \:\ \:\__\   \::/__/      \:\  \    \:\ \:\__\    \:\  /:/  / 
   \:\/:/  /    \:\__\       \:\  \    \:\/:/  /     \:\/:/  /  
    \::/  /      \/__/        \:\__\    \::/  /       \::/  /   
     \/__/                     \/__/     \/__/         \/__/    
```

## 环境要求
golang >= 1.19

## 入门示例
``` bash
> git clone git@github.com:silfoxs/silgo.git
> cd silgo
> go mod tidy
> go go run main.go silgo-api

Using config file: ./configs/dev.yaml
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /ping                     --> github.com/silfoxs/silgo/internal/app.Run.func1 (3 handlers)
```

## 参数示例
### config 指定启动配置文件默认./configs/dev.yaml
--config=./configs/dev.yaml
