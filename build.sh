#!/bin/bash

LDFLAGS='-linkmode external -extldflags "-static"' TAGS="bindata sqlite sqlite_unlock_notify" make build
