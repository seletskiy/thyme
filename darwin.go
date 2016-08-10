package thyme

import "fmt"

type DarwinTracker struct{}

var _ Tracker = (*DarwinTracker)(nil)

// scriptableAppWindows fetches the windows of all scriptable applications.  It
// iterates through each application process known to System Events and attempts
// to script the application with the same name as the application process. If
// such an application exists and is scriptable, it prints the name of every
// window in the application. Otherwise, it just prints the name of every
// visible window in the application. If no visible windows exist, it will just
// print the application name.  (System Events processes only have windows in
// the current desktop/workspace.)
var scriptableAppWindows = `
tell application "System Events"
  set listOfProcesses to (every application process where background only is false)
end tell

repeat with proc in listOfProcesses
  set procName to (name of proc)
  log "(PROCESS) "& procName
  try
    tell application procName
      repeat with i from 1 to (count windows)
        log "(WINDOW) " & (name of window i) as string
      end repeat
    end tell
  on error e
    -- Fall back to listing all visible windows
    repeat with win in windows of proc
      log "(WINDOW) " & (name of win as string)
    end repeat
  end try
end repeat
`

// visibleWindows fetches all windows in the current desktop

func (t *DarwinTracker) Deps() string {
	// TODO: mention requirement of enabling privileged Accessibility tools in preferences
	return "TODO"
}

func (t *DarwinTracker) Snap() (*Snapshot, error) {
	return nil, fmt.Errorf("TODO")
}
