default: build

new:
    export DBUSER="root"
    export DBPW="root"
    export DATABASEDB="cs53"
    export MAILAUTH="CHANGE THIS to whatever you like"
    export MAILUSER="csams.noreply@gmail.com"
    export MAILPW="CHANGE THIS to actual password"
    export MAILPROVIDER="smtp.gmail.com"

    sudo docker-compose up --build -d

build:
	git pull
	sudo docker-compose up --build -d

run:
	sudo docker-compose up -d

clean:
	sudo rm -r -f dbservice/data
	git pull
	sudo docker-compose up --build -d

stop:
	sudo docker-compose down

restart:
    sudo docker-compose down
    sudo docker-compose up -d