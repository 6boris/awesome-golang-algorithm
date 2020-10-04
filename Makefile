VERSION = `git rev-parse --short HEAD`

SHELL := /bin/bash

all: test-nowcoder test-leetcode
	@echo all
	#make test-nowcoder
	#make test-leetcode

test-nowcoder:
	@echo start test neworder 
	sh test.sh nowcoder
	@echo end  test neworder 

test-leetcode:
	@echo start test leetcode 
	sh test.sh leetcode
	@echo end  test leetcode 
		

