#/bin/sh

go build .

cd ./web

bun install

bun run build

rm -rf ../backend/frontend

mkdir ../backend/frontend

mv ./dist/* ../backend/frontend

cd ..

go build

./Controller