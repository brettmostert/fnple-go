#!/usr/bin/env bash

source "$(dirname "$0")/common.sh"
getProjectDir

echo -e "==> ${CYAN}Bootstrap...${NC}"

if [[ -d ${DIR}/bin ]]; then
	rm -rf ./tmp
fi

echo -e "==> Installing golint'"
brew install golangci-lint

echo -e "==> Installing 'gofumpt'"
go install mvdan.cc/gofumpt@latest

echo -e "==> Setup Git Hooks"
git config core.hooksPath .githooks
cd "${DIR}" || exit

echo -e "==> ${YELLOW}Done... ${GREEN}ALL GOOD!!!${NC}"
echo
exit 0
