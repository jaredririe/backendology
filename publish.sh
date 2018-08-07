#!/bin/bash

echo -e "\033[0;32mPublishing updates to content/posts...\033[0m"

# Publish the updates
cd publish
go build && ./publish

# Come back up to the project's root.
cd ..
