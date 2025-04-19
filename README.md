# NTO

## Задачи

- [x] Migrator
- [x] Crud service generator (based on gorm gen)
- [x] Easy init without git repository (like create-nto-app)
- [x] Excel export
- [x] Models validation logic (callbacks and other stuff)
- [x] Models [linter](https:/github.com/kuzgoga/gormlint)
- [x] Auto-generated frontend

## Инициализация
Для того чтобы использовать этот шаблон, выполните команды:
```shell
curl -O https://proxy.golang.org/github.com/kuzgoga/nto-boilerplate/@v/v1.1.6.zip
```
Для Windows:
```shell
Invoke-WebRequest -Uri https://proxy.golang.org/github.com/kuzgoga/nto-boilerplate/@v/v1.1.6.zip -OutFile v1.1.6.zip
```

## Установка

Для разработки необходимы следующие инструменты:

- [Node.js](https://nodejs.org/en)
- [Golang](https://go.dev/dl/)
- [VSCode](https://code.visualstudio.com/) или [Goland](https://cloud.gogacoder.ru/s/pZJzpeNSprmWDB8)
- ~~[UPX](https://github.com/upx/upx/releases/latest) - для сжатия.~~
- [Git](https://git-scm.com/)
- [Obsidian](https://obsidian.md/)
- Wails3
- MinGw-64 (необязательно): https://jmeubank.github.io/tdm-gcc/

## Разработка

### Hot Reload

Для запуска приложения в режиме разработчика используйте эту команду в директории проекта:

```
wails3 dev
```

### Генерация TS биндингов

Для обновления API для TypeScript используйте команду:

```
wails3 generate bindings -ts
```

## Начало работы
Установите следующие утилиты перед началом работы:
```shell
go install -v github.com/wailsapp/wails/v3/cmd/wails3@latest
```
```shell
go install github.com/opbnq-q/nto-cli@latest
```
```shell
go install github.com/kuzgoga/crudgen/cmd/crudgen@latest
```
```shell
go install github.com/kuzgoga/gormlint/cmd/gormlint@latest
```

## Сборка

Для финальной сборки запустите эту команду в директории проекта:

На Linux/Mac OS:
```
PRODUCTION=true wails3 build
```

На Windows:
```
$env:PRODUCTION="true"; wails3 build
```

**Перед релизом не забыть**:
- убедиться, что дефолтные данные правильные
- убедиться, что приложение запускается
- (опционально) поместить все нужные asset'ы в папку assets
- приложить сопроводительную записку.

## Работа без GitHub

Настройте прокси для скачивания зависимостей через прокси:

```
go env -w GOPROXY="https://proxy.golang.org,direct"
```

## CRUD generator

Установите crudgen:

```
go install git.gogacoder.ru/NTO/crudgen/cmd/crudgen@latest
```

Сгенерируйте DAL:

```
cd gen
go run gen.go
```

Запустите crudgen:

```
crudgen -p internal
```

Не забудьте добавить новые модели в `Entities`, а сервисы в `Services`.
