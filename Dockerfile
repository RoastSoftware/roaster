FROM golang:1.11
LABEL authors="William Wennerström <william@willeponken.me>; Philip Hjortsberg <philip@hjortsberg.me>"

COPY ./deploy /roaster

WORKDIR /roaster

EXPOSE 5000
CMD ["/roaster/roasterd"]
