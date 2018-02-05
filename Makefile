proto:
	for d in */; do \
		for f in $$d/**/proto/*.proto; do \
			if [ -f "$${GOPATH}/src/github.com/srizzling/gotham/$$f" ]; \
  			then \
    			protoc --proto_path="$${GOPATH}/src" --go_out=plugins=micro:$${GOPATH}/src $${GOPATH}/src/github.com/srizzling/gotham/$$f; \
				echo compiled: $$f; \
			fi\
		done \
	done

build-srv:
	for d in services/*; do \
		echo "building $$d"; \
		docker build --build-arg SERVICE=$$d .; \
	done
