gen: 
	protoc \
	    --plugin=./bin/protoc-gen-go \
	    --go_out=. \
		--go_opt=paths=source_relative \
	    location.proto