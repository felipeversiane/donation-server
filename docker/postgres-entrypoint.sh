#!/bin/bash
set -e

# Configure cron job using runtime env vars
printf "%s postgres BACKUP_PASSWORD=%s BACKUP_DIR=%s MAX_BACKUPS=%s BACKUP_LOG_FILE=%s /usr/local/bin/backup.sh >> %s 2>&1\n" \ 
  "${BACKUP_SCHEDULE}" "${BACKUP_PASSWORD}" "${BACKUP_DIR}" "${MAX_BACKUPS}" "${BACKUP_LOG_FILE}" "${BACKUP_LOG_FILE}" > /etc/cron.d/postgres-backup
chmod 0644 /etc/cron.d/postgres-backup

# Start cron service
service cron start

exec docker-entrypoint.sh "$@"

