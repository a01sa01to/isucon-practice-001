Get-ChildItem logs/mysql -Exclude mysql-slow.log | Remove-Item
Get-ChildItem logs/nginx -Exclude access.log, error.log, access.ndjson | Remove-Item
