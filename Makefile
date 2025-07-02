user-rpc-dev:
	@make -f deploy/mk/user-rpc.mk release-test
user-api-dev:
	@make -f deploy/mk/user-api.mk release-test
social-rpc-dev:
	@make -f deploy/mk/social-rpc.mk release-test
social-api-dev:
	@make -f deploy/mk/social-api.mk release-test
im-rpc-dev:
	@make -f deploy/mk/im-rpc.mk release-test
im-api-dev:
	@make -f deploy/mk/im-api.mk release-test
im-ws-dev:
	@make -f deploy/mk/im-ws.mk release-test
user-api-test:
	@make -f deploy/mk/user-api.mk run-test
user-rpc-test:
	@make -f deploy/mk/user-rpc.mk run-test
social-rpc-test:
	@make -f deploy/mk/social-rpc.mk run-test
social-api-test:
	@make -f deploy/mk/social-api.mk run-test
im-rpc-test:
	@make -f deploy/mk/im-rpc.mk run-test
release-test: user-rpc-dev user-api-dev social-rpc-dev social-api-dev im-rpc-dev im-api-dev im-ws-dev

install-server:
	cd ./deploy/script && chmod +x release-test.sh && ./release-test.sh



