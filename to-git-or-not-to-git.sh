curl -s https://platform.zone01.gr/assets/superhero/all.json | jq -r '.[] | select(.id == 170) | (.name), (.powerstats.power), (.appearance.gender)'
