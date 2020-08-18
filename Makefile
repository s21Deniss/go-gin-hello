
grepSearch ?= s21deniss/go-gin-hello

consul:
	docker ps | grep $(grepSearch) | \
	awk '{ print $$1 }' | \
	xargs -I'{}' docker inspect --format '{{ .NetworkSettings.IPAddress }}' {} | \
	xargs -I'{}' bash -c "sed -i 's/dummyIP/{}/g' go-gin-hello-consul.json; consul services register go-gin-hello-consul.json; sed -i 's/{}/dummyIP/g' go-gin-hello-consul.json"
	