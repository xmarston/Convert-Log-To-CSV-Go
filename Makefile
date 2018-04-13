PLATFORMS := linux/amd64 windows/amd64 linux/arm darwin/amd64
PACKAGE := convert_log_to_csv
VERSION := 0.0.1

temp = $(subst /, ,$@)
os = $(word 1, $(temp))
arch = $(word 2, $(temp))

.PHONY: release $(PLATFORMS)
release: $(PLATFORMS)

$(PLATFORMS):
	GOOS=$(os) GOARCH=$(arch) go build -o '$(PACKAGE)-$(VERSION)-$(os)-$(arch)' $(PACKAGE)