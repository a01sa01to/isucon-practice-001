Copy-Item logs/mysql/mysql-slow.log logs/mysql/$(date +%Y%m%d-%H%M%S).log
Copy-Item logs/nginx/access.log logs/nginx/$(date +%Y%m%d-%H%M%S).log
Copy-Item logs/nginx/error.log logs/nginx/$(date +%Y%m%d-%H%M%S).log
Copy-Item logs/nginx/access.ndjson logs/nginx/$(date +%Y%m%d-%H%M%S).ndjson

Clear-Content logs/mysql/mysql-slow.log
Clear-Content logs/nginx/access.log
Clear-Content logs/nginx/error.log
Clear-Content logs/nginx/access.ndjson
