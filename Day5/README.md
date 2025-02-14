# Day 5

## Lab - Download jenkins and launching it on the RPS Lab machine

Download the jenkins.war LTS from jenkins.io/download page

You can then launch jenkins from the terminal 
```
cd ~/Downloads
java -jar ./jenkins.war
```

Expected output
![image](https://github.com/user-attachments/assets/47254921-622d-4ec1-a881-900e19e1ac45)
![image](https://github.com/user-attachments/assets/682e7fa1-2ac7-45b8-b5fb-ad5ef266470a)

You can access the Jenkins dashboard
<pre>
http://localhost:8080  
</pre>

Expected output
![image](https://github.com/user-attachments/assets/87625aae-995f-44be-8926-8912760ce3f6)

You can copy/paste the url to copy the one time admin password in the web browser
![image](https://github.com/user-attachments/assets/5c721468-7464-4bc4-9bf7-5921e3e602fd)

Click on continue, after pasting the initial admin password
![image](https://github.com/user-attachments/assets/ace8903c-02cb-4830-9eda-6d2a5291f1a3)
![image](https://github.com/user-attachments/assets/26ae3458-ef80-46f0-9998-18adbc4961aa)

Select "Install suggested plugins"
![image](https://github.com/user-attachments/assets/11d0fa27-365f-43cd-883d-a25c7bc1bfb5)
![image](https://github.com/user-attachments/assets/376aef9e-3f81-4f24-81da-5755604db594)
![image](https://github.com/user-attachments/assets/dfe9ca0b-2ce0-4f92-b80a-0a21da2abdf6)

You need to create a user
![image](https://github.com/user-attachments/assets/5219f585-6332-4266-bdc0-5cb7b64235df)
![image](https://github.com/user-attachments/assets/5240aae6-a7c5-4a85-b41e-3bd7ff185144)
Click on "Save and Continue" button

![image](https://github.com/user-attachments/assets/76cea878-d1e8-4c3e-8294-68ec9eabae17)
Click on "Save and Finish"

![image](https://github.com/user-attachments/assets/95c60ee4-c818-46bd-90b4-56f074dbec3f)
Click on "Start using Jenkins"
![image](https://github.com/user-attachments/assets/66bfb2df-f873-4d9f-8bb1-bb50bfc09b76)

Click "Manage Jenkins" 
![image](https://github.com/user-attachments/assets/fc21cfbc-91e5-4b42-ad0b-34144b4fed6e)

Click "Plugins"
![image](https://github.com/user-attachments/assets/1a34c552-11be-44cb-a762-545a4b078c6c)

Click "Available Plugins"
![image](https://github.com/user-attachments/assets/5b9d5f95-644b-4f61-b48a-438bed408f18)

## Lab - Let's create CI Job to run the Terraform scripts whenever there is a code commit to our terraform-feb-2025.git code repository
Navigate to Jenkins Dashboard
![image](https://github.com/user-attachments/assets/a80fd698-0843-4892-8557-ba048fe5f686)

Click "Create a Job"
![image](https://github.com/user-attachments/assets/f28e0eb7-75a3-4780-9b02-9e893c727e45)

Select "Freestyle Job" and type "Run Terraform Scripts" under "Enter an item name"
![image](https://github.com/user-attachments/assets/435de084-dea5-46c6-ac0d-350d8ac521e8)
Click "Ok"
![image](https://github.com/user-attachments/assets/a2117ec5-fc5e-439a-99f1-d1620e817cec)

General section Description
![image](https://github.com/user-attachments/assets/4a64fc06-5892-4942-b944-9ead8c81958f)

Under Source Code Managment, select "Git"
![image](https://github.com/user-attachments/assets/3558f0ea-7206-45aa-8674-44a99e5ea4c4)

Triggers
![image](https://github.com/user-attachments/assets/def35dbd-8b5f-4c02-afa7-074e3a10466e)
Configure Jenkins to Poll GitHub repository every 2 minutes
<pre>
H/02 * * * *
</pre>
![image](https://github.com/user-attachments/assets/0463f0c1-5d9f-48be-b422-fc74f8c13aa8)

Click "Save" to avoid loosing the changes done so far
![image](https://github.com/user-attachments/assets/c0ce0f97-0e38-4b6d-a514-ce2cc7eb9bd1)
![image](https://github.com/user-attachments/assets/d0bbc35d-ebc8-477e-a19b-9e2ac7a54d66)

To continue with the Job configuration, click on Configure
![image](https://github.com/user-attachments/assets/e445a7a4-f22d-4835-ad83-a6d87b735980)
![image](https://github.com/user-attachments/assets/d3ca2c2a-93e5-403f-8ddc-82d2f5e9a50b)
![image](https://github.com/user-attachments/assets/cda71edb-793b-4a55-a8ba-930bf500786d)
![image](https://github.com/user-attachments/assets/ba5fd11f-63ea-449a-923e-d79bb2aefc86)
![image](https://github.com/user-attachments/assets/159a5ac4-0ce7-42e7-9424-51ccfe5b9a67)

Build Steps
Click "Add Build Step"
![image](https://github.com/user-attachments/assets/9d1899b0-ab22-4c0b-923a-099b54f096c7)
Let's select "Execute shell"
![image](https://github.com/user-attachments/assets/a8cb8df8-46e2-423d-acbd-84b67983ac21)
Type the below commands
<pre>
cd Day5/invoking-ansibleplaybook-from-terraform
terraform init
terraform apply --auto-approve  
</pre>
![image](https://github.com/user-attachments/assets/74b20c1a-a8cf-4f1b-8028-998eeb685d92)

Click "Save" button
![image](https://github.com/user-attachments/assets/6ddcc2e8-9a88-40dc-96ad-5163ffb82e4f)
![image](https://github.com/user-attachments/assets/98789a61-3033-483e-b235-9076024911b6)

Click "Build Now"
![image](https://github.com/user-attachments/assets/213478db-0199-4efe-9f74-19f4a3863714)
![image](https://github.com/user-attachments/assets/54b45fe3-9ccf-4098-b45a-ec9b3354baaf)

Let's wait for the build to get triggered automatically
![image](https://github.com/user-attachments/assets/f1887101-b628-4e05-a203-866ffbfc8c31)
![image](https://github.com/user-attachments/assets/25a09248-9987-4001-bceb-7fa0d2e5be09)
![image](https://github.com/user-attachments/assets/2e649190-2488-4dac-89ee-1c59eea6a083)
![image](https://github.com/user-attachments/assets/3a56ef3b-648e-4692-b079-2f631f33c22b)
![image](https://github.com/user-attachments/assets/1a293766-2069-4ce8-a13a-11002cafb2f8)

## Lab - Let's create another Freestyle job in Jenkins
Currently Jenkins Dashboard looks as below
![image](https://github.com/user-attachments/assets/2aa043c0-29ff-4164-8107-f62c98dfdf38)

Click "New item"
![image](https://github.com/user-attachments/assets/14f59f1c-170a-488d-9aea-b96c8459cc77)
Select "Freestyle project" and type "Build Docker Image" under "Enter an item name"
![image](https://github.com/user-attachments/assets/4fef533b-dc80-4db9-bd89-4743865f8ead)
Click "Ok"
![image](https://github.com/user-attachments/assets/c2304d59-81f0-4677-a0b7-08c2c291bc95)
![image](https://github.com/user-attachments/assets/994a242b-5d47-4b78-a673-604a4028199c)

Source Code Management section
![image](https://github.com/user-attachments/assets/cf3aa36f-59fb-4af1-ade0-85e5d2ea4b22)
Select "Git"
![image](https://github.com/user-attachments/assets/6d1f8c3f-29cf-4e60-b248-3aa6375c65d5)
![image](https://github.com/user-attachments/assets/77c24d88-c73f-4672-9e79-07512f97dc85)
![image](https://github.com/user-attachments/assets/2fbc4422-adf1-4e13-9eb6-8cccceadec07)

Triggers section
![image](https://github.com/user-attachments/assets/15a0425e-39cf-4ae1-bd74-83194a75396a)
![image](https://github.com/user-attachments/assets/e5b01103-db28-4a0a-b542-a6c96bfe4580)

Build Steps
![image](https://github.com/user-attachments/assets/b67731a2-af32-4305-91ef-408717a31978)
![image](https://github.com/user-attachments/assets/52b6b4d0-3055-4641-90f4-acf7322d0fe0)
Select "Execute shell"
Type the below commands
<pre>
cd Day5/CustomDockerAnsibleNodeImages/ubuntu-ansible
cp ~/.ssh/id_ed25519.pub authorized_keys
docker build -t tektutor/ubuntu-ansible:1.0 .
rm authorized_keys
docker images
</pre>
![image](https://github.com/user-attachments/assets/305ac7b3-5a30-4bd5-b343-994203a33543)
![image](https://github.com/user-attachments/assets/c6209c4f-3a81-4028-8689-ccfff5e6a120)
Click "Save" button
![image](https://github.com/user-attachments/assets/df0f5219-1d35-4375-958f-a435c9059924)
![image](https://github.com/user-attachments/assets/d2db31ee-a6f7-4708-953f-bc70f930e801)


Buid Progress
![image](https://github.com/user-attachments/assets/21ffe510-6607-4312-89ec-454a90f97963)
![image](https://github.com/user-attachments/assets/53bd7e2b-d4fe-4aa5-95f8-4e7833b873ad)
![image](https://github.com/user-attachments/assets/de2fbab2-23ac-4509-8e71-f34004d72f93)
![image](https://github.com/user-attachments/assets/6540f030-63c0-45f5-89e3-b733a68df638)

## Lab - Let's create a CI/CD Pipeline with the 2 Jenkins Jobs we created
![image](https://github.com/user-attachments/assets/9ba5e16b-2389-4b1b-bd48-4f3f1e1c20b6)
Click "+" New view
![image](https://github.com/user-attachments/assets/a891012f-4fe9-4c0c-803a-e6565a46e3e5)
![image](https://github.com/user-attachments/assets/342d72cd-d6e8-4325-8fb0-bc1885310960)
![image](https://github.com/user-attachments/assets/60a8f78f-8da1-4c9d-9fb5-ca7b88caa68b)
Click "Create" button
![image](https://github.com/user-attachments/assets/3031b94f-ab3a-4dbe-9d55-da1650dea462)

## Lab - let's a declarative devops pipeline
![image](https://github.com/user-attachments/assets/e3df7c83-3e3f-4c86-9890-e4016a37507e)
![image](https://github.com/user-attachments/assets/0707c504-5d4e-411c-ba34-81ded4b02f64)
Click "Ok"
![image](https://github.com/user-attachments/assets/1948997e-c9cd-4a5d-bcc1-09dde64156bc)

General
![image](https://github.com/user-attachments/assets/75b5dca2-0064-49a3-bf8b-e161baca3cd4)

Triggers
![image](https://github.com/user-attachments/assets/aa81472d-e42c-483c-86e0-827ae18069ed)

Pipeline
![image](https://github.com/user-attachments/assets/5733945a-db2e-455b-b3f6-a763529cde10)
![image](https://github.com/user-attachments/assets/38a358f8-993d-4ec7-9e42-f39dc97ada6e)
![image](https://github.com/user-attachments/assets/b89fc7e9-9dc1-4a77-9528-5fc985e88149)

Script Path
![image](https://github.com/user-attachments/assets/32283b35-3ace-4c05-83ff-4c5588185843)

Save
![image](https://github.com/user-attachments/assets/30e2c23a-413f-44aa-a12b-b1b9fe8f5955)

Build Progress
![image](https://github.com/user-attachments/assets/d4bafe27-f26c-45a1-9cce-752544d2e038)

