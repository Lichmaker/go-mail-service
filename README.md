# go-mail-service
golang 编写的简易单邮件发送服务，使用阿里云邮件推送服务

## 配置

### 修改 `.env` 文件
```
cp .env.example .env
vim .env
```

`.env` 示例
```
ALIYUN_ACCESS_KEY_ID=
ALIYUN_ACCESS_KEY_SECRET=
ALIYUN_REGION_ID=
ALIYUN_EMAIL_ACCOUNT_NAME=
```
说明：
- access_key_id 和 access_key_secret 都是通过阿里云的RAM管理获得。 详细可阅读文档： https://help.aliyun.com/document_detail/53045.html
- region_id 是与API服务地址有关。详细可阅读文档：https://help.aliyun.com/document_detail/96856.html
- ALIYUN_EMAIL_ACCOUNT_NAME 是阿里云中“邮件推送控制台”中的“发信地址”

## 启动
直接启动
```
go run main.go
```

使用docker启动
```
docker build -t lichmaker/go-mail-service .
docker run --rm -p 3001:3001 -d lichmaker/go-mail-service --name="go-mail-service"
```

## 调用

```
curl --location --request POST 'localhost:3001' \
--form 'email="lich.wu2014@gmail.com"' \
--form 'title="测试发送邮件-标题"' \
--form 'body="测试发送邮件-内容"'
```
