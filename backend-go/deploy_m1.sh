#!/bin/bash
export DOCKER_DEFAULT_PLATFORM=linux/arm64

heroku container:login

echo "Build image"
heroku container:push web --app=polar-oasis-34616

echo "Release new version"
heroku container:release web --app=polar-oasis-34616