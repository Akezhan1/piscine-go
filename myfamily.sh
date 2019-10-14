#!/bin/bash
curl https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json | jq -r --arg HERO_ID $HERO_ID '.[] | select( .id == ($HERO_ID|tonumber)) | .connections.relatives'
