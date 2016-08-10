#!/usr/bin/env osascript

tell application "System Events"
    set window_name to name of first window of (first application process whose frontmost is true)
end tell

log window_name
