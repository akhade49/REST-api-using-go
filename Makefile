generate-mock:
	go generate -v ./...


clean:
	go clean -testcache

test-simple: clean
	 cd repository && go test ./... && cd ..

unit-test:
	cd service && go test ./... && cd ..

