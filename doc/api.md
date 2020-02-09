#  啟動伺服器

- Golang Version：go1.13.4

- Use Golang Gin：`go run main.go`

- Build and start the server

  ~~~go
  go build
  ./app
  ~~~



# 資料庫 （PostgreSQL)

* **Table**

  1. **auth**：身份認證

     ~~~
     id // [primary_key, serial] int
     account // varchar 學生學號
     password // varchar 學生密碼
     role // varchar 身份
     ~~~

  2. **users**：用戶資料

     ~~~
     id // [foreign_key, serial] int
     name // varchar
     phone // varchar
     ~~~



# Restful API

- `baseURL` = http://localhost:8080/api

  `v1URL` = http://localhost:8080/api/v1

- **Request Format (jwt authorization)**

  ~~~json
  header: {
  	"Authorization": "Bearer " + {$token}
  }
  ~~~

- **Response Format**

  ```json
  resp = {
    "msg": "", // 訊息
    "result": "", // 回傳的資料
  }
  ```

- **Status Code**

  | Status Code             | Description                   |
  | ----------------------- | ----------------------------- |
  | 200 OK                  | 請求成功                      |
  | 400 Bad Request         | (客戶端錯誤) 收到無效語法     |
  | 404 Not Found           | (客戶端錯誤) 找不到請求的資源 |
  | 503 Service Unavailable | (伺服器錯誤) 服務無法使用     |

* **baseAPI**

  | API Method | API URL        | Desc | Req Params                     | Resp Result |
  | ---------- | -------------- | ---- | ------------------------------ | ----------- |
  | POST       | baseURL/signup | 註冊 | account, password, name, phone |             |
  | POST       | baseURL/login  | 登入 | account, password              | token       |

* **usersAPI**

  | API Method | API URL            | Desc         | Req Params | Resp Result |
  | ---------- | ------------------ | ------------ | ---------- | ----------- |
  | GET        | v1URL/users/detail | 取得用戶資料 | id         |             |

  