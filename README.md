# Go Log Analyzer

Что умеет анализатор логов:
- Считывать и парсить Nginx/Apache форматов
- Выводить Топ-N IP-адресов
- Выводить статистику по статусам

## Установка и использование
```bash
git clone https://github.com/RomanKovalev007/go-log-analyzer
cd go-log-analyzer
go build -o logs/analyzer ./cmd/analyzer
./analyzer -f access.log -n 10
```
## Используемые флаги
-f путь к лог-файлу (по умолчанию: logs/access.log)\n
-n количество топN IP (по умолчанию: 5)


## Пример вывода
HTTP Status Statistics:\n
404: 2 (client error)\n
200: 3 (success)\n
503: 1 (server error)\n
302: 1 (redirect)\n

Top 10 IPs:\n
127.0.0.1: 4 requests\n
192.168.1.1: 2 requests\n
2001:db8::1: 1 requests\n
Memory Usage: Alloc = 0.18 MB   TotalAlloc = 0 MiB      Sys = 6 MiB     NumGC = 0
