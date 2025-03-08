##Write different operating system services

Answer:- An operating system is software that acts as an intermediary between the user and computer hardware 
.It is a program with the help of which we are able to run cariout applications . It is the open program that is running all the time .Every computer must have any operating system to smoothly execute other programs 

- The Os coordinates te use of the hardware and application programs for various users. It provides a platform for other application program to work .
- It controls Input - Output devices execution of programs managing files etc
- Servies of OS include
- 1) Program excution
  2) Input Output operations
  3) File mangement
  4) Memory Management
  5) User Interface
  6) Networking
  7) etc
 
  ##What is System Call ?
  Answer:- System call provied the services of operating system to the user programsn via Application Program Interface(API)
  -It provied an interface between a process and operating system to allow user level processes to request services of operating system
  -In Operating systems there are two main modes " Kernel Mode" and "User mode"
  -In Kernel mode Operating system works on it owns  . It handels the Hardware and performs its fucntion with precision
  -On the Other hand User mode is where the user interact with the system by giving various commands
  - When User issudes commands need help of operating system it calls kernel help with this "System call "
  - System calls acts as bridgde between the Users and Kernel 
  - The Operating system has a predefined set of system calls , For example Windows as around 700 system calls
  =>  How System calls work:

      Requesting Services----> Mode Switching----> The Call Execution -----> Return To User mode______________>Once the system call is completed the operating system returnd the program to user mode 
            |                                |                       | 
            v                                V                        ______________
        In this state                  By default your program run                  |        
    Users Request Serivces from        in User mode which restricts                 V 
    The OS and then A system call      accessing the hardware and sensitive     With the program in Kernel mode 
    is made from the user mode to      parts of the system .When the            the Os indetifies the specific system call 
    kernel of the os                   System call is called then the mode       requested and then executes the associated code 
                                        temporarily switches to kernel mode      example: Reading a file, making a file, Reading from the 
                                        to gain access of the hardware.  -       input devices
 
  
  
  
    
