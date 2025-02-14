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


