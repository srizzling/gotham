proto:
	for d in services; do \
		for f in $$d/**/proto/*.proto; do \
			protoc --go_out=plugins=micro:. $$f; \
			echo compiled: $$f; \
		done \
	done

build:
	for d in services; do \
		echo "building $1/$d" \
		pushd $d/$1>/dev/null \
		rocker build -var SERVICE=$1 \
		popd >/dev/null \
	done
