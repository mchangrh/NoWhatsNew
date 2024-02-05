#!/usr/bin/env bash

# LICENCE: MIT

steampaths=("$1" "$1/steamui/css" "$HOME/.local/share/Steam/steamui/css" "$HOME/.steam/steam/steamui/css" "$HOME/Steam/steamui/css")
# ~/.local/share/Steam - Arch
# ~/.steam/steam - Debian
# ~/Steam - ???

# see https://github.com/ValveSoftware/steam-for-linux/issues/6976

# search for file in given paths
for folder in "${steampaths[@]}"
do
  if [ -d "$folder" ]
  then
    FILENAME=$(grep -rl --include=*.css "libraryhome_UpdatesContainer" "$folder")
    if $FILENAME; then break; fi
  fi
done
# set display to none
patch="display: none;"
# find the libraryhome_UpdatesContainer_ reference
line=$(grep -Po 'libraryhome_UpdatesContainer_[0-9a-zA-Z]+?{[^{]+?padding.+?}' "$FILENAME")
# find the padding value
padding=$(echo "$line" | grep -Po 'padding.+?;')
# find the difference in length between the padding and patch
lendiff=$((${#padding}-${#patch}))
# add padding spaces to the patch
patch=$(printf "${patch}%*s" "$lendiff" "")
# create modified reference
line_edited=${line//$padding/$patch}
# replace entire line in file
sed -i "s/${line}/${line_edited}/g" "$FILENAME"
