#!/bin/bash

# TODO: adjust the url based on the env.
atlas migrate apply --env gorm --url "postgres://postgres:postgres@localhost:5432/coproxpert_db?sslmode=disable&search_path=public" --revisions-schema true
