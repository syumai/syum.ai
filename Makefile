PROJECT_NAME="syum-ai"
PORT?=8080

.PHONY: dev
dev:
	dev_appserver.py app.yaml --port=$(PORT)

.PHONY: deploy
deploy:
	gcloud app deploy --project $(PROJECT_NAME)
