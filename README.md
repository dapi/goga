![goga logo](https://raw.githubusercontent.com/dapi/goga/master/assets/goga.png)

# goga

Языконезависимый менеджер однофайловых пакетов (версионированный `copy-paste`)

[![Build Status](https://travis-ci.org/dapi/goga.svg?branch=master)](https://travis-ci.org/dapi/goga)

## Зачем?

Бывают участки кода, которые, слегка модифицируясь копируются из проекта в проект. Ради них не рационально заводить и регистрировать отдельный пакет в типовых пакетных менеджерах, потому что:

1. Уходит время на создание и поддержку - необходимо создать тестовую инфраструктуру, описать спецификации, фактически придется создать ещё один проект
2. Придется отдельно вносить изменения в самом пакете, обновлять пакет в основном приложении и заново заниматься его интеграционным тестированием, а если пойдет что-то не так, обрытно переключаться в проект пакета, вносить там изменения в тесты и код, публиковать и тп.
3. Часто бывает, что проще слегка модифицировать код под новый проект, чем создавать универсальное решение подходящее всем проектам.

При этом хочется видеть историю изменений, централизованно хранить и делиться такими участками кода. Для таких случаев подходит `goga`.

`goga` это не замена `gem`, `bundle`, `npm`, `yarn` и тп, а дополнение. Он живет рядом и хорошо делает свою маленькую работу.

`goga`-модуль это обычный исходник на любом языке программирования, в который первой строкой автоматически добавлен `goga`-комментарий с адресом источника.

## Что на борту?

* Подключение однофайловых пакетов одной командой.
* Использует Ваш git и gists для хранения пакетов.
* Публикация пакета в общий репозиторий одной командой.
* Живет в системе контроля версий вашего приложения совместно с основным пакетным менеджером.
* Не зависит от языка программированя.
* Легко вность изменения в исходный код пакетов.
* Мультиплатформенное решение, написано на `golang`

## Установка

1. Установите `golang` в вашу ОС по [инструкции](https://golang.org/doc/install)

2. Установите `goga`

> go get github.com/dapi/goga

## Использование

Список всех команд:

> goga

### goga add <URL> [local destination]

Скачивает по ссылке файл, кладет по указанному пути и добавляет первой строкой комментарий с ссылкой на источник.

Например: 

> goga add https://github.com/dapi/elements/blob/master/spinner.js ./app/javascripts/

Далее вы можете подключать и перемещать файл по проекту как вам удобно, используя проектную систему контроля версий.

### goga status <dir>

Сканирует указанную или текущую директорию на `goga`-модули (файлы с магическим комментарием) и сравнивает их с источником, сообщает если нашел изменения. Например:

```sh
> go run ./goga.go status
Scanning directory: /home/danil/code/goga_samples/
Found /home/danil/code/goga_samples/spinner.js checking - 4 diffs found
```

### goga push <file>

Заливает измененния в удаленный источник указанный в магическом комментарии. Например:

> goga push ./app/javascripts/spinner.js

### goga diff <file>

Показывает разницу между локальным файлом и его источником. Например:

> goga diff

### goga syntax

Выводит список поддерживаемых расширений и видов комментариев. Например:

```sh
> go run syntax

List of available file extensions and its comments syntax:
.py     # goga URI
.pl     # goga URI
.sh     # goga URI
.haml   // goga URI
.c      // goga URI
.cs     // goga URI
.go     // goga URI
.php    // goga URI
.sql    -- goga URI
.xml    <!-- goga URI -->
.html   <!-- goga URI -->
.swift  -- goga URI
.js     // goga URI
.java   // goga URI
.slim   // goga URI
.rb     # goga URI
```

## Магический комментарий

При скачивании модуля через `goga` он автоматически добавляет первой строкой магичекий комментарий со ссылкой на источник, например:

```javascript
// goga https://github.com/dapi/elements/blob/master/spinner.js
```

## Планы

* [ ] Поддержка доступа к git-репозиторию по логин/пароль (сейчас только ssh-ключи)
* [ ] Поддержка gist
* [ ] Поддержка автоматического определения gitlab/bitbucket-репозиториев.
* [ ] Улучшить вывод diff
* [ ] shell-автокомплит

## Поддерживаемые источники модулей

* Любая прямая ссылка на файл, например: https://site.com/spinner.js
* Ссылка на исходник в github репозитории наблюдаемый в браузере типа https://github.com/dapi/elements/blob/master/spinner.js автоматически преобразуется в https://raw.githubusercontent.com/...

## Языковая поддержка

На данный момент поддерживюатся исходники следующих языков программирования
(определяется через расширение файла) - https://github.com/dapi/goga/blob/master/cmd/syntax.go#L31

Если Вам нужна поддержка других языкой - создайте issue или пришлите PR
