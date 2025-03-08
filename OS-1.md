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

##  Monolithic Kerndel and Microkernel 
- There are 5 type of kernels
- A Microkernel which only contains basic fucntionality
- A Monolithic Kernel Which contains many device drivers
- Hybrid Kernel
- Exo-kernel
- Nano kernel

## Micro-Kernel
  The kernel manages the operations of the Computer . In microkernel The user servies and kernel servies are implemented in different address space and kernel Servies are kept under the kernel address space
- A microkernel is a minimalistic approach to designing an operating system , In a microkernel architecture only the most essential fucntion are included in the kernel such as basic communication between hardware and software and simple proces management other servies like device drivers, file systems and network protocols are run in user space as separate process

## Advantages of Micro-kernel 
- Size of Microkernel is smaller, so it is easy to use
- Easy to Extend Micro kernel
- Easy to Port Microkernel
- Less prone to errors because all the process works indepently 
- Example of such a Operatinf system is MacOS

## Disadvantages of Micro-kernel 
- Slower

## Monolithic-Kernel 

In the Monolithic Kernel the entire operating system runs as a single program in kernel mode . The user serives and kernel servies are implemeneted in the same address space 

- In this type of architecture Every core fucntions lie memory management , process management, Device drivesm and file system is intergrated into a single large clock of code running in a single address space
- This system can make the kernel because the process can interact with each other directly
- But this can also make the system more complex and harder to maintain
- And a Bug can bring down entire system
- 
## Advantages of Monolithic Kervel 
- Monolithic has all in one piece user servcies and kernel services under same address space
- Faster speed
- Example for such that Operatig system is Linux

## Disafvantages of Monolothic Kernel 
- Hard to port
- Large is Size
- More Prone to Errors


  ## Multiprocessing
  In a uni processor system only one proces executes ar a time . Multiprocessing makes use of two or mores CPU within a single computer system . The term also refers to the ability of a system to support more than one processor within a single computer system since ther are multiple procesors available , Multiple processes can be executed at a time . Since there are Multiple processess can be executed at a time these Multiprocess share the computer bus , memory etc

- With the help of multipleprocessing, Many process can be executed simultaneously For example, P1,P2,P3,P4 are waiting for execution . In a sinlge processor system all of this process get executed one by one .But in multipleprocessing, Each process can be assignes to a different process can be executed simultaneously and Thus two times faster and so on

- The main advantages of this multiprocessing is that it get more work done in short time
- THese type of system are used when high excution speeds and huge volume of data is required
-  We can say that Multiple processing occurs like  Parallel processing

## Note that Multiprocessing and MultipleProgramming are two different things 

## MultipleProgramming 
In a multi-programmed system, the CPU selects a function from all the waiting functions and starts performing it and when the process needs to wait for any IO operation then the CPU switches to another process and starts executing it. It is also known as the concept of context switching. Therefore, this keeps the CPU busy and active all the time and reduces CPU idle time. It can be done with slower CPU processing. It requires a small amount of memory (RAM or ROM) to operate.

## Objective
The multiprogramming objective is to allow more processes to run at the same time and keep the CPU busy for as long as possible and simply focus on maximizing CPU utilization or CPU time.

## Advantages of Multi Programming Operating System
-High CPU utilization.
-Shorter Response Time.
-Shorter Waiting Time.
-Increased Throughput.
-It may be extended to multiple users.
-It has the ability to assign priority to different processes.


## Define: Context Switching
 - Context Switching is the process of storing the state of a running process or thread so that it can be restored and resume executaion at a late point and then loading the context or state of another or thread and run  it
 - Context Switching is necessary for multitasking and efficient resourse management
 - Contest switching involes saving the following infomation
 - the contest of the CPU registers which store the current state of the process
 - THe process or thread memory map which maps the process or thread virtual memory address to its physical address
   The process fucntion call stack and other infomation which is needed to reume its execution

## Context Switching is OS
-IT Allows MultiTasking
- eFFICIENT RESOURCE MANAGEMENT
- 
 - 

