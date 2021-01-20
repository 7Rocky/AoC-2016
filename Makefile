test:
	@for d in $(shell ls -d ${PWD}/*/); do \
		cd $${d} && go test; \
	done

.PHONY: test
