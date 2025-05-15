#!/bin/bash

BACKUP_DIR="/backups"
MAX_BACKUPS=7 
DATE=$(date +%Y%m%d_%H%M%S)
TEMP_FILE="$BACKUP_DIR/temp_$DATE.sql.gz"
BACKUP_FILE="$BACKUP_DIR/backup_$DATE.sql.gz.gpg"

if [ -z "$BACKUP_PASSWORD" ]; then
    echo "[$(date)] ERROR: Backup password not set (BACKUP_PASSWORD environment variable)"
    exit 1
fi

cleanup() {
    ls -t $BACKUP_DIR/backup_*.sql.gz.gpg 2>/dev/null | tail -n +$((MAX_BACKUPS + 1)) | xargs rm -f 2>/dev/null

    find $BACKUP_DIR -name "*.log" -mtime +30 -exec rm {} \; 2>/dev/null

    rm -f $BACKUP_DIR/temp_*.sql.gz
}

echo "[$(date)] Starting encrypted backup..."

if pg_dump -U "$POSTGRES_USER" "$POSTGRES_DB" | gzip > "$TEMP_FILE" && \
   gpg --batch --yes --passphrase "$BACKUP_PASSWORD" -c "$TEMP_FILE" && \
   mv "$TEMP_FILE.gpg" "$BACKUP_FILE"; then
    
    rm -f "$TEMP_FILE"
    
    echo "[$(date)] Backup completed successfully: $BACKUP_FILE"
    echo "[$(date)] Backup size: $(du -h "$BACKUP_FILE" | cut -f1)"
    
    cleanup
    
    echo "[$(date)] Available encrypted backups:"
    ls -lh $BACKUP_DIR/backup_*.sql.gz.gpg 2>/dev/null | awk '{print $9, $5}'
else
    echo "[$(date)] ERROR: Backup failed"
    rm -f "$TEMP_FILE" "$TEMP_FILE.gpg" "$BACKUP_FILE" 
    exit 1
fi 