FROM node:8

ENV PORT 8080
EXPOSE 8080

COPY . /app
WORKDIR /app

RUN npm install

CMD ["node", "app.js"]