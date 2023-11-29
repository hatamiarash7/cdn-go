#!/bin/bash

cat .github/generator/.gitignore >>.gitignore
mv README.md docs/HOW-TO.md
cp .github/generator/*.md .
sed -i 's/GIT_USER_ID/arash-r1c/g' go.mod test/*.go docs/*.md
sed -i 's+GIT_REPO_ID+cdn-go-sdk+g' go.mod test/*.go docs/*.md
# sed -i 's+github.com/arvancloud+git.arvancloud.ir/arvancloud+g' go.mod test/*.go docs/*.md
sed -i 's+(docs/+(+g' docs/*.md
sed -i 's+../README.md#+HOW-TO.md#+g' docs/*.md
sed -i 's+(../README.md)+(HOW-TO.md)+g' docs/*.md
