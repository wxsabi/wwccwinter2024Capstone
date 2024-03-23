#!/bin/bash

# This is a shell script written by Jay Cisneros<jaycisneros@jacm.io>
# This script is to be run within a git initialized project folder
# it will add all the changes is the folder, and open nano to
# allow a multiline message commit with a very small template. It
# will then update the local changelog, push to the remote repository,
# delete any temp messages and finally update the remote changelog.

#Welcome message
echo "Running cacp.sh 'CHANGELOG - ADD - COMMIT - PUSH'"
sleep 1s

# Check if remote repository exists
if ! git ls-remote --exit-code origin &> /dev/null
then
    echo "Remote repository does not exist or you do not have permission to access it"
    exit 1
else
    echo "Remote repository exists"
fi

# Check if nano is installed
if ! command -v nano &> /dev/null
then
    echo "nano could not be found, please install it first"
    exit 1
else
    echo "nano has been found in this system"
fi

# Check for changes
if git diff-index --quiet HEAD --
then
    echo "No changes to commit"
    exit 0
else
    echo "Found changes to the local repository"
fi

# Add all changes to staging
git add .

# Create a template file
echo -e "Time spent:\n\n[]:" > commit_msg.txt

# Write commit message to a temporary file
nano commit_msg.txt

# Commit changes
git commit -F commit_msg.txt

# Save current git log to a temporary file
string='```'
git log -1 --pretty=format:"<hr>%n%n### %ad%n#### Author: %an <%ae>%n#### commit: \`%H\`%n%n$string%n%B%n$string%n" > latest_commit.txt

# Create a temporary file
touch temp.txt

# If CHANGELOG.md exists, concatenate its contents to latest_commit.txt
if test -f CHANGELOG.md
then
    cat latest_commit.txt CHANGELOG.md > temp.txt
else
    cp latest_commit.txt temp.txt
fi

# Rename temp.txt to CHANGELOG.md
mv temp.txt CHANGELOG.md

# Remove the original latest_commit.txt
rm latest_commit.txt

# Push changes to main branch
git push origin main

# Delete the temporaty commit message file
echo "Deleting temporary files..."
rm commit_msg.txt

# Update remote changelog
echo "Update remote repository changelog file"
git add CHANGELOG.md
git commit -m "[AUTOMATED]Update CHANGELOG.md"
git push origin main

# Print success message
echo "SUCCESS!"
