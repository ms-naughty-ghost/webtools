.PHONY: build up utils stop

ifeq ($(ENV),)
ENV := develop
endif

COMPOSE := ENV=$(ENV) docker-compose -f docker-compose.yaml
build:
	$(COMPOSE) build

up:
	$(COMPOSE) up -d

utils:
	$(COMPOSEUTIL) up -d

stop:
	$(COMPOSE) stop

down: stop 
	$(COMPOSE) down