build:
	rm -f terraform-provider-zicong
	go build -o terraform-provider-zicong

install: build
	mkdir -p out/0.0.1
	mv terraform-provider-zicong out/0.0.1

clean:
	rm -rf out/0.0.1

