# Card Service

### gRPC сервис для управления банковскими картами и пользователями с PostgreSQL базой данных.

Написан для практики создания автотестов на Go

# 🚀 Функциональность

### Методы

* CreateUser - Создание нового пользователя в системе
* GetUser - Получение информации о пользователе по ID
* GetAllUsers - Получение списка всех пользователей
* UpdateUser - Обновление информации о пользователе
* DeleteUser - Удаление пользователя из системы
* CreateCard - Создание новой банковской карты
* GetCard - Получение информации о карте по ID
* GetAllCards - Получение списка всех карт
* DeleteCard - Удаление банковской карты
* GenerateCard - Автоматическая генерация карты для пользователя

# 📋 Требования

* Go
* PostgreSQL
* Protocol Buffers (protoc)

# ⚒️ Установка

1. Установка Protocol Buffers

Следуйте [инструкции](https://www.geeksforgeeks.org/installation-guide/how-to-install-protocol-buffers-on-windows/)

2. Клонирование проекта

Выполнить в консоли команду: `git clone https://github.com/Vasto-GoQa/card_service`

Перейти внутрь проекта выполнив команду: ``cd card_service``

Создать папку для сгенерированных Protocol Buffers файлов: ``mkdir generated``

3. Установка Go плагинов из go.mod

Выполнить в консоли команду: ```go mod download```

4. Генерация Protocol Buffers файлов

Выполнить в консоли команду:
``protoc --go_out=generated --go_opt=paths=source_relative --go-grpc_out=generated --go-grpc_opt=paths=source_relative proto/card_service.proto``

5. Инициализация БД

Выполнить скрипт из ``init/db/init.sql`` в своей СУБД

6. Изменение настроек подключения БД

В файле ``config.go`` изменить значения на те, что используются в вашем подключении

6. Запуск сервера

При первом запуске выполнить команду: ``go mod tidy``

Выполнить в консоли команду: ``go run cmd/server/main.go cmd/server/config.go``

7. Запуск автотестов

⚒️В разработке⚒️

Author: [Vasto-GoQa](https://github.com/Vasto-GoQa)