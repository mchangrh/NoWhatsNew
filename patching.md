# Automatic patching with JS searching
1. Find child container since parent is anonymous and pull out `classname`
2. Recursive grep for `classname` in `*.js`
3. Find the corresponding js variable name - `UpdatesContainer` that matches `classname`
4. Search for replaceable references for class. We can target the `.BasicUI .classname` selector
5. Construct parent selector to match it

```diff
- .BasicUI ._17uEBe5Ri8TMsnfELvs8-N{background:none}⏎
+ div:has(>._17uEBe5Ri8TMsnfELvs8-N{display:none})  ⏎
```

# Manually patching anonymous container (react)
## Setup
1. Find the container that holds the updates box. This can be done by launching steam with the `-dev` flag
2. Find a identifying class of the child container we can target.
   1. In this case we used `._17uEBe5Ri8TMsnfELvs8-N`
3. construct a **chained** selector targeting the parent of this class with the `display:none` property
   1. In this case we use `}div:has(>._17uEBe5Ri8TMsnfELvs8-N){display:none;}`
   2. We have to chain it with the previous container, so the `}` is very important as to not break the CSS
4. Find a target property in the child block, ideally it is longer than our target string and also at the end.
   1. In this case we use `background-image:linear-gradient(to top, #171d25 0%, #2d333c 80%)}` since it is plenty long
5. Steam only check the length of the file so after we replace the last property, we pad our new selector with whitespaces
```diff
height:324px;overflow:hidden;
- background-image:linear-gradient(to top, #171d25 0%, #2d333c 80%)}⏎
+ }div:has(>._17uEBe5Ri8TMsnfELvs8-N){display:none;}                ⏎
```

# Manually patching named container (pre-react)
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

# $STEAMROOT locations
- Windows: `C:\Program Files (x86)\Steam`
- Arch: `~/.local/share/Steam`
- Debian/ Ubuntu: `~/.steam/steam`