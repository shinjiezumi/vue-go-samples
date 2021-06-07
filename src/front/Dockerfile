FROM node:14.17.0-alpine as build

WORKDIR /app

COPY . .

RUN apk update && \
    npm install && \
    npm run build

# --- production ---
FROM nginx:1.15.12-alpine as production
COPY --from=build /app/dist /usr/share/nginx/html
COPY --from=build /app/nginx/default.conf /etc/nginx/conf.d/default.conf
EXPOSE $PORT
CMD ["nginx", "-g", "daemon off;"]