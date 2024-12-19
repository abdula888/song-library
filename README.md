# Онлайн библиотека песен
**Тестовое задание для Effective Mobile**  
## Стек технологий
- **Backend**: Golang
- **База данных**: PostgreSQL
- **Контейнеризация**: Docker Compose  

## REST методы  
1. Получение данных библиотеки с фильтрацией по всем полям и пагинацией.
2. Получение текста песни с пагинацией по куплетам.
3. Удаление песни.
4. Изменение данных песни.
5. Добавление новой песни в формате JSON
{
 "group": "Muse",
 "song": "Supermassive Black Hole"
}.

При добавлении песни делается запрос во внешние API (AUDD_API и LASTFM_API) для обогащения информации (дата релиза, текст песни и ссылка на youtube).  
## Логирование
Проект поддерживает уровни логирования `debug` и `info`. Уровень задаётся в `.env` файле.

Пример:
```env
LOG_LEVEL=info
```
## Запуск
Соберите контейнеры:
```bash
docker compose build
```
Запустите приложение:
```bash
docker compose up
```
Приложение будет доступно по адресу:
```
http://localhost:8080/songs
```
