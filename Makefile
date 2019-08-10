PROJECT_NAME="syum-ai"

.PHONY: deploy
deploy:
	cd cmd && gcloud app deploy --project $(PROJECT_NAME)
