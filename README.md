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
```

## Info - Installing Docker Community Edition in Ubuntu
```
```

## Info - Installing Terraform in Ubuntu
## Installing Terraform in Ubuntu
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

Checking terraform installation
```
terraform -version
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

