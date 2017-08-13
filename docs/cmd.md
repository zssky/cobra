# sqlserver相关操作  
## sqlserver 命令行工具安装说明  
参考文档：https://docs.microsoft.com/en-us/sql/linux/sql-server-linux-setup-tools#ubuntu  
步骤:  
curl https://packages.microsoft.com/keys/microsoft.asc | sudo apt-key add -  
curl https://packages.microsoft.com/config/ubuntu/16.04/prod.list | sudo tee /etc/apt/sources.list.d/msprod.list  
sudo apt-get update  
sudo apt-get install mssql-tools unixodbc-dev  
echo 'export PATH="$PATH:/opt/mssql-tools/bin"' >> ~/.bash_profile  
source ~/.bashrc  

执行完之后即可正常的使用sqlcmd访问了  

## sqlcmd连接sqlserver  
执行以下命令输入密码即可  
sqlcmd -S tcp:172.17.0.3,1433 -U sa  

## sqlcmd命令  
SELECT name from sys.databases 查询所有的数据列表  
SELECT Name FROM SysObjects Where XType='U' ORDER BY Name 查询数据库下面所有的用户表  

