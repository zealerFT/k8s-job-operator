# Use Node.js to build frontend
FROM node:latest as build-stage
WORKDIR /app/frontend/
COPY frontend/package*.json ./
RUN npm install
COPY frontend/ ./
RUN npm run build

FROM alpine:latest

MAINTAINER fermi <spuerfei@gmail.com>

    LABEL name="k8s-job-operator" \
description="k8s-job-operator service by fermi" \
owner="spuerfei@gmail.com"

WORKDIR /bin

ADD bin/k8s-job-operator /bin/k8s-job-operator
ADD bin/zoneinfo.zip /bin/zoneinfo.zip
COPY config/* /bin/config/
# Copy built frontend from build stage
COPY --from=build-stage /app/frontend/build /bin/frontend/build

COPY docker/entrypoint.sh /entrypoint.sh
# 处理 TimeLocation 无法找到的问题
ENV ZONEINFO=/bin/zoneinfo.zip
ENTRYPOINT ["/entrypoint.sh"]
