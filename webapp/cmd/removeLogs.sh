cp logs/mysql/mysql-slow.log logs/mysql/$(date +%Y%m%d-%H%M%S).log
cp logs/nginx/access.log logs/nginx/$(date +%Y%m%d-%H%M%S).access.log
cp logs/nginx/error.log logs/nginx/$(date +%Y%m%d-%H%M%S).err.log
cp logs/nginx/access.ndjson logs/nginx/$(date +%Y%m%d-%H%M%S).ndjson
echo > logs/mysql/mysql-slow.log
echo > logs/nginx/access.log
echo > logs/nginx/error.log
echo > logs/nginx/access.ndjson
