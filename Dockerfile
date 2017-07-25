FROM golang

ARG app_env
ENV APP_ENV $app_env

COPY . /go/src/github.com/danshao/ppapp
WORKDIR /go/src/github.com/danshao/ppapp

RUN go get ./
RUN go build

CMD if [ ${APP_ENV} = production ]; \
	then \
	app; \
	else \
	go get github.com/pilu/fresh && \
	fresh; \
	fi
	
EXPOSE 8080