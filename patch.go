package main

import (
	"errors"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"syscall"
	"unsafe"
)

func alert(usermsg string, err bool) {
	// https://learn.microsoft.com/en-us/windows/win32/api/winuser/nf-winuser-messageboxw
	var user32DLL = syscall.NewLazyDLL("user32.dll")
	var procMessageBox = user32DLL.NewProc("MessageBoxW")
	const (
		MB_OK = 0x00000000

		MB_ICONSTOP        = 0x00000010
		MB_ICONQUESTION    = 0x00000020
		MB_ICONWARNING     = 0x00000030
		MB_ICONINFORMATION = 0x00000040
	)

	lpCaption, _ := syscall.UTF16PtrFromString("NoWhatsNew")
	lpText, _ := syscall.UTF16PtrFromString(usermsg)
	icon := uintptr(MB_ICONINFORMATION)
	// replace with warning icon if error
	if err {
		icon = MB_ICONWARNING
	}

	syscall.SyscallN(procMessageBox.Addr(),
		0,
		uintptr(unsafe.Pointer(lpText)),
		uintptr(unsafe.Pointer(lpCaption)),
		MB_OK|icon,
	)
}

func glob(root string) []string {
	var files []string
	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err == nil && strings.HasPrefix(info.Name(), "chunk") && filepath.Ext(path) == ".css" {
			files = append(files, path)
		}
		return nil
	})
	return files
}

func main() {
	println("Report issues to https://github.com/mchangrh/NoWhatsNew/issues")
	files := glob("C:\\Program Files (x86)\\Steam\\steamui\\css")
	success := false
	var err = errors.New("no candidate files found for patching - Make sure Steam is installed at C:\\Program Files (x86)\\Steam\\")
	for _, file := range files {
		err = readPatchFile(file)
		if err.Error() == "already patched" {
			alert("Already patched", false)
			println("already patched")
			return
		}
		if err == nil {
			success = true
		}
	}
	if !success {
		alert(err.Error(), true)
		println("error: " + err.Error())
	}
}

func readPatchFile(file string) error {
	content, readFileErr := os.ReadFile(file)
	if readFileErr != nil {
		return errors.New("error reading file " + readFileErr.Error())
	}
	newfile, patchFileErr := patch(string(content))
	if patchFileErr != nil {
		return patchFileErr
	}
	if newfile != "" {
		println("Patching file")
		writeFileErr := os.WriteFile(file, []byte(newfile), 0644)
		if writeFileErr != nil {
			return errors.New("error writing file " + writeFileErr.Error())
		}
	}
	alert("Done - Patched File", false)
	println("Done - Patched File")
	return nil
}

func patch(file string) (result string, err error) {
	patch := "display: none;"
	// find line
	lineRegex := regexp.MustCompile(`libraryhome_UpdatesContainer_[0-9a-zA-Z]+?{[^{]+?padding.+?}`)
	line := lineRegex.FindString(file)
	if line == "" {
		// find patched line
		patchedRegex := regexp.MustCompile(`libraryhome_UpdatesContainer_[0-9a-zA-Z]+?{[^{]+?display: none;.+?}`)
		patched := patchedRegex.FindString(file)
		if patched != "" {
			return "", errors.New("already patched")
		}
		return "", errors.New("line not found - the file might be missing or changed")
	}
	// find padding
	cssPaddingRegex := regexp.MustCompile(`padding.+?;`)
	cssPadding := cssPaddingRegex.FindString(line)
	// find difference in length
	paddingLength := len(cssPadding) - len(patch)
	stringPadding := strings.Repeat(" ", paddingLength)
	patch = patch + stringPadding
	// replace in line for
	patch = strings.Replace(line, cssPadding, patch, 1)
	// replace in file
	file = strings.Replace(file, line, patch, 1)
	// Write file
	return file, nil
}
