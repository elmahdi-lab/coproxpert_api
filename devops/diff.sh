#!/bin/bash
atlas migrate hash --env gorm
atlas migrate diff --env gorm
