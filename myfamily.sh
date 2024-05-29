#! /bin/bash

relatives=$(curl -s https://platform.zone01.gr/assets/superhero/all.json | jq '.[] | select(.id=='$HERO_ID') | .connections.relatives')
relatives=$(echo "$relatives" | sed 's/"//g')
echo "$relatives"