FROM mysql:5.7

ENV MYSQL_ROOT_PASSWORD pass
ENV MYSQL_ROOT_HOST %

COPY ./scripts/ /docker-entrypoint-initdb.d/
EXPOSE 3306
