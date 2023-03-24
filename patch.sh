#!/bin/bash

FILENAME="./chunk~547eb3232.css"
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