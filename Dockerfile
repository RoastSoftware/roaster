FROM ubuntu:bionic
LABEL authors="William Wennerström <william@willeponken.me>; Philip Hjortsberg <philip@hjortsberg.me>"

RUN apt-get update && apt-get install -y python3 python3-pip

COPY ./build /roaster

RUN du -h -d1 /roaster

WORKDIR /roaster
RUN pip3 install --no-cache-dir -r requirements.freeze.txt

EXPOSE 5000
CMD ["/roaster/roasterd"]
