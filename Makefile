up:
	@echo "Re-deploy container"
	podman-compose down
	podman-compose up -d
