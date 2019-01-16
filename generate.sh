#!/usr/bin/env bash#!/usr/bin/env bash

#   GitBook
gitbook serve --config book.json . gitbook
#gitbook build --config book.json . gitbook

#   all-contributor

npm install --save-dev all-contributors-cli
npm run contributors:generate
all-contributors add kylesliu code,blog,design,doc
all-contributors generate