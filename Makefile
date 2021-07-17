PROJECT_NAME="syum-ai"
GOOGLE_ACCOUNT_NAME="syumai@gmail.com"
PORT?=8080

.PHONY: dev
dev:
	dev_appserver.py app.yaml --port=$(PORT)

.PHONY: deploy
deploy:
	GO111MODULE=on gcloud app deploy --project $(PROJECT_NAME) --account $(GOOGLE_ACCOUNT_NAME)
