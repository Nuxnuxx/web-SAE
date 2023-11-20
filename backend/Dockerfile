FROM node:18-alpine

RUN apk add dumb-init

COPY ./src/ /app/src/
COPY ./package*.json /app/
COPY ./tsconfig.json /app/

WORKDIR /app/
RUN npm install
RUN npm run build



EXPOSE 5001

USER node

CMD dumb-init npm run start
