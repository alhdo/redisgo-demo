FROM node:18-alpine3.18
WORKDIR /app
COPY package.json package-lock.json ./
RUN npm ci


COPY . .

EXPOSE 8000

CMD [ "npm", "run", "dev"]