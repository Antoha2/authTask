FROM golang:1.18

WORKDIR /usr/src/app

COPY . .

EXPOSE 8181

CMD ["go","run", "cmd/main.go"]



##COPY . D:/code/etodo/docker

#EXPOSE 8080

#CMD [ "golang", "main.go" ]



#RUN mkdir D:/code/etodo/docker
#ADD . D:/code/etodo/docker
#WORKDIR D:/code/etodo/docker
#RUN go build -o main . 





