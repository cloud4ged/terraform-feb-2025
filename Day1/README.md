# Day 1

## Info - Configuration Management Tool Overview
<pre>
- assuming we already have a machine with Windows/Linux/Mac pre-installed, we can automate configuration management on those machines
- helps automating system administrative tasks like software installation
- helps inmanaging linux/windows users giving specific permissions, etc.,  
- Examples
  - Puppet
  - Chef
  - Ansible
</pre>

## Info - Programming languages classification
<pre>
1. Imperative
   - they are very powerful languages
   - For example
     - if I need to automate copying a file from one server to another server using imperative language
       - I need to write code to install What I wish to automate ( What )
       - I also have to write code to copy the file from one machine to the other ( How )
   - examples
     - C++, Java, C#, Python, Shell scripts, Powershell, Perl, LUA, etc.,
2. Declarative
   - You just have to say What needs to be automated, the how part is taken care
   - Example
     - Ansible, Puppet, Chef are all declarative

</pre>  

## Info - Ansible Overview
<pre>
- is a configuration management tool that helps in automation system administration activities
  - installing/uninstalling/upgrading softwares
  - managing users
    - creating a new linux/windows user
    - modifying user credentials
    - providing access rights to different users
    - deleting users
  - can manages
    - Windows servers
    - Linux servers
    - Mac servers/machines
    - Cisco routers/switches
- ansible is idempotent
  - no matter how many times you run the automation script, the outcome will be consistent
- 
</pre>
