#!/usr/bin/env bash

export MYSQL_PWD="$MARIADB_ROOT_PASSWORD" 

mariadb -uroot -sse "SELECT SCHEMA_NAME FROM INFORMATION_SCHEMA.SCHEMATA WHERE SCHEMA_NAME = '$1';"
