BINARY_NAME=gclear
INSTALL_DIR=/usr/local/bin/

all: build 

install:
	go build -o ${BINARY_NAME}	
	sudo mv ${BINARY_NAME} ${INSTALL_DIR}

build:
	go build -o ${BINARY_NAME}

clean:
	go clean
	rm -f ${BINARY_NAME}
