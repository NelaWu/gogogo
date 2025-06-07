
## API 文檔

## 項目結構

├── config.yaml # 配置文件  
├── main.go # 主入口文件  
├── handler/  
│ └── upload.go # 上傳處理器  
├── s3client/  
│ └── client.go # S3 客戶端  
└── README.md # 說明文檔

## 配置說明

配置文件 `config.yaml` 參數說明：

- `endpoint`: MinIO 服務器地址
- `region`: S3 區域設置
- `access_key`: MinIO 訪問密鑰
- `secret_key`: MinIO 密鑰
- `bucket`: 存儲桶名稱

## 依賴項目

- github.com/gorilla/mux: 路由處理
- github.com/aws/aws-sdk-go-v2: AWS SDK
- github.com/spf13/viper: 配置管理

## 開發注意事項

1. 確保 MinIO 服務器正在運行
2. 檢查配置文件中的訪問憑證是否正確
3. 上傳大文件時注意服務器的內存使用情況

## 故障排除

1. 如果遇到連接錯誤，檢查 MinIO 是否正常運行
2. 確認配置文件中的 endpoint 地址是否正確
3. 檢查 access_key 和 secret_key 是否匹配

## 啟動項目
1. `go run main.go`
2. ```
   docker run -d -p 9000:9000 -p 9001:9001 \
    -e "MINIO_ROOT_USER=minioadmin" \
    -e "MINIO_ROOT_PASSWORD=minioadmin" \
    quay.io/minio/minio server /data --console-address ":9001"
    ```
3. 執行API
    -upload  
        `curl -X POST -F "file=@/Users/nela/Documents/kobe.png" http://localhost:8080/upload`
