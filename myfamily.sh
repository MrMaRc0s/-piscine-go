family=$(curl -s https://platform.zone01.gr/assets/superhero/all.json | jq '.[] | select(.id == $HERO_ID') | .connections.relatives')
family=$(echo "$family" | sed 's/"//g')
echo "$family"