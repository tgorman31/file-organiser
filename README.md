# File-Organiser
A CLI program written in Go to help me manage the storage on my Laptop

This is my first Go program so hoping it helps me learn a few things along the way

Going to work my way up to the final program by implementing basic functionallity first. Will also use this as an opportunity to learn to use Git properly

## Expected final function
 - The user provides a directory path
 - The tool goes through the directory working out the size of all files and directories within that directory
 - The tool lists out each subdirectory and its size
 - The user can select a subdirectory and the tool will display the top 10 dir's or files in it by size

## Dev Stages
### 1. Basic Directory Listing
Create a Go program that can list all files and subdirectories in a given directory

**Goals**
- [ ] Parse the command line to accept a directory path
- [ ] List all items in the specified directory

### 2. Calculate Directory Size
Given a directory path calculate the total size of all the files in the directory

**Goals**
- [ ] Implement a function to calculate the size of a single directory
- [ ] Print the total size of the directory

### 3. Recursive Directory Traversal
Extend the functionality to calculate the size of directories recursively.

**Goals**
- [ ] Modify the directory size function to perform recursive calculations.
- [ ] Print the total size of the directory and each sub-directories

### 4. Sorting and Ranking Directories
After calculating sizes, sort the directories by size to easily idetify the largest ones

**Goals**
- [ ] Sort directories based on their size
- [ ] Print directories in order from largest to smallest

### 5. Identifying Top Items
Find and list the largest items within each directory. Includes both files and subdirectories

**Goals**
- [ ] Create a function to find the top N largest files in a directory

### 6. User Interface and Interactivity
Improve usability for the CLI tool by adding help commands, interactive prompts, flag controls etc.

**Goals**
- [ ] Add command-line flags for different functionalities
- [ ] Add help documentation

### 7. Performance Optimisation and Testing
Refine the performance and test reliability

**Goals**
- [ ] Add error handling to manage permissions issues or unreadable files
- [ ] Optimise performancce for large directories
- [ ] Write unit tests