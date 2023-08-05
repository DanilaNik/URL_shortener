# URL Shortener 

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

![UrlShortener](/Media/URL-Short.jpg)

Микросервис для сокращения ссылок, сохранение и удаление ссылок и их псевдонимов, редирект по псведонимам.(Для сохранения и удаления URL необходима авторизация)

Используемые технологии:
- SQLite (в качестве хранилища данных)
- Chi (веб фреймворк)
- go-sqlite3 (драйвер для работы с SQLite)
- golang/testify (для тестирования)

## Examples
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
