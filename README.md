# goga

Менеджер однофайловых модулей (версионирование copy-paste примеров). Использует Ваш git и gists для хранения пакетов.

* Альтернатика copy-paste
* Небольшие готовые решения можно подключать в проект целиком
* Легко вность изменения и контроллировать их и в истории самого проекта и в истории пакета, при этом остается связь с исходниом.
* Легко публиковать

## Пример использования

### Добавление модуля в проект

1. Найдите ссылку на нужный вам модуль - https://dapi.github.com/goga

2. Добавьте его в проект под нужным

Например так добавляется модуль `https://github.com/dapi/elements/blob/master/spinner.js` в файл `./app/javascripts/spinner.js`

> goga pull https://github.com/dapi/elements/blob/master/spinner.js ./app/javascripts/spinner.js

или можно упустить `pull`

> goga https://github.com/dapi/elements/blob/master/spinner.js ./app/javascripts/

`goga` скопирует файл как есть и добавит первой строкой специальный комментарий с адресом источника, например:

`// goga https://github.com/dapi/elements/blob/master/spinner/index.js`

Не удаляйте эту строку.

### Публикация изменений

> goga push ./app/javascripts/spinner.js

`goga` находит первую строку комментария с текстом `gogа`, берет из нее ссылку и пытается под вашими доступами запушать туда изменения.
