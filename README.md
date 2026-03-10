# CoalMine

## Список технологий

### Frontend

| Технология                       | Назначение                     |
|----------------------------------|--------------------------------|
| [Vue 3](https://vuejs.org/)      | Фреймворк                      |
| [Vite](https://vite.dev/)        | Сборщик и dev-сервер           |
| [Axios](https://axios-http.com/) | HTTP-клиент для запросов к API |

### Backend

Go 1.25 + [Gorilla Mux](https://github.com/gorilla/mux)

---

## Памятка по запуску

Требования: установленные **Docker** и **Docker Compose**.

```bash
cd app
docker compose up --build
```

- Фронтенд: [http://localhost:5173](http://localhost:5173)
- Бэкенд API: [http://localhost:9091](http://localhost:9091)

Остановить:

```bash
docker compose down
```