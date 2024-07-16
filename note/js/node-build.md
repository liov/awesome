
# typescript

## Locally in your project.
npm install -D typescript
npm install -D ts-node

## Or globally with TypeScript.
npm install -g typescript
npm install -g ts-node

## Depending on configuration, you may also need these
npm install -D tslib @types/node


# docker build node
docker run --rm --privileged=true -v /home/ghoper:/work -w /work/website/vhoper node:16-alpine3.16 pnpm run build
docker run -v /home/ghoper/static:/static --net=host --restart=always --cpus=0.2 -d --name vhoper  vhoper:1.2
