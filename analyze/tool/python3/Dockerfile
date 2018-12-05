FROM python:3

WORKDIR /usr/src/analyzer

RUN mkdir -p /usr/src/analyzer/dist/run

ENV PYTHONUSERBASE=/usr/src/analyzer/dist/run

RUN python -m venv dist/venv && . ./dist/venv/bin/activate

COPY requirements.txt analyze.py ./

CMD [ "pip", "install", "--user", "--no-cache-dir", "--upgrade", "-r", "requirements.txt" ]
