# URL Shortener 

[![forthebadge](https://forthebadge.com/images/badges/made-with-go.svg)](https://forthebadge.com)

![UrlShortener][UrlShortener]

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

![UrlSaveOK][UrlSaveOK]


Сохранение с авторизацией и зарезервированным alias:

![UrlSaveError][UrlSaveError]

Сохранение без указания alias:
![saveNoAlias][saveNoAlias]

Сохранение с неудачной попыткой авторизации:

![NoAuth][NoAuth]


### Удаление <a name="save"></a>

Удаление существующего alias:

![deleteAlias][deleteAlias.png]

Удаление несуществующего alias:

![deleteNoAlias][deleteNoAlias]

### Редирект <a name="save"></a>

![redirect](/Media/redirect.gif)


[UrlShortener]: https://github.com/DanilaNik/URL_shortener/tree/main/Media/URL-Short.jpg
[UrlSaveOK]: https://github.com/DanilaNik/URL_shortener/tree/main/Media/saveOK.png
[UrlSaveError]: https://github.com/DanilaNik/URL_shortener/tree/main/Media/saveError.png
[NoAuth]: https://github.com/DanilaNik/URL_shortener/tree/main/Media/NoAuth.png
[saveNoAlias]: https://github.com/DanilaNik/URL_shortener/tree/main/Media/saveNoAlias.png
[deleteAlias]: https://github.com/DanilaNik/URL_shortener/tree/main/Media/deleteAlias.png
[deleteNoAlias]: https://github.com/DanilaNik/URL_shortener/tree/main/Media/deleteNoAlias.png
