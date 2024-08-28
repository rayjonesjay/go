#!/bin/bash

# Directory to search in (current directory by default)
search_dir="${PWD}"
# Text to search for
search_text="stty"

# Find all .go files in the directory and search for the text
find "$search_dir" -type f -name "*.go" -exec grep -Hn "$search_text" {} \;
