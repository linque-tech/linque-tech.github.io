#!/usr/bin/env bash
# Usage: ./deploy.sh user@your-server.com
# Builds the binary for Linux/amd64 and copies it to the server.

set -euo pipefail

TARGET="${1:?Usage: ./deploy.sh user@your-server.com}"
REMOTE_DIR="/opt/linque-contact"

echo "→ Building linux/amd64 binary..."
GOOS=linux GOARCH=amd64 go build -o contact-handler .

echo "→ Copying binary to ${TARGET}:${REMOTE_DIR}/"
ssh "$TARGET" "mkdir -p ${REMOTE_DIR}"
scp contact-handler "$TARGET:${REMOTE_DIR}/contact-handler"
ssh "$TARGET" "chmod +x ${REMOTE_DIR}/contact-handler && systemctl restart contact-handler && systemctl is-active contact-handler"

echo "✓ Deployed and service restarted."
