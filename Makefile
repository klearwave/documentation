virtualenv:
	@python3 -m venv venv
	@. venv/bin/activate && \
		pip install --upgrade pip && \
		pip install -r requirements.txt

submodules:
	@scripts/submodules.sh

build:
	@. venv/bin/activate && \
		mkdocs build

run:
	@. venv/bin/activate && \
		mkdocs serve

release:
	@. venv/bin/activate && \
		mkdocs gh-deploy --force
