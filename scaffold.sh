#!/usr/bin/env bash

sleep_time=1

case "$1" in
  "down")
    docker-compose down
    ;;
  "full-rebuild")
    docker-compose down
    docker-compose build
    case "$2" in
      "-prune")
        docker volume prune -f
        sleep_time=30
        ;;
    esac
    docker-compose up -d
    echo "waiting $sleep_time second(s) til database is ready.."
    sleep $sleep_time
    docker-compose restart swapiapp
    echo 'SWAPIAPP is Ready!'
    ;;
  "up")
    "$0" down
    docker-compose up -d
    ;;
  "logs")
    case "$2" in
      "-app")
        docker logs -f swapiapp
        ;;
      "-db")
        docker logs -f db
        ;;
    esac
    ;;
  *)
    echo "You have failed to specify what to do correctly."
    exit 1
    ;;
esac
