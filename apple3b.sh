#!/usr/bin/env osascript

tell application "System Events"

set theProcesses to application processes
repeat with theProcess from 1 to count theProcesses
  tell process theProcess
     log "(process) " & theProcess as string
     repeat with x from 1 to (count windows)
       log "# process " & theProcess as string
       log (name of window x) as string
     end repeat
   end tell
end repeat

end tell
