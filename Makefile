FORCE:
.DEFAULT_GOAL := help

define USAGE
  conductor
endef

clean:

help:
	$(info $(USAGE))
	@:

include protos/Makefile
include queue/Makefile
include runner/Makefile

clean-all ca: clean

.PHONY: clean clean-all ca
