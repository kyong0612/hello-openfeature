run:
	go build -o hello-openfeatures && ./hello-openfeatures

flagd:
	docker run -p 8013:8013 \
	-v ./:/etc/flagd/ \
	-it ghcr.io/open-feature/flagd:latest start \
	--uri file:/etc/flagd/flags.flagd.json

test-curl:
	curl http://localhost:8080/hello
