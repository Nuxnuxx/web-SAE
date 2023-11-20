FROM node:18-alpine AS builder

COPY package*.json /app/
COPY ./tsconfig.json /app/
COPY ./svelte.config.js /app/
COPY ./vite.config.ts /app/
COPY ./src/ /app/src/
COPY ./static /app/

WORKDIR /app

RUN npm ci
RUN npm run build
RUN npm prune --production

EXPOSE 3000

ENV NODE_ENV=production


USER node
CMD [ "node", "build" ]
