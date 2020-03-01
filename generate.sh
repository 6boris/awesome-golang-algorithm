#!/usr/bin/env bash#!/usr/bin/env bash

# Git
#git checkout master
#git reset --hard
#git fetch --all
#git reset --hard origin/master
git pull

#   GitBook
#npm install gitbook-cli -g
#gitbook serve --config book.json . gitbook
#gitbook build --config book.json . gitbook

#   all-contributor

#npm install --save-dev all-contributors-cli
#npm run contributors:generate
#all-contributors add kylesliu code,blog,design,doc
#all-contributors generate

rm -rf public_new
rm -rf gitbook

mkdir public_new
mkdir gitbook

cp -rfv src/* gitbook
cp README.md gitbook
cp SUMMARY.md gitbook
cp SUMMARY-LIST.md gitbook
cp CONTRIBUTOR.md gitbook
gitbook build --config book.json gitbook public_new --timing --log debug
#gitbook serve --config book.json gitbook public

# nohup sh generate.sh >generate.log 2>&1 &

rm -rf public && mv public_new public
