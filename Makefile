#include .env

up:
	@docker compose up

down:
	@docker compose down -v

background:
	@docker compose up -d

.PHONY: up down background