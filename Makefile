.PHONY: help
help:
	@echo "help:"
	@echo "use [ make gen ] cmd to gen {cmd}_test.go file, if it not exists"

.PHONY: gen 
gen:
	@python3 scripts/gen_cmd_test.py

.DEFAULT_GOAL=help