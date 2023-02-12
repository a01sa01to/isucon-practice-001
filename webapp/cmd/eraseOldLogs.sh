find logs/mysql -type f -not -name "mysql-slow.log" -exec rm -f {} \;
find logs/nginx -type f -not -name "access.log" -not -name "error.log" -not -name "access.ndjson" -exec rm -f {} \;
