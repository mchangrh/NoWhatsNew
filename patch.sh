#!/usr/bin/env bash

# Compatibility: 2024-03-12 Beta | 1710281934
# LICENCE: MIT

steampaths=("$1" "$1/steamui/css" "$HOME/.local/share/Steam/steamui/css" "$HOME/.steam/steam/steamui/css" "$HOME/Steam/steamui/css")
# ~/.local/share/Steam - Arch
# ~/.steam/steam - Debian
# ~/Steam - ???

# see https://github.com/ValveSoftware/steam-for-linux/issues/6976

function patch() {
  echo "$FILENAME"
  # add selector to hide updates section
  patch="}div:has(>._17uEBe5Ri8TMsnfELvs8-N){display:none;}"
  # find the libraryhome_UpdatesContainer_ reference
  line=$(grep -Po '\._17uEBe5Ri8TMsnfELvs8-N{.+?;(background-image:.+?})' "$FILENAME")
  # find the padding value
  padding=$(echo "$line" | grep -Po 'background-image.\+\?;')
  # find the difference in length between the padding and patch
  lendiff=$((${#padding}-${#patch}))
  # add padding spaces to the patch
  patch=$(printf "${patch}%*s" "$lendiff" "")
  # create modified reference
  line_edited=${line//$padding/$patch}
  # replace entire line in file
  sed -i "s/${line}/${line_edited}/g" "$FILENAME"
  echo "patched $FILENAME"
  exit 0
}

# search for file in given paths
for folder in "${steampaths[@]}"
do
  if [ -d "$folder" ]
  then
    FILENAME=$(grep -rl --include=*.css "\._17uEBe5Ri8TMsnfELvs8-N{.\+\?;background-image:.\+\?}" "$folder")
    if [ -n "$FILENAME" ]; then
      patch
    fi
  fi
done
echo "line not found - the file might be missing or changed"
exit 1