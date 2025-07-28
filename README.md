```shell
docker compose -f docker-compose.Checker-Orchestrator.yaml up --build
```
в логах от чекера будет uuid - его надо самостоятельно отправить сервису checkerSystem по gRPC хендлеру CreateVulnService 
```json
{
"ip": "127.0.0.1",
"service_id": "{uuid из лога чекера}",
"web_port": 4040
} 
```
