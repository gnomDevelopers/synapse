#!/bin/bash

if [ -f /opt/synapse/build.lock ]; then
  exit 0
fi

touch /opt/synapse/build.lock

docker compose up -d --build --remove-orphans

rm /opt/synapse/build.lock
