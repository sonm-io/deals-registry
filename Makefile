
GO=go
TRUFFLE=./node_modules/truffle/build/cli.bundled.js
TESTRPC=./node_modules/ethereumjs-testrpc/build/cli.node.js

default:
	${TRUFFLE} compile

node_modules:
	npm install

testrpc:
	nohup ${TESTRPC} -l ${GAS_LIMIT} &

test: testrpc
	${TRUFFLE} test

test_ci: node_morules testrpc test

deploy:
	cp -R build deployed
