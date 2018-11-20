FROM golang:1.11
LABEL authors="William Wennerström <william@willeponken.me>; Philip Hjortsberg <philip@hjortsberg.me>"

COPY ./bin/roasterd /roaster/roasterd
COPY ./www /roaster/www

WORKDIR /roaster

EXPOSE 5000
CMD ["/roaster/roasterd"]
