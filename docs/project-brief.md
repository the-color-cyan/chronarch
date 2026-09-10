# chronarch Project Brief

## Elevator Pitch

**chronarch** is a time tracking application designed for tracking and categorizing time spent working on tasks. At minimum, it will support categorizing time under projects. Broader and narrower categorization will exist in the form of collections of projects and collections of tasks within projects, as well as things like tags at each level.

---

## Problem Statement

I'm currently using [TimeScribe](https://github.com/WINBIGFOX/timescribe) and it covers a lot of my use-cases. I like the focus on projects, but it lacks categorization outside of this such as:
- tasks to break up projects
- tags for projects, tasks, etc.
- ways to categorize projects (by job, organization, etc.)

Additionally with how often I am working out of the terminal, having something with first-class CLI support is essential. This also has pointed me towards developing a TUI component, as there are visual aspects to time tracking that are difficult to display in a readable fashion with typical terminal output.

---

## Existing Systems / Dependencies

I would like to at least have surface-level support for importing the time I currently have tracked in TimeScribe. Whether or not this is valuable as a lasting feature is yet to be determined.
