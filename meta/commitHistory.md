# Commit History
## Commit 1: Passing Config Lesson
This commit was the first state of the program which passes the tests for the Config lesson of the project guidance. It doesn't actually meet all the demands of the lesson though, so I'll make another commit related to the config lesson before moving on. 

I'm honestly not entirely clear why the lesson wants us to start this way though. I guess it helps flex the muscles of reading files in the first place, but there's really not a clear reason for a config file in the first place yet. If we're supposed to be learning good programming workflows (which admittedly isn't listed in the requirements so far) then why do these guided projects keep making us start these projects without a clue why we'd be doing what we're doing? 

Anyway, 
### Changes
- Initiated Go Module
- Added internal/config/config.go with a Read function and functionality to return a struct with the information from a `.gatorconfig.json` file held in the home directory for some reason.
- Added main.go with functionality to call the Read function and print the returned struct.
- Added temporary debug code demonstrating that each variable worked and did not remove yet.
- Added `.gitignore` to prevent `.swp` files from being committed.
- Successfully built go module.
- Added `README.md` with only a temporary introduction
- Added this `meta/commitHistory.md` for documentation
- Added `meta/lessonPages/` to record the project instructions for the lessons I've focused on as they were when I was working on this commit. These are titled with their dates in case of changes to the pages.
