# Task-Tracker
Here's a description for your Task Tracker CLI project. You can expand or modify this to suit your style or focus areas:

# Task Tracker CLI

**Task Tracker CLI** is a simple, command-line-based application for managing and organizing tasks efficiently. Designed to help users keep track of what needs to be done, what’s in progress, and what has already been completed, this CLI tool provides a straightforward approach to task management directly from the terminal.

### Key Features

- **Add Tasks**: Quickly add tasks with a simple description.
- **Update Task Descriptions**: Modify existing tasks to keep information up-to-date.
- **Delete Tasks**: Remove completed or unnecessary tasks.
- **Status Management**: Mark tasks as "in-progress," "done," or reset to "todo."
- **Task Listing**: List all tasks or filter by status (e.g., show only completed tasks).
- **Persistent Storage**: Tasks are stored in a JSON file, preserving them between sessions.

### How It Works

The Task Tracker CLI stores all tasks in a JSON file within the current directory, allowing users to view, add, modify, or delete tasks at any time. Each task entry includes:

- **ID**: A unique identifier for the task.
- **Description**: A brief summary of the task.
- **Status**: Current task status (either "todo," "in-progress," or "done").
- **Created Date**: Timestamp when the task was initially added.
- **Updated Date**: Timestamp showing the last time the task was modified.

### Why Use Task Tracker CLI?

This CLI is a great tool for anyone looking to boost productivity and organize their workload without relying on complex project management tools. It’s also ideal for command-line enthusiasts who prefer managing tasks directly in the terminal.

### Sample Commands

Here’s how to get started with Task Tracker CLI:

```bash
# Adding a new task
task-cli add "Finish project documentation"

# Listing all tasks
task-cli list

# Updating an existing task
task-cli update 1 "Finish and review project documentation"

# Marking a task as in progress
task-cli mark-in-progress 1

# Marking a task as done
task-cli mark-done 1

# Deleting a task
task-cli delete 1
```

### Who is it for?

Task Tracker CLI is perfect for:
- Developers who want to manage to-do lists within the command line.
- Anyone who prefers a lightweight, terminal-based task tracker.
- Users who need a simple yet effective way to organize tasks.

Whether you're tracking personal tasks, managing a small project, or organizing work-related goals, Task Tracker CLI makes it easy to stay on top of what’s important.
