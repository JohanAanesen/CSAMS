default: build

build:
	sudo docker-compose up --build -d

run:
	sudo docker-compose up -d

clean:
	sudo rm -r -f dbservice/data
	git pull
	sudo docker-compose up --build -d

stop:
	sudo docker-compose down