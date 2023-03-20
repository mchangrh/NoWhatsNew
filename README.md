# NoWhatsNew

currently (2023-03-20), the filename is `chunk~547eb3232.css` under `C:\Program Files (x86)\Steam\steamui\css`  
Steam only checks the length of the file, so we replace the `padding:` property with `display: none` and pad it with the appropiate amount of spaces.

## Linux/ WSL
1. cd to the directory where `chunk~547eb3232.css` is located or `Steam\steamui\css` and run `bash patch.sh`  
2. After updates, just run the script again?

## Windows
1. Navigate to `C:\Program Files (x86)\Steam\steamui\css`
2. Find the top string and replace it with the bottom - it is essential to keep the extra spaces in.

This might not work between patches, a more permanent solution is coming soon TM
```
.libraryhome_UpdatesContainer_17uEB{box-sizing:border-box;padding:16px 24px 0px 24px;position:relative;height:324px;
.libraryhome_UpdatesContainer_17uEB{box-sizing:border-box;display: none;             position:relative;height:324px;
```