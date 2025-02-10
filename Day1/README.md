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
## Info - Puppet/Chef Overview
<pre>
- configuration management tools
- they are difficult to setup
- steep learning curve
- The DSL(Domain Specific Language - language used to automate) used by Puppet is Puppet's proprietary language
- The DSL(Domain Specific Language - language used to automate) used by Chef is Ruby scripting language
- they both follow client/server architecture
- they both have proprietary agents that must be installed on the servers that needs to managed by Puppet/Chef
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
- is agentless ( doesn't follow client/server achitectures )
- it is a kind of stand-alone tool
- follows PUSH architecture
</pre>

## Puppet Overview
<pre>
- is the oldest configuration management tool
- it follows client/server architecture
- every configuration management supports a specific domain specific language (DSL) to automate stuffs 
- the DSL used by Puppet is Puppet language ( a proprietary declarative language )
- Puppet installation is very complex and time consuming
- learning curve is quite steep
- uses proprietary tools on the servers that needs to be managed by Puppet
- Puppet architecture is very complex
</pre>


## Chef Overview
<pre>
- is a configuration management tool
- it follows client/server architecture
- the domain specific language (DSL) - the automation language used by Chef
- the DSL used by Chef is Ruby ( scripting language )
- Chef installation is very complex and time consuming as Puppet
- Chef provides loads of tools, hence its very powerful and confusing
- learning curve is quite steep
- uses proprietary tools on the servers that needs to be managed by chef
- chef architecture is very complex
</pre>

## Ansible Overview
<pre>
- is a configuration management tool
- it is agent-less
- easy to learn
- easy to install
- follows simple architecture
- ansible nodes
  - these are servers we can perform automation using ansible 
  - dependent softwares
    - Unix/Linux/Mac Server
      - Python
      - SSH Server
    - Windows Server
      - Powershell
      - WinRM
- Ansible Controller Machine
  - the machine where ansible is installed is called Ansible Controller Machine(ACM)
  - it could a laptop/desktop
  - officially Ansible is only supported on Linux machines, but it works in Unix/Mac
  - Windows machine can't be used as a Ansible Controller Machine
  - Windows machine can be managed by Ansible
- Inventory
  - is a plain text file which follows an INI style format
  - captures connectivity details, IP address/hostname, username, password, ssh-key's etc
- comes in 3 flavours
  1. Ansible core - open source variant supports only command line
  2. AWX - supports webconsole, opensource, built on top of Ansible core
  3. Red Hat Ansible Tower - enterprise commercial product, built on top of AWX 
</pre>

## Ansible Core
<pre>
- this is developed in Python by Michael Deehan
- Michael Deehan is a former employee of Red Hat
- Michael Deehan founded a company called Ansible Inc and developed Ansible core as an open source product
- perfect alternate to Puppet/Chef
- supports only command line
- very well documented open source product
- agent less
- can be installed in Linux, Unix and Mac
- can manage Windows, Linux, Mac, etc., ansible nodes
- doesn't support role based access ( can't create different types of ansible users )
- doesn't historial logging mechanism
</pre>

## AWX
<pre>
- is developed on top of Ansible core
- supports webconsole but no command line
- it can be installed on a centralized server within your organization
- can be accessed from web browser only
- supports role based access control
- supports logs for each playbook execution
- you don't get any support
- can't develop ansible playbook, you can only run them
- which means we need ansible core to develop/write playbook
</pre>

## Red Hat Ansible Tower
<pre>
- it is developed on top of AWX
- functionally AWX and Ansible Tower(Ansible Automation Platform) are same
- you will world-wide support from Red Hat (an IBM company)
- which means we need ansible core to develop/write playbook
</pre>

## Ansible Modules
<pre>
- ansible supports many built-in ansible modules to automate
- for instance 
  - file module helps in creating files and folders with specific permissions
  - copy module helps in copying from/to ACM to ansible nodes and vice versa
  - all unix/linux/mac ansible modules are developed as Python scripts
  - all windows ansible modules are developed as Powershell scripts
  - we can also write out own custom ansible modules, when there is no readily available module to automate certain rare stuffs
</pre>

## Ansible Plugins
<pre>
- ansible plugins helps us extend the core functionality of ansible
- for instance
  - become plugin helps us perform certain tasks as sudo(administrative) users
</pre>

## Ansible Roles
<pre>
- is way we could follow best practices and ensure our automation code can be reused across many ansible playbooks
- ansible roles can't be executed directly, while they can be invoked via ansible playbooks
- ansible roles can be downloaded and installed via ansible-galaxy tool
- we could also develop our own ansible role
- For example
  - we could develop an ansible role to install Oracle Database in Windows 2016/2019 Server, Ubuntu Linux, etc
</pre>

## Ansible Playbook
<pre>
- is a YAML file 
- it invokes bunch of Ansible module, roles in a particular order to automate configuration management activity
</pre>

## Ansible Collections
<pre>
- is a reusable code that has many different kinds of reusable code in ansible
- it could have one or more roles, custom modules, plugins, filters, etc.,
- it's a way we could package and distribute all the related playbooks, modules, plugins, etc in a single collection
</pre>

## Lab - Cloning TekTutor Training Repository ( one time activity )
```
git clone https://github.com/tektutor/terraform-feb-2025.git
cd ~/terraform-feb-2025
```

## Lab - Let's create a custom ubuntu ansible node docker image
```
cd ~/terraform-feb-2025
git pull
cd Day1/CustomDockerAnsibleNodeImages
ssh-keygen
cp ~/.ssh/id_ed25519.pub authorized_keys
docker build -t tektutor/ubuntu-ansible-node:latest .
docker images
```

Expected output
![image](https://github.com/user-attachments/assets/acd7f6a1-7c44-480a-a325-e18d206932f8)
![image](https://github.com/user-attachments/assets/d8da70f2-542c-4ce0-9241-c8026e6f8377)
![image](https://github.com/user-attachments/assets/57f53d6f-68c7-4b86-8b9b-30f16d3dcfc8)

## Lab - Let's create 2 containers using our custom docker image
```
docker run -d --name ubuntu1 --hostname ubuntu1 tektutor/ubuntu-ansible-node:latest
docker run -d --name ubuntu2 --hostname ubuntu2 tektutor/ubuntu-ansible-node:latest
```

List the running containers
```

```



