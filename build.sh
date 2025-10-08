#/bin/sh

go build .

cd ./web

bun install

bun run build

rm -rf ../frontend/*

mv ./dist/* ../frontend -f

cd ..

./Controller
