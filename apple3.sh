#!/usr/bin/env osascript

tell application "System Events"
     set listOfProcesses to (every application process where background only is false)
end tell

repeat with proc in listOfProcesses
     set procName to (name of proc)
     log procName
     try
     tell application procName
     	  repeat with i from 1 to (count windows)
	  	 log (name of window i) as string
	  end repeat
     end tell
     on error e
       -- Note: these only show up if they're visible
       log "## ERROR: " & e
       repeat with win in windows of proc
         log "# a window: " & (name of win as string)
       end repeat
     end try
end repeat
