Перед запуском чекера:
1. запустить checkerService
2. сервису checkerService отправить запрос по gRPC (registerChecker), из ответа взять uuid, который сгенерировал checkerService
3. создать переменную окружения: 
```bash
export UUID="{{uuid из пункта 2}}"
```
4. запустить чекер:
```bash
go run main.go
```