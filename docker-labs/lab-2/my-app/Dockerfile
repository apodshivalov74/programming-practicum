FROM node:20-alpine AS builder
WORKDIR /app

COPY package*.json ./
RUN npm install

ENV PATH /app/node_modules/.bin:$PATH

COPY . .
RUN npm run build

FROM nginx:alpine
COPY myapp_nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=builder /app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]