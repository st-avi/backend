include ./hack/hack-cli.mk
include ./hack/hack.mk

.PHONY: init
init:
	cp ./hack/config.yaml.example ./hack/config.yaml
	cp ./manifest/config/config.yaml.example ./manifest/config/config.yaml
	@echo "\033[32mInitialization complete. Please edit the config files as needed.\033[0m"

# 啟用 git hooks（commit 前自動執行 golangci-lint）
.PHONY: hooks
hooks:
	git config core.hooksPath hack/git-hooks
	@echo "\033[32mGit hooks enabled.\033[0m"
