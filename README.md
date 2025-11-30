# Hack&Change 2025 - Образовательная платформа

~~всем привет~~ Мы команда "Здесь могла быть ваша реклама". Мы участвуем в хакатоне Hack&Change 2025 и работаем над кейсом по созданию прототипа цифровой образовательной платформы от ПСБ.

## Наша оманда

- Игаев Даниил — Baskend Developer
- Постников Даниил — Backend Developer
- Малчевский Максим — Backend Developer
- Жаров Степан — Frontend Developer
- Ермолина Софья — UI/UX Designer

## Демонстрация

[Видео демонстрация работы проекта](https://drive.google.com/drive/folders/1GAXZ9o8U_bKTVxRnyhqlWOW5k8tdn26I?usp=sharing)

---

## Структура проекта

### Фронтенд (Vue 3 + TypeScript)

```
client/
├── public/                # Статические файлы
├── src/                   # Исходный код
│   ├── api/              # API интеграция и конфигурация
│   ├── assets/           # Изображения и стили
│   ├── pages/            # Компоненты страниц
│   ├── widgets/          # Переиспользуемые UI компоненты
│   ├── router/           # Маршруты приложения
│   ├── stores/           # Управление состоянием (Pinia)
│   └── utils/            # Вспомогательные функции
└── Конфигурация          # package.json, vite.config.ts, tsconfig.json
```

### Бэкенд (Go)

```
server/
├── cmd/                   # Точка входа приложения
├── internal/              # Внутренний код
│   ├── api/              # HTTP маршруты и обработчики
│   ├── config/           # Конфигурация и переменные окружения
│   ├── entities/         # Структуры данных и модели
│   ├── logger/           # Логирование
│   ├── mock_data/        # Тестовые данные
│   └── repository/       # Работа с БД
├── docs/                 # Документация API
└── Конфигурация          # go.mod, Dockerfile, Makefile
```

---

## Запуск

### Требования

- Docker & Docker Compose
- Git

### Локальный запуск

```bash
# Клонируем репозиторий
git clone https://github.com/gnomDevelopers/synapse.git
cd synapse

# Копируем конфигурацию
cp .env.example .env

# Запускаем контейнеры
docker compose up -d
```

Приложение доступно по адресу `http://localhost:10800`

---

## Дополнительно

- [Frontend README](./client/README.md) — документация фронтенда
- API документация находится в папке `server/docs`
