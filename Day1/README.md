# Day 1

## Please provide your first day feedback
```
https://survey.zohopublic.com/zs/xafPLk
```

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

## Info - Linux Repository Server
<pre>
- it could be JFrog Artifactory or Sonatype Nexus or some FTP Servers
- the repository servers maintains all the opensource compatible/tested softwares for each version of linux distributions
- For instance, ubuntu has its own Repository server to host all the software packages that are supported by Ubuntu OS
</pre>  

## Info - Why Ubuntu 16.04 installs older version of Python, while Ubuntu 24.04 installs latest version of Python
<pre>
- the repository server url that comes by default with Ubuntu 16.04 points to python v3.5.2
- the repository server url that comes by default with Ubuntu 24.04 points to python v3.6.x 
</pre>  

## Info - Linux Package Manager
<pre>
- it is a utility that helps installing/uninstalling/upgrading/downgrading softwares in Linux OS
- package managers depends on repository servers to download and install softwares
- the repository server urls are maintained in restricted folders like /etc/apt in case of Ubuntu
- each Linux family supports a particular type of package manager
- For instance,
  - Ubuntu - uses apt(apt-get) package manager
  - Fedora - uses rpm/yum/dnf package manager
  - RHEL - uses rpm/yum/dnf package manager
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
cd Day1/CustomDockerAnsibleNodeImages/ubuntu-ansible
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
Delete any container with name ubuntu1 and ubuntu2 in case they exists
```
docker rm -f ubuntu1 ubuntu2
```

Create two ansible node containers
```
docker run -d --name ubuntu1 --hostname ubuntu1 -p 2001:22 -p 8001:80 tektutor/ubuntu-ansible-node:latest
docker run -d --name ubuntu2 --hostname ubuntu2 -p 2002:22 -p 8002:80 tektutor/ubuntu-ansible-node:latest
```

List the running containers
```
docker ps
```
Expected output
![image](https://github.com/user-attachments/assets/481f48a9-dd23-4192-a4a1-f5b5c0fea001)

Let's verify if we are able to SSH into the ubuntu1 ansible node container without password
```
ssh -p 2001 root@localhost
```

Expected output
![image](https://github.com/user-attachments/assets/481f48a9-dd23-4192-a4a1-f5b5c0fea001)

Let's verify if we are able to SSH into the ubuntu2 ansible node container without password
```
ssh -p 2002 root@localhost
```

Expected output
![image](https://github.com/user-attachments/assets/7e33de38-ff90-452f-877f-298c59af5508)


## Lab - Rebuilding Custom Docker Ansible node images

First of all you need to delete ubuntu1 and ubuntu2 containers
<pre>
docker rm -f ubuntu1 ubuntu2  
</pre>

Update the known hosts
```
echo "" > ~/.ssh/known_hosts
```

Next, you need to rebuild the docker image after git pull
```
cd ~/terraform-feb-2025
git pull
cd Day1/CustomDockerAnsibleNodeImages/ubuntu-ansible
docker build -t tektutor/ubuntu-ansible-node:latest .
docker images
```

Recreate, the ubuntu1 and ubuntu2 containers
```
docker run -d --name ubuntu1 --hostname ubuntu1 -p 2001:22 -p 8001:80 tektutor/ubuntu-ansible-node:latest
docker run -d --name ubuntu2 --hostname ubuntu2 -p 2002:22 -p 8002:80 tektutor/ubuntu-ansible-node:latest
```

List the running containers
```
docker ps
```

See if ansible is able to communicate with the newly created ansible node containers
```
cd ~/terraform-feb-2025
git pull
cd Day1/ansible
ansible -i inventory all -m ping
```

Expected output
![image](https://github.com/user-attachments/assets/32e2394a-fa8b-4778-8e6b-37879762ac71)

## Lab - Ansible Ad-hoc command to collect facts
Ansible facts are ansible node(remote server) meta-data

We can collects facts about the remote machines using setup module
```
ansible -i hosts ubuntu1 -m setup
```

Expected output
![image](https://github.com/user-attachments/assets/a11ea182-10bb-4fd3-bf1d-32b4cc5311db)
![image](https://github.com/user-attachments/assets/54424851-e492-426a-ae4c-9e9959c451aa)
![image](https://github.com/user-attachments/assets/250786f9-7d51-40be-9450-526aa1e25fe5)

## Lab - Copying a file from docker container to local machine
```
docker cp ubuntu1:/etc/nginx/sites-available/default .
```

## Lab - Running your first ansible playbook
```
cd ~/terraform-feb-2025
git pull
cd Day1/ansible
ansible-playbook -i hosts ping-playbook.yml
```

Expected output
![image](https://github.com/user-attachments/assets/d418342e-3c98-4e7f-af07-a9dfefbd9828)

## Lab - Installing nginx web server via ansible playbook
```
cd ~/terraform-feb-2025
git pull
cd Day1/ansible
ansible-playbook -i hosts install-nginx-playbook.yml
```

Expected output
![image](https://github.com/user-attachments/assets/779ec180-1653-4414-9db7-b01d4a2ed7a6)
![image](https://github.com/user-attachments/assets/9e527991-d465-40ff-adbe-8777bdf12ac5)
![image](https://github.com/user-attachments/assets/0818636e-50a1-4b43-9c10-5145776dc6da)
![image](https://github.com/user-attachments/assets/4b8370ad-d7ec-4f11-bf3d-8c6a77b733f8)
![image](https://github.com/user-attachments/assets/ba0dba20-2a20-45b9-afc2-37d2ed666ad4)
![image](https://github.com/user-attachments/assets/35051b75-3c54-4444-9a19-44103e7626d7)
![image](https://github.com/user-attachments/assets/5e734eb3-a1b2-4b50-9bd5-f47a318ccf74)
![image](https://github.com/user-attachments/assets/1dade90e-a400-469a-b993-8318dc9560e9)

## Lab - Getting help for ansible module
```
ansible-doc service
```

Expected output
![image](https://github.com/user-attachments/assets/85635179-1fc9-4ef2-925f-3ada6bb0f6af)

## Lab - Creating a custom ansible role to install nginx, configure nginx and deploy custom web page
```
ansible-galaxy init nginx
```

You could pull the role from github if you don't prefer doing it yourself
```
cd ~/terraform-feb-2025
git pull
cd Day1/ansible/roles
tree
```


Expected output
![image](https://github.com/user-attachments/assets/7998b670-dd9e-41b6-bd69-7513301ddd34)
