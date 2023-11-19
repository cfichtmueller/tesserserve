FROM golang:1.21.4-bookworm AS BUILDER

RUN apt-get update -qq && apt-get install -y -qq libtesseract-dev libleptonica-dev

ENV TESSDATA_PREFIX=/usr/share/tesseract-ocr/5/tessdata/

WORKDIR /build

ADD . .

RUN go build -o tesserserve ./cmd/server/main.go

FROM ubuntu:23.10

ENV TESSDATA_PREFIX=/usr/share/tesseract-ocr/5/tessdata/
ENV BIND_IP=0.0.0.0
ENV PORT=8000

RUN apt-get update -qq \
    && apt-get install -y -qq \
    libtesseract-dev \
    libleptonica-dev \
    tesseract-ocr-eng \
    tesseract-ocr-deu

COPY --from=BUILDER /build/tesserserve /usr/bin/tesserserve

ENTRYPOINT ["tesserserve"]