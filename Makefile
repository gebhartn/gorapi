SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c

ifneq ("$(wildcard .env)","")
	include .env
	export
endif

NO_COLOR		:= \x1b[0m
ERROR_COLOR	:= \x1b[31;01m

env-%:
	@if [ "${${*}}" = "" ]; then \
		printf "$(ERROR_COLOR)"; \
		echo "**** ERROR: Required environment variable $* not set ****"; \
		printf "$(NO_COLOR)"; \
		echo; \
		exit 1; \
	fi

build: env-APP_DB_USERNAME env-APP_DB_PASSWORD env-APP_DB_NAME
	go build *.go
