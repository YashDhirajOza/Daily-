# Operating System Services

An **Operating System (OS)** is software that acts as an intermediary between the user and computer hardware. It enables the execution of applications and manages system resources. The OS runs continuously and ensures the smooth execution of other programs.

## Functions of an Operating System
- **Resource Coordination:** The OS manages hardware resources and application programs.
- **Platform for Applications:** It provides a stable environment for software execution.
- **Control Over I/O Devices:** Manages input and output operations efficiently.
- **Execution of Programs:** Loads and runs user applications.
- **File Management:** Organizes and secures system files.

## Services Provided by an OS
1. **Program Execution** - Loads programs into memory and executes them.
2. **Input/Output Operations** - Facilitates interaction with input and output devices.
3. **File Management** - Organizes, stores, retrieves, and secures files.
4. **Memory Management** - Allocates and deallocates memory as needed.
5. **User Interface** - Provides CLI (Command Line Interface) and GUI (Graphical User Interface).
6. **Networking** - Manages network connections and system communication.

---

# System Calls
A **System Call** provides OS services to user programs via an **Application Program Interface (API)**. It acts as an interface between a process and the operating system, allowing user-level processes to request OS services.

## Modes in an Operating System
An OS operates in two modes:
- **Kernel Mode:** The OS runs at a privileged level, managing hardware and system operations.
- **User Mode:** The user interacts with the system by executing programs with restricted access to system resources.

When a user issues a command that requires OS assistance, a **system call** is made to the kernel.
System calls serve as a **bridge between users and the kernel**.

For example, Windows has around **700 predefined system calls**.

## How System Calls Work
1. **Requesting Services** → The user requests a service from the OS.
2. **Mode Switching** → The system call is triggered, switching from user mode to kernel mode.
3. **Call Execution** → The OS processes the request and executes the appropriate system call.
4. **Return to User Mode** → Once execution is complete, control is returned to user mode.

### Example Workflow
- A user requests to **read a file**.
- The request triggers a **system call**, switching the mode to **kernel mode**.
- The OS processes the request and accesses the required file.
- The file data is retrieved and returned to the user, switching back to **user mode**.

## Common System Calls
System calls allow the OS to perform tasks like:
- **Process Control:** Creating and terminating processes.
- **File Management:** Reading, writing, and deleting files.
- **Device Management:** Requesting device access and releasing it.
- **Information Maintenance:** Retrieving system status and process attributes.
- **Communication:** Facilitating interprocess communication and networking.

---
This document provides an overview of **Operating System Services** and **System Calls**, explaining their functionality and importance in computing.

