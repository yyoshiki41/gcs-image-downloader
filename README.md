# downloader - Bulk Image DownLoader.

Download images using Google Custom Search API.
It's fast because of downloading multiple files in parallel using goroutines.


## Installation

```shell
$ go get -u github.com/yyoshiki41/gcs-image-downloader/cmd/downloader
```

## Usage

```shell
$ downloader --help
usage: downloader --query=QUERY [<flags>]

Image downloader for Google Custom Search API.

Flags:
      --help               Show context-sensitive help (also try --help-long and --help-man).
  -c, --conf="conf"        Config file path
  -o, --outputs="outputs"  Outputs directory
  -q, --query=QUERY        Query
  -n, --number=10          Number of files
  -s, --safe=SAFETY-LEVEL  Safety level: high, medium, off
  -t, --type=IMG-TYPE      Images of a type: clipart, face, lineart, news, photo
```

## Example

```shell
$ downloader -q gopher
2016/10/18 23:59:06 Start!
Number: 10
2016/10/18 23:59:23 Download has completed!
Total: 10, Success: 10, Failure: 0
```

## Settings

[Docs](https://developers.google.com/custom-search/docs/overview).

1. Create a project in the [Google Developers Console](https://console.cloud.google.com/) and turn on the Custom Search API
2. [Get API Key](https://console.cloud.google.com/apis/credentials)
3. [Create new Custom Search Engine](https://cse.google.com/cse/all) and get Custom Search Engine ID


## Configuration

```shell
$ cd $GOPATH/src/github.com/yyoshiki41/gcs-image-downloader
$ cp ./conf/credentials.toml.skel ./conf/credentials.toml
# Edit a config file
$ vim ./conf/credentials.toml
```

## License 

The MIT License

## Author

Yoshiki Nakagawa
