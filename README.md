## README

1. Change directory into the cloned repository
  
    `$ cd $GOPATH/github.com/danshao/ppapp`
   
2. Run MySQL Docker Container
    
    `docker run --name mysql -d -e MYSQL_ROOT_PASSWORD=root mysql:5.7`

3. Give MySQL about 15 seconds to boot. Then create database.

    `docker run -it --link mysql:mysql --rm mysql sh -c 'exec mysql -uroot -hmysql -proot -e "create database pp;"'`

3. Build API Server Docker Container
  
    `docker build ./ --build-arg app_env=production -t ppapp`

4. Run API Server Docker Container
  
    `docker run -d --name ppapp --link mysql:mysql -p 8080:8080 ppapp`
