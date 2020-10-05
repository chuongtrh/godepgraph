#!/bin/bash

heroku container:push web --app=polar-oasis-34616

heroku container:release web --app=polar-oasis-34616