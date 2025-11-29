# syntax=docker/dockerfile:1.7

##
## Build stage
##
FROM node:20-alpine AS builder

WORKDIR /app

COPY package*.json ./
RUN npm ci

COPY . .
ENV NODE_ENV=production

RUN npm run build


##
## Runtime stage
##
FROM nginx:1.27-alpine AS runtime

WORKDIR /usr/share/nginx/html

# Configure nginx to serve the SPA on port 3800 with custom logs
COPY default.conf /etc/nginx/conf.d/default.conf

RUN mkdir -p /var/log/app && \
    chown -R nginx:nginx /var/log/app

COPY --from=builder /app/dist ./

EXPOSE 3800

CMD ["nginx", "-g", "daemon off;"]
