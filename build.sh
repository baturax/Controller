#/bin/sh

cd ./web

bun run build

rm -rf ../frontend/*

mv ./dist/* ../frontend -f