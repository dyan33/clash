
install_clash:
	go install .


build_windows:
		export CGO_ENABLED=0;\
    	export GOOS=windows;\
    	export GOARCH=amd64;\
    	go  build  -ldflags '-w -s' .;\

# openapi 规则
install_openapi_extend:
	go build -o ${GOPATH}/bin/clash_openai_config cmd/clash_openai_config.go
