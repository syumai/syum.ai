PROJECT_NAME="syum-ai"

.PHONY: dev
dev:
	dev_appserver.py app/app.yaml

.PHONY: deploy
deploy:
	cd app && gcloud app deploy --project $(PROJECT_NAME)
