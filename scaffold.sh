#!/usr/bin/env bash



case "$1" in
  "down")
    echo 'Running down!'
    docker-compose down
    ;;
  "build")
    echo 'Running build!'
    docker-compose build
    ;;
  "runtests")
    echo 'Running tests!'
    go test -v -coverpkg=./... -coverprofile=profile.cov ./...; go tool cover -func profile.cov
    ;;
  "full-rebuild")
    "$0" down
    "$0" build
    t=1
    shift
    if [[ "$1" != "" ]]; then
      if [[ "$1" == "prune" || "$1" == "-prune" ]]; then
        t=30
        "$0" "$1"
      else
        "$0" "$1"
      fi
    fi
    shift
    if [[ "$1" != "" ]]; then
      if [[ "$1" == "prune" || "$1" == "-prune" ]]; then
        t=30
        "$0" "$1"
      else 
        "$0" "$1"
      fi
    fi
    
    echo 'Running Up'
    docker-compose up -d
    echo "waiting $t second(s) til database is ready.."
    sleep $t
    echo 'Running restart swapiapp!'
    docker-compose restart swapiapp
    echo 'SWAPIAPP is Ready!'
    ;;
  "-prune")
    echo 'Running Prune!'
    docker volume prune -f
    ;;
  "prune")
    echo 'Running prune!'
    docker volume prune -f
    ;;
  "-runtests")
    "$0" runtests
    ;;
  "up")
    echo 'Running up!'
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
    exit 1
    ;;
esac
