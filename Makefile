VERSION = `git rev-parse --short HEAD`

SHELL := /bin/bash

all: test-lcof test-leetcode
	@echo all
	#make test-nowcoder
	#make test-leetcode

test:
	@echo start test
	sh test.sh
	@echo end  test
 
hugo:
	@echo start hugo
	hugo server
	@echo end  hugo
deploy:
	@echo start deploy
	git submodule update --init --recursive 
	git pull
	hugo -D
	@echo end  deploy 

