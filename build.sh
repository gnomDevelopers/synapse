touch /opt/synapse/build.lock
docker compose up -d --build --remova-orphans
rm /opt/synapse/build.lock
