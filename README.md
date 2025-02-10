# Terraform 10-14 Feb 2025 Batch

## Kindly complete your pre-test ( Don't use the demo credentials )
<pre>
- Login to the RPS Cloud lab machine
- In the cloud machine, you will find an excel sheet in the Ubuntu Desktop
- With the login credentials shared in the excel sheet, proceed with the pre-test
- Once you are done with the pre-test, kindly inform me so we know when to start the training
</pre>

## Info - Installing Ansible core in Ubuntu
```
sudo apt update
sudo apt install software-properties-common
sudo add-apt-repository --yes --update ppa:ansible/ansible
sudo apt install ansible -y
```

## Info - Installing Docker Community Edition in Ubuntu
```
# Add Docker's official GPG key:
sudo apt-get update
sudo apt-get install ca-certificates curl
sudo install -m 0755 -d /etc/apt/keyrings
sudo curl -fsSL https://download.docker.com/linux/ubuntu/gpg -o /etc/apt/keyrings/docker.asc
sudo chmod a+r /etc/apt/keyrings/docker.asc

# Add the repository to Apt sources:
echo \
  "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.asc] https://download.docker.com/linux/ubuntu \
  $(. /etc/os-release && echo "$VERSION_CODENAME") stable" | \
  sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update

sudo apt-get install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin -y

sudo systemctl enable docker
sudo systemctl start docker
sudo usermod -aG docker $USER
sudo su $USER
docker --version
docker images
```

## Info - Installing Terraform in Ubuntu
```
sudo apt-get update && sudo apt-get install -y gnupg software-properties-common

wget -O- https://apt.releases.hashicorp.com/gpg | \
gpg --dearmor | \
sudo tee /usr/share/keyrings/hashicorp-archive-keyring.gpg > /dev/null

gpg --no-default-keyring \
--keyring /usr/share/keyrings/hashicorp-archive-keyring.gpg \
--fingerprint

echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] \
https://apt.releases.hashicorp.com $(lsb_release -cs) main" | \
sudo tee /etc/apt/sources.list.d/hashicorp.list

sudo apt update
sudo apt-get install terraform
terraform -install-autocomplete
```

## Check your lab
Check if ansible core is installed
```
ansible --version
```
Expected output
![image](https://github.com/user-attachments/assets/e93ca9df-15d3-45bc-964f-646b834d5b26)

Check if docker community edition is installed
```
docker info
docker --version
```

Expected output
![image](https://github.com/user-attachments/assets/eb37436d-1b15-4ce6-ba45-d59ef457ee65)
![image](https://github.com/user-attachments/assets/cbdc12d2-458d-4035-941b-157cfb89c942)
![image](https://github.com/user-attachments/assets/f9146be5-4789-4888-b9bf-0fcdad1ed8e7)

Check if Terraform is installed
```
terraform version
```

