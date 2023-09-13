FRONTEND_DIR=frontend
BUILD_DIR=build

clean:
	cd $(FRONTEND_DIR); \
	if [ -d $(BUILD_DIR) ]; then rm -rf $(BUILD_DIR); fi

static: clean
	cd $(FRONTEND_DIR); \
	pnpm i; \
	pnpm build

build: 
	docker compose build

run:
	docker compose up

stop:
	docker compose down