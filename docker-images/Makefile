images += alpine
images += prometheus
images += grafana

default:
	@echo Possible targets: $(images)

$(images):
	[ ! -f $@/Makefile ] || (cd $@; $(MAKE))
	$(eval tag := $(shell ./version.sh $@))
	$(eval image := sourcegraph/$@)
	docker build -t $(image):$(tag) $@
	docker tag $(image):$(tag) $(image):latest
	docker push $(image):$(tag)
	docker push $(image):latest

.PHONY: $(images)
