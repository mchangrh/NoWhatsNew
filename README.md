# NoWhatsNew

# Patching with the scripts
Run `patch.sh` on linux, `NoWhatsNew-v*.exe` from [Releases](https://github.com/mchangrh/NoWhatsNew/releases) on windows

# Manually patching
1. Search for the css file that includes the `WhatsNew` container
   1. Currently (2023-04-27) this file is named `sp.css` under `$STEAMROOT\steamui\css`
2. Locate the block that modifies the container - `libraryhome_UpdatesContainer`...
3. Find a property that is long enough to fit the modification and not important, in this case we choose `padding`
4. Steam only checks the length of the file so we replace `padding:` with `display:none` and pad it with the appropiate amount of spaces
```diff
- padding:16px 24px 0px 24px;⏎
+ display:none;              ⏎
```
5. Exit/ Quit out of Steam and restart the steam client.

## $STEAMROOT
- Windows: `C:\Program Files (x86)\Steam`
- Arch: `~/.local/share/Steam`
- Debian/ Ubuntu: `~/.steam/steam`