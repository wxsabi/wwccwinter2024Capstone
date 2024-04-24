#!/bin/bash

# This script will generate the latest changelog

# Save current git log to a temporary file
git log --pretty=format:"<hr>%n%n### %ad%n#### Author: %an <%ae>%n#### commit: \`%H\`%n%n%n%B%n%n" > latest_changelog.md
