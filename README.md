# Тестовое задание стажера в юнит AvitoPRO
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Feeonevision%2Favito-pro-test.svg?type=shield)](https://app.fossa.io/projects/git%2Bgithub.com%2Feeonevision%2Favito-pro-test?ref=badge_shield)


### Необходимо реализовать сервис на ~~PHP~~Golang для генерации случайных значений.

Сервис реализует JSON API работающее по HTTP. Каждой генерации присваивать уникальный id, по которому можно получить результат генерации методом retrieve.

Реализовать методы
* POST /api/generate/ - генерация случайного значения и его идентификатора
* GET /api/retrieve/ - получение зачения по id, которое вернулось в методе generate

Задача может быть решена со следующими усложнениями
* возможность задать входные параметры для метода /api/generate/
  - type - тип возвращаемого случайного значения (строка, число, guid, цифробуквенное, из заданных значений)
  - длина возвращаемого значения
* возможность идемпотентных запросов (несколько запросов с одним requestId вернут то же самое число)
* сервис поставляется как Docker образ, опубликованный в публичном репозитории
* написать Unit тесты

### Способы запуска сервиса

Тестовый запуск http-сервера с помощью CLI
```
go run cmd/avitornd/avitornd.go
```

Запуск через Docker-образ
```
docker build -t avitornd:local -f deployments/docker/avitornd-prod.Dockerfile .
docker run -p 8888:8888 avitornd:local
```

Запуск из публичного репозитория DockerHub
```
docker run -p 8888:8888 maintianone/avitornd:latest
```

### Документация к RESTful API
Документация написана по стандарту OASv3 и расположена в директории
```
api/openapi-spec
```

## License
[![FOSSA Status](https://app.fossa.io/api/projects/git%2Bgithub.com%2Feeonevision%2Favito-pro-test.svg?type=large)](https://app.fossa.io/projects/git%2Bgithub.com%2Feeonevision%2Favito-pro-test?ref=badge_large)