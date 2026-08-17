# netviewer
Утилита для просмотра открытых портов. Сделано для практики работы с сетями и конкурентности.

## Пример вывода:
```bash
netviewer % ./bin/app scan --host localhost --from 1 -- to 1024
02:59:55 INF Scanning localhost ports 1-1024
02:59:55 INF OPEN localhost:88
02:59:55 INF OPEN localhost:445
```