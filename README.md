# File-Organiser
A CLI program written in Go to help me manage the storage on my Laptop

This is my first Go program so hoping it helps me learn a few things along the way

Going to work my way up to the final program by implementing basic functionallity first. Will also use this as an opportunity to learn to use Git properly

## Expected final function
 - The user provides a directory path
 - The tool goes through the directory working out the size of all files and directories within that directory
 - The tool lists out each subdirectory and its size
 - The user can select a subdirectory and the tool will display the top 10 dir's or files in it by size

## Result
 - Helped with identifing over 60gb of items to be cleared off my device

## What I've Learned
 - How to structure a Go project and how the files interact with each other
 - Some basic Go functionality incl structs, ranges & slices
 - Reading Directories and FileInfo in Go with the os stdlib
 - Recursive Function calling
 - A file navigation system would have worked better for my desired goals
 - Reading user input with bufio
 - Reading command line flags
 - Slice sorting
 - A basic understanding of Bubbletea

## Lessons Learned
 - Go feels quite natural to code in
 - There are a lot of concepts that I still need to work on
 - A file organiser should allow for better traversing of directories (v2?)
 - Definitely need to do more with Bubbletea, love the concept but definitely need to do more in it to understand it better
 - The Bubbletea component added more complexity where a file output may have sufficed
 - I have no idea on how to do testing
 - Not fully grasping some concepts led me down paths which required re-writes

## Dev Stages
### 1. Basic Directory Listing
Create a Go program that can list all files and subdirectories in a given directory

**Goals**
- [x] Parse the command line to accept a directory path
- [x] List all items in the specified directory

### 2. Calculate Directory Size
Given a directory path calculate the total size of all the files in the directory

**Goals**
- [x] Implement a function to calculate the size of a single directory
- [x] Print the total size of the directory

### 3. Recursive Directory Traversal
Extend the functionality to calculate the size of directories recursively.

**Goals**
- [x] Modify the directory size function to perform recursive calculations.
- [x] Print the total size of the directory and each sub-directories

### 4. Sorting and Ranking Directories
After calculating sizes, sort the directories by size to easily idetify the largest ones

**Goals**
- [x] Sort directories based on their size
- [x] Print directories in order from largest to smallest

### 5. Identifying Top Items
Find and list the largest items within each directory. Includes both files and subdirectories

**Goals**
- [x] Create a function to find the top N largest files in a directory

### 6. User Interface and Interactivity
Improve usability for the CLI tool by adding help commands, interactive prompts, flag controls etc.

**Goals**
- [x] Add command-line flags for different functionalities
- [x] Add help menu

### 7. Performance Optimisation and Testing
Refine the performance and test reliability

**Goals**
- [x] Add error handling to manage permissions issues or unreadable files
- [ ] Optimise performancce for large directories
- [ ] Write unit tests