 A golang web project
 * based on beego
 * using redis/mysql
 
 # dev mode
  cd src/myapp && bee run
# prod mode
  build.sh
  cd src/myapp &&./myapp
# publish to docker hub
  docker_publish.sh
# deploy in k8s
  https://github.com/yc-alex-xu/microk8s  
    