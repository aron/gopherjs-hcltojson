vendor:
	go mod vendor

hcltojson.js hcltojson.js.map: vendor
	GOOS=linux $$GOPATH/bin/gopherjs build -m hcltojson.go

.PHONY: build
build: hcltojson.js

.PHONY: test
test: hcltojson.js
	node test.js

.PHONY: clean
clean:
	rm -rf hcltojson.js hcltojson.js.map vendor

.DEFAULT: build
