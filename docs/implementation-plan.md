# Implementation Plan

Implementation plan for **chronarch**, a time tracking/management app. Initial development will focus on core and CLI components before moving into any sort of TUI or GUI.

## First Slice

The ability to start and stop a time tracking session that is stored and categorized with under a project. This will also necessitate the ability to add projects.

### Example Sequence

1. Create a project
2. Start a time tracking session under the created project
3. Stop the session
4. View session length, start/stop times, and project associated with the session

### Necessary Systems

- CLI Core
- Projects
- Tracking Sessions
- Data Viewing
