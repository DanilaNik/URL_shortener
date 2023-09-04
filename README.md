# URL Shortener 

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

![UrlShortener](/Media/URL-Short.jpg)

### Задача:
Микросервис для сокращения ссылок, сохранение и удаление ссылок и их псевдонимов, редирект по псведонимам.(Для сохранения и удаления URL необходима авторизация)

### Используемые технологии:
- Go v1.20
- Роутер                - [go-chi/chi](https://github.com/go-chi/chi) 
- SQLite        
- Драйвер БД            - [go-sqlite3](github.com/mattn/go-sqlite3)
- Валидатор пакетов     - [go-playground/validator](https://github.com/go-playground/validator)
- Логгер                - [slog](https://pkg.go.dev/golang.org/x/exp/slog)   

### Examples
- [Сохранение](#save)
- [Удаление](#delete)
- [Редирект](#redirect)

### Сохранение <a name="save"></a>

Сохранение с авторизацией и не зарезервированным alias:

![UrlSaveOK](/Media/saveOK.png)

Сохранение с авторизацией и зарезервированным alias:

![UrlSaveError](/Media/saveError.png)

Сохранение без указания alias:
![saveNoAlias](Media/saveNoAlias.png)

Сохранение с неудачной попыткой авторизации:

![NoAuth](/Media/NoAuth.png)


### Удаление <a name="delete"></a>

Удаление существующего alias:

![deleteAlias](/Media/deleteAlias.png)

Удаление несуществующего alias:

![deleteNoAlias](/Media/deleteNoAlias.png)

### Редирект <a name="redirect"></a>

![redirect](/Media/redirect.gif)
