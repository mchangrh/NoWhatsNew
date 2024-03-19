#!/usr/bin/env bash

# Compatibility: 2024-03-12 Beta | 1710281934
# LICENCE: MIT

steampaths=("$1" "$1/steamui" "$HOME/.local/share/Steam/steamui" "$HOME/.steam/steam/steamui" "$HOME/Steam/steamui")
# ~/.local/share/Steam - Arch
# ~/.steam/steam - Debian
# ~/Steam - ???

# see https://github.com/ValveSoftware/steam-for-linux/issues/6976

function patch() {
  patch="div:has(>.${CLASSNAME}){display:none}"
  # check if patch is already applied
  if grep -q "$patch" "$FILENAME"; then
    echo "patch already applied"
    exit 0
  fi
  # find and patch line
  candidate=$(grep -oP "\.BasicUI \.${CLASSNAME}{.+?}" "$FILENAME")
  # find the difference in length between the padding and patch
  lendiff=$((${#candidate}-${#patch}))
  # add padding spaces to the patch
  paddedpatch=$(printf "${patch}%*s" "$lendiff" "")
  # replace entire line in file
  sed -i "s/${candidate}/${paddedpatch}/g" "$FILENAME"
  echo "patched $FILENAME"
  exit 0
}

# search for file in given paths
for folder in "${steampaths[@]}"; do
  if [ -d "$folder" ]; then
    # find cssname from JS
    CLASSNAME=$(grep -hoPr --include "*.js" 'UpdatesContainer:"(.+?)"' "$folder" | grep -oP '"(.+?)"' | cut -d '"' -f 2)
    FILENAME=$(grep -lPr --include=*.css "\.${CLASSNAME}" "$folder")
    if [ -n "$FILENAME" ]; then
      patch
    else
      echo "line not found - the file might be missing or changed - Please open an issue on https://github.com/mchangrh/NoWhatsNew/issues/new"
      exit 1
    fi
  fi
done
echo "Steam install not found"
exit 1
