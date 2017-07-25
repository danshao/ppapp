## README

1. Run MySQL Docker Container

   `docker run --name mysqldb --restart always -p 3306:3306 -e MYSQL_ROOT_PASSWORD=root -d mysql:5.7`

2. Check MySQL container IPAddress

    `docker inspect mysqldb | grep IPAddr`

3. Modify MySQL Connection IP Address in `db/connection.go` 
    
    `MYSQLIPADDR  = "MYSQL_CONTAINER_IP_ADDRESS"`

4. Change directory into the cloned repository
  
   `$ cd $GOPATH/github.com/danshao/ppapp`

4. Build API Server Docker Container
  
   `docker build ./ --build-arg app_env=production -t ppapp`

5. Run API Server Docker Container
  
   `docker run -d --name ppapp -p 8080:8080 ppapp`