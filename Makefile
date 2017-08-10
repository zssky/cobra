all: cobra

mysql:
	sudo docker run  -e MYSQL_ROOT_PASSWORD=123456 --name mysql -d mysql

mssql:
	sudo docker run -e 'ACCEPT_EULA=Y' -e 'SA_PASSWORD=123456' -p 1433:1433 -d microsoft/mssql-server-linux

cobra: 
	go build -o bin/cobra ./server/main.go
