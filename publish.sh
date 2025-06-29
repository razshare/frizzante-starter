#!/usr/bin/env bash
# shellcheck disable=SC2155
export app_version=$(< version)
git add .
git commit -m"chore(app): tagging version $app_version"
git tag "$app_version"
git push --tags
git push