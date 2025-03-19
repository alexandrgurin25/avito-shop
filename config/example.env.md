
##### Пример файла .env для настройки подключения к PostgreSQL и аутентификации

* Адрес хоста PostgreSQL, либо для докера можно указать для связи между контейнерами имя сервиса
    ```
    POSTGRES_HOST=postgresql
    ```

* Порт для подключения к PostgreSQL
    ```
    POSTGRES_PORT=5432
    ```

* Имя пользователя для подключения к PostgreSQL
    ```
    POSTGRES_USER=postgres
    ```

* Пароль для подключения к PostgreSQL
    ```
    POSTGRES_PASSWORD=102104
    ```

* Имя базы данных PostgreSQL
    ```
    POSTGRES_DB=shop
    ```

* Минимальное количество соединений в пуле
    ```
    POSTGRES_MIN_CONN=5
    ```

* Максимальное количество соединений в пуле
    ```
    POSTGRES_MAX_CONN=10
    ```

* Секретный ключ для аутентификации
    ```
    AUTH_SECRET_KEY="avito"
    ```

##### Пример файла test.env для настройки подключения к PostgreSQL и аутентификации
>Данный файл нужен, чтобы проводить интеграционные тесты и гарантировать, что с нашей основной базой данных ничего не случиться

* Адрес хоста PostgreSQL, либо для докера можно указать для связи между контейнерами имя сервиса
    ```
    POSTGRES_HOST=postgresqlTest
    ```

* Порт для подключения к PostgreSQL
    ```
    POSTGRES_PORT=5433
    ```

* Имя пользователя для подключения к PostgreSQL
    ```
    POSTGRES_USER=postgres
    ```

* Пароль для подключения к PostgreSQL
    ```
    POSTGRES_PASSWORD=102104
    ```

* Имя базы данных PostgreSQL
    ```
    POSTGRES_DB=TestShop
    ```

* Минимальное количество соединений в пуле
    ```
    POSTGRES_MIN_CONN=5
    ```

* Максимальное количество соединений в пуле
    ```
    POSTGRES_MAX_CONN=10
    ```

* Секретный ключ для аутентификации
    ```
    AUTH_SECRET_KEY="avito"
    ```