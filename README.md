# apitest
testing api stuff

NOTE: there are a number of ways to do this, i'm not trying to link containers or anything at the moment.
the end result will be 2 containers, one running redis and the other running the apitest app.

to build and run   
install docker https://docs.docker.com/installation/mac/  
git clone https://github.com/grahamgreen/apitest  
cd apitest  
docker build -t image_name  
this will build the app but you still need redis  
docker pull redis   
docker run -p 6379:6379 --name api-redis -d redis  
docker run --publish 8080:8080 --name test --rm apitest  

now you should be able to 
telnet 192.168.59.103 6379  
and connect to redis  
running the commnad "monitor" on redis will show you the commands as they come through
and  
curl -i -H "Content-Type: application/json" -d '{"name": "account name", "stuff": "some more stuff"}' http://192.168.59.103:8080/add/account/
to add accounts  
to add actions you'll need to note an account id and then run  
curl -i -H "Content-Type: application/json" -d '{"account": "<ACCOUNT_ID>","name": "the action name", "thingy": "someotherthings"}' http://192.168.59.103:8080/add/action/  
to make a donation you'll need 2 accounts and an action  
curl -i -H "Content-Type: application/json" -d '{"action": "<ACTION_ID>", "to": "<ACCOUNT_ID>", "from": "<DIFFERENT_ACCOUNT_ID>"}' http://localhost:8080/donate/
