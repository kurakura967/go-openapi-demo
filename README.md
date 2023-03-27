# go-openapi-demo
[OpenAPI Generator](https://github.com/deepmap/oapi-codegen)を用いてGo言語のコード生成を試してみるレポジトリ

## 動作確認

```bash
# コードの自動生成
make codegen
# APIサーバーの起動
make run

# 正常系
$ curl -s localhost:8080/v1/user/1 | jq .
{
  "age": 10,
  "gender": "1",
  "id": 1,
  "name": "taro"
}

# invalid_parameter
$ curl -s localhost:8080/v1/user/hoge | jq .
{
  "error": {
    "field": {
      "field_name": "userId",
      "message": "Invalid format for parameter userId: error binding string parameter: strconv.ParseInt: parsing \"hoge\": invalid syntax"
    },
    "message": "Request parameters are invalid",
    "type": "invalid_parameter"
  }
}
```
