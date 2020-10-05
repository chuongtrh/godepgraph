#!/bin/bash

echo "Build image"
heroku container:push web --app=polar-oasis-34616

echo "Release new version"
heroku container:release web --app=polar-oasis-34616