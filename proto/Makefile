format:
	buf format -w

lint:
	buf lint

lock:
	buf mod update

clean:
	rm -rf gen

generate: format lint clean
	buf generate && generate-go-embed4assets -directory gen
