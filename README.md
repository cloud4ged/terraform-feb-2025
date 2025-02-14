# Terraform 10-14 Feb 2025 Batch

## My LinkedIn Profile
https://www.linkedin.com/in/jeganathan-swaminathan-2a6a6a6/

## My YouTube Channel
https://www.youtube.com/@JeganathanSwaminathan/videos

## My Blogs
<pre>
https://medium.com/@jegan_50867/docker-overview-be840f727b3

https://medium.com/tektutor/container-engine-vs-container-runtime-667a99042f3

https://medium.com/@jegan_50867/docker-commands-ba19387383b4
</pre>

## Info - Installing Ansible core in Ubuntu
Ansible core is already installed on the lab machine, hence this is just for your reference purpose.

```
sudo apt update
sudo apt install software-properties-common
sudo add-apt-repository --yes --update ppa:ansible/ansible
sudo apt install ansible -y
```

## Info - Installing Docker Community Edition in Ubuntu

Docker is already installed on the lab machine, hence this is just for your reference purpose.
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
Terraform is already installed on the lab machine, hence this is just for your reference purpose.

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

Check if golang is installed
```
go version
```

## Info - Installing AWX in Ubuntu ( Please don't do this in our lab environment - this is just for your reference )

First, let's install minikube
```
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube && rm minikube-linux-amd64
minikube start  --vm-driver=docker --addons=ingress
```

Expected output

Let's install kubectl 
```
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
chmod +x kubectl
sudo mv kubectl /usr/local/bin
```

Check if minikube is installed properly
```
kubectl get nodes
```

Expected output

Let's install AWX operator in minikube
```
sudo apt install make -y
cd ~
git clone https://github.com/ansible/awx-operator.git
cd awx-operator/
git checkout 2.19.0
export NAMESPACE=ansible-awx
make deploy

kubectl get pods -n ansible-awx
```

Expected output

Track the progress of awx installation
```
kubectl logs -f deployments/awx-operator-controller-manager -c awx-manager -n ansible-awx
```

Expected output


Access the AWX dashboard
```
cp awx-demo.yml awx-ubuntu.yml
kubectl create -f awx-ubuntu.yml -n ansible-awx
kubectl get pods -n ansible-awx
kubectl get svc -n ansible-awx
minikube service awx-ubuntu-service --url -n ansible-awx
```
Expected output


AWX Login Credentials
```
username - admin
```

To retrieve password
```
kubectl get secret -n ansible-awx | grep -i password
kubectl get secret awx-ubuntu-admin-password -o jsonpath="{.data.password}" -n ansible-awx | base64 --decode; echo
```

Expected output


If everything went smooth, you are expected to see similar page


To access the awx dashboard from other machines, you need to do port-forwarding
```
kubectl port-forward service/awx-ubuntu-service -n ansible-awx --address 0.0.0.0 10445:80 &> /dev/null &
```

You may now access the dashboard from other machines as
```
http://10.0.1.72:10445
```

## Installing Microsoft Visual Studio Code Editor in Ubuntu
```
sudo snap install code --classic
```

Expected output
![image](https://github.com/user-attachments/assets/7d8842c2-a84c-43fd-861e-2ceafeefbb7b)

