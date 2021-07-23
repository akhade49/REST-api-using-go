generate-mock:
	go generate -v ./...


clean:
	go clean -testcache


test-simple: clean
	 cd repository && go test ./... && cd ..

