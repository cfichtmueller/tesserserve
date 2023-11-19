# tesserserve

> A simple Tesseract over REST server

## API

> To be documented.

In short:

```bash
curl --location 'http://localhost:8000/api/recognize?iterator-level=block&lang=deu' \
--header 'Content-Type: image/png' \
--data image-with-text.png'
```

Where iterator-level is one of `symbol`, `word`, `textline`, `para`, `block`.

Parameter `lang` is optional. It gives the engine a hint of what the expected language of the text is.


## Configuration

Configuration is done through the environment.

```bash
BIND_IP=127.0.0.1 # the IP the server binds to (optional)
PORT=8000         # the port the server binds to (optional)
```

## Running Locally (on Mac)

First, install tesseract and leptonica (via brew).  
Then point to the tesseract and leptonica libraries and header files

```
CGO_ENABLED=1 \
CGO_CXXFLAGS="-I/opt/homebrew/Cellar/leptonica/1.83.1/include -I/opt/homebrew/Cellar/tesseract/5.3.3/include" CGO_LDFLAGS="-L/opt/homebrew/Cellar/leptonica/1.83.1/lib -L/opt/homebrew/Cellar/tesseract/5.3.3/lib" \
go run cmd/server/main.go
```

## Known Issues

Running the software in docker on a mac (ARM) can result in very poor performance (>10s for completing a recognition task).
This issue is being ignored for now since it can't be reproduced in the targeted execution environment (x86 linux server).