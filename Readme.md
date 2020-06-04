# Navicat 連線密碼查尋
將連線記錄中的連線密碼解碼

# 適用版本
Navicat 12

# 使用方式
先啟動Navicat，並將連線設定匯出成*.ncx檔案（Navicat預設檔案），
執行命令
```
NaviPassRead -f [檔案路徑]
```
檔案路徑預設為 ./navi.ncx

# 輸出範例
```
[
    {
        "ConnType":"MYSQL",
        "ConnectionName":"music",
        "Host":"localhost",
        "Password":"55688",
        "Port":"3306",
        "SSH_Host":"123.45.67.89",
        "SSH_UserName":"bob",
        "SSH_Password":"bobbobobo",
        "ServiceProvider":"Default",
        "UserName":"bob"
    },
    {
        "ConnType":"MYSQL",
        "ConnectionName":"GCP",
        "Host":"88.77.66.55",
        "Password":"password123",
        "Port":"3306",
        "SSH_Host":"",
        "SSH_UserName":"",
        "SSH_Password":"",
        "ServiceProvider":"GoogleCloud",
        "UserName":"alice"
    }
]

```
