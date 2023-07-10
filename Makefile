
install_clash:
	go install -a  .


build_windows:
		export CGO_ENABLED=0;\
    	export GOOS=windows;\
    	export GOARCH=amd64;\
    	go  build  -ldflags '-w -s' .;\


build_linux:
			export CGO_ENABLED=0;\
        	export GOOS=linux;\
        	export GOARCH=amd64;\
        	go  build  -ldflags '-w -s' -o clash_linux .;\

# openapi 规则
install_openapi_extend:
	go build -o ${GOPATH}/bin/clash_openai_config cmd/clash_openai_config.go
