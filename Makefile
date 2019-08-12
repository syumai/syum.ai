PROJECT_NAME="syum-ai"

.PHONY: dev
dev:
	dev_appserver.py cmd/app.yaml

.PHONY: deploy
deploy:
	cd cmd && gcloud app deploy --project $(PROJECT_NAME)
