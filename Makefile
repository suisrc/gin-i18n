.PHONY: start build

NOW = $(shell date -u '+%Y%m%d%I%M%S')

APP = i18n
RELEASE_SERVER = example/${APP}

run: start

# 初始化mod
mod:
	go mod init github.com/suisrc/gin-${APP}

install:
	go install ./example

# 修正依赖
tidy:
	go mod tidy

build:
	go build -ldflags "-w -s" -o $(SERVER_BIN) ./example

start:
	go run example/example.go

# go get -u github.com/nicksnyder/go-i18n/v2/goi18n
# goi18n -help
i18n:
	goi18n extract -outdir example -sourceLanguage zh-CN
i18n-1:
	cd example && touch translate.en-US.toml translate.ja-JP.toml
i18n-2:
	cd example && goi18n merge -sourceLanguage zh-CN active.zh-CN.toml translate.en-US.toml translate.ja-JP.toml
i18n-en:
	cd example && goi18n merge -sourceLanguage en-US active.en-US.toml translate.en-US.toml
i18n-ja:
	cd example && goi18n merge -sourceLanguage ja-JP active.ja-JP.toml translate.ja-JP.toml