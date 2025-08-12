include ./hack/hack-cli.mk
include ./hack/hack.mk

init:
	cp ./hack/config.yaml.example ./hack/config.yaml
	cp ./manifest/config/config.yaml.example ./manifest/config/config.yaml
	cp .env.example .env
	@echo "\033[32mInitialization complete. Please edit the config files as needed.\033[0m"