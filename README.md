# final_project

Project structure 

cmd //main file
pb //protocol buffer files 
pkg
    client //clients for consuming grpc service 
    user_service
        entities //entities abstractions 
    user_service //main directory for services , endpoints , transport , coder/encoder and repository


# Generate proto files

Go to the pb directory and use the following command :

protoc --go_out=. --go-grpc_out=. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative user.proto

# Connect remotely to mysql

docker run -it --network main_network --rm mysql mysql -hmysql -uuser_db -p

