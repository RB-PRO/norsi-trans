all: run

run:
	go run cmd/main/main.go

push:
	git push git@github.com:RB-PRO/norsi-trans.git

pull:
	git pull git@github.com:RB-PRO/norsi-trans.git

pushW:
	git push https://github.com/RB-PRO/norsi-trans.git

pullW:
	git pull https://github.com/RB-PRO/norsi-trans.git

pushCar:
	scp main root@194.87.107.129:go/norsi-trans/

build-config:
	go env GOOS GOARCH

build-windows-to-linux:
	set GOARCH=amd64 set GOOS=linux go build .\cmd\main\main.go  

build-linux-to-windows:
	export GOARCH=amd64 export GOOS=windows go build .\cmd\main\main.go

build-car:
	export GOARCH=amd64
	export GOOS=linux
	export CGO_ENABLED=0
	go env GOOS GOARCH
	go build -o main ./cmd/main/main.go
	scp main Delivery.xlsx config_test.json consumer_key secret_key proxy root@194.87.107.129:go/norsi-trans/

build-carW:
	set GOARCH=amd64
	set GOOS=linux
	set CGO_ENABLED=0
	go env GOOS GOARCH
	go build -o main ./cmd/main/main.go
	scp main Delivery.xlsx config_test.json consumer_key secret_key proxy root@194.87.107.129:go/norsi-trans/