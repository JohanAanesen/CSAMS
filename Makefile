build:
    sudo docker-compose up --build

run:
    sudo docker-compose up

clean:
    sudo rm -r -f dbservice/data
    sudo docker-compose up --build