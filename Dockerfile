FROM golang:1.11
LABEL authors="William Wennerström <william@willeponken.me>; Philip Hjortsberg <philip@hjortsberg.me>"

COPY ./build /roaster

WORKDIR /roaster

RUN du -h -d1 *

EXPOSE 5000
CMD ["/roaster/roasterd"]
