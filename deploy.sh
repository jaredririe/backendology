#!/bin/bash

echo -e "\033[0;32mDeploying updates to GitHub...\033[0m"

# Stash uncomitted and untracked changes
git stash --all

# Remove the contents of the current /public folder
rm -rf ./public/*

# Generate the static site in the default /public folder.
hugo

# Add changes to the git submodule.
cd public
git checkout README.md
git add .

# Commit changes.
msg="Rebuilding site on `date`"
if [ $# -eq 1 ]
  then msg="$1"
fi
git commit -m "$msg"

# Push changes to origin.
git push origin master

# Come back up to the project's root.
cd ..

# Unstash
git stash pop
