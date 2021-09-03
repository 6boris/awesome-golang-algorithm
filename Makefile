VERSION = `git rev-parse --short HEAD`

SHELL := /bin/bash

all: test-lcof test-leetcode
	@echo all
	#make test-nowcoder
	#make test-leetcode

test-lcof:
	@echo start test lcof 
	sh test.sh lcof
	@echo end  test lcof 

test-leetcode:
	@echo start test leetcode 
	sh test.sh leetcode
	@echo end  test leetcode 
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

