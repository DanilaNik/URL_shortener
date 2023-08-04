# URL Shortener 

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

![UrlShortener](https://github.com/DanilaNik/URL_shortener/tree/main/Media/URL-Short.jpg)

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

![UrlSaveOK](https://github.com/DanilaNik/URL_shortener/tree/main/Media/saveOK.png)

Сохранение с авторизацией и зарезервированным alias:

![UrlSaveError](https://github.com/DanilaNik/URL_shortener/tree/main/Media/saveError.png)

Сохранение без указания alias:
![saveNoAlias](https://github.com/DanilaNik/URL_shortener/tree/main/Media/saveNoAlias.png)

Сохранение с неудачной попыткой авторизации:

![NoAuth](https://github.com/DanilaNik/URL_shortener/tree/main/Media/NoAuth.png)


### Удаление <a name="save"></a>

Удаление существующего alias:

![deleteAlias](https://github.com/DanilaNik/URL_shortener/tree/main/Media/deleteAlias.png)

Удаление несуществующего alias:

![deleteNoAlias](https://github.com/DanilaNik/URL_shortener/tree/main/Media/deleteNoAlias.png)

### Редирект <a name="save"></a>

![redirect](/Media/redirect.gif)
