run:
	docker run -p 3000:3000 --rm ascii 

build:
	docker build -t ascii .
