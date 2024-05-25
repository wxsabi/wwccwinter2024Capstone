#!/bin/bash

# Welcome message
echo "Running cacp.sh 'CHANGELOG - ADD - COMMIT - PUSH'"
sleep 1

# Check if remote repository exists
if ! git ls-remote --exit-code origin &> /dev/null; then
    echo "Remote repository does not exist or you do not have permission to access it"
    exit 1
else
    echo "Remote repository exists"
fi

# Check if nano is installed
if ! command -v nano &> /dev/null; then
    echo "nano could not be found, please install it first"
    exit 1
else
    echo "nano has been found in this system"
fi

# Add all changes to staging
git add .

# Check for changes
if git status --porcelain | grep -q '^[MADRC\?\?]'; then
    echo "Found changes to the local repository"
else
    echo "No changes to commit"
    exit 0
fi

# Create a template file
echo -e "[[]:\n\nTime spent:\n\n-" > commit_msg.txt

# Write commit message to a temporary file
nano commit_msg.txt

# Read from the temporary file and print it out
commit_message=$(cat commit_msg.txt)
echo "Commit message: $commit_message"

# Remove the original temporary file
rm commit_msg.txt

# Commit changes with the new message
git commit -m "$commit_message"