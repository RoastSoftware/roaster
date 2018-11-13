FROM alpine:edge
LABEL authors="William Wennerström <william@willeponken.me>; Philip Hjortsberg <philip@hjortsberg.me>"

RUN apk --no-cache add go git musl-dev openssh-client ca-certificates nodejs nodejs-npm

ENV GO111MODULE=on

COPY . /roaster
WORKDIR /roaster

RUN go build github.com/LuleaUniversityOfTechnology/2018-project-roaster/cmd/roasterd
RUN cp roasterd /usr/bin/roasterd

WORKDIR /roaster/www
RUN npm install
RUN npm run build

WORKDIR /roaster
EXPOSE 5000
CMD ["/usr/bin/roasterd"]
