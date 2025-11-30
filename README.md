Всем привет! Мы команда "Здесь могла быть ваша реклама". Мы участвуем в хакатоне Hack&Change 2025 и работаем над кейсом по созданию прототипа цифровой образовательной платформы от ПСБ.
Наша команда состоит из пяти участников:
- Игаев Даниил - девопс
- Постников Даниил - бэкенд-разработчик
- Малчевский Максим - бэкенд-разработчик
- Жаров Степан - фронтенд-разработчик
- Ермолина Софья - UI/UX дизайнер

## Демонстрация решения
https://drive.google.com/drive/folders/1GAXZ9o8U_bKTVxRnyhqlWOW5k8tdn26I?usp=sharing

## Структура проекта
### Фронтенд
├── client/
│ ├── public/
│ ├── src/
│ │ ├── api/
│ │ ├── assets/
│ │ ├── pages/
│ │ ├── router/
│ │ ├── stores/
│ │ ├── utils/
│ │ ├── widgets/
│ │ ├── App.vue
│ │ └── main.ts
│ ├── .gitignore
│ ├── default.conf
│ ├── Dockerfile
│ ├── env.d.ts
│ ├── index.html
│ ├── package.json
│ ├── package-lock.json
│ ├── tsconfig.app.json
│ ├── tsconfig.json
│ ├── tsconfig.node.json
│ └── vite.config.ts

### Бэкенд
├── server/
│ ├── cmd/
│ ├── docs/
│ ├── internal/
│ │ ├── api/
│ │ ├── config/
│ │ ├── entities/
│ │ ├── logger/
│ │ ├── mock_data/
│ │ ├── repository/
│ │ └── util/
│ ├── Dockerfile
│ ├── go.mod
│ ├── Makefile
│ ├── tmp/
│ ├── .env
│ ├── env.example
│ ├── .gitignore
│ ├── docker-compose.yml

## Запуск
Для запуска решения локально нужно написать эти команды:
```bash
git clone https://github.com/gnomDevelopers/synapse
```

```bash
docker compose up -d
```