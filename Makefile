DEST=out/0.0.1
#DEST=$(HOME)/.terraform.d/plugins/registry.terraform.io/zicong/zicong/0.0.1

build:
	rm -f terraform-provider-zicong
	go build -o terraform-provider-zicong

install: build
	mkdir -p ${DEST}
	mv terraform-provider-zicong ${DEST}

clean:
	rm -rf ${DEST}

