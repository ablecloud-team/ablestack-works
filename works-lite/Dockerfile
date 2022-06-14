FROM node:10-alpine

RUN mkdir -p /app
WORKDIR /app

COPY package.json ./
RUN npm install
COPY . .

EXPOSE 8080
CMD [ "node", "index.js" ]
