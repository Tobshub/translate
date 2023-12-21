# Translate - A Language Translator CLI

Translates input text into a target language using [libretranslate](https://libretranslate.com)

## Installation

```bash
$ go install github.com/tobshub/translate
```

## Config

```bash
$ translate -config url=<libretranslate_server_url> key=<api_key> lang=<default_target_language_code>
```

### Defaults

- `url`: https://libretranslate.de
- `key`: (empty)
- `lang`: en

## Usage

```bash
$ translate -l <source_language> -t <target_language> <text>
```

If `-l` (source language) is omitted, it will be detected automatically
and the confidence in the language detection will be reported.

If `-t` (target language) is omitted,
it will be set to the default target language specified in the config.
