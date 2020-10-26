# gamecharacter

# Pre Setup
- Install Golang
- go to https://golang.org/dl/ download and install golang
- Install postgresql
- Install docker
- install postman
- go to https://www.getpostman.com/downloads/ download and install

# Install Package Library
- go mod init github.com/tmluthfiana/gamecharacter

# Setup
- Clone this Project git@github.com:tmluthfiana/gamecharacter.git
- Run it on local go run main.go
- Run with docker : docker-compose up --build

# Usage
- use postman to test it. 
get all data => GET : http://localhost:8080/items
update data => PUT : http://localhost:8080/items/{id}
create data = > POST : http://localhost:8080/items

# automated tests
- running file test in tests/modeltest folder with : go test -v -run (function name). example : go test -v -run TestFindAllItems
- running with docker test : docker-compose -f docker-compose.test.yml up --build --abort-on-container-exit