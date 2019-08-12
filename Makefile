PROJECT_NAME="syum-ai"

.PHONY: dev
dev:
	dev_appserver.py app.yaml

.PHONY: deploy
deploy:
	gcloud app deploy --project $(PROJECT_NAME)
