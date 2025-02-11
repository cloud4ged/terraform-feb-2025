# Day 2

## Lab - Starting Ansible automation platform
```
minikube status
minikube start
kubectl get pods -n ansible-awx
kubectl get svc -n ansible-awx
minikube service awx-demo-service --url -n ansible-awx
```

Expected output
![image](https://github.com/user-attachments/assets/8d5f02cf-75cd-43a6-a84d-f6ce776b9846)

AWX Login Credentials
```
username - admin
```

To retrieve password
```
kubectl get secret -n ansible-awx | grep -i password
kubectl get secret awx-demo-admin-password -o jsonpath="{.data.password}" -n ansible-awx | base64 --decode; echo
```

To access the awx dashboard from other machines, you need to do port-forwarding
```
kubectl port-forward service/awx-demo-service -n ansible-awx --address 0.0.0.0 10445:80 &> /dev/null &
```

You may now access the dashboard from other machines as
```
http://10.0.1.72:31225
```

Expected output
![image](https://github.com/user-attachments/assets/3081631a-89f0-46cc-b84d-412c0ea41dd0)
![image](https://github.com/user-attachments/assets/6c8a2104-c97f-4f74-80cd-d9adb6443dc0)


## Lab - Creating a project in Ansible Tower
Navigate to Resource --> Projects
![image](https://github.com/user-attachments/assets/ee41a93c-96cc-44f5-b799-bc2913d25f2c)

Click "Add" button
![image](https://github.com/user-attachments/assets/0169dd9c-89c1-42d5-a006-9a99674b3f6a)
Select "Git" under Source Control Type
![image](https://github.com/user-attachments/assets/9d9b571c-dfa4-4a04-84e0-69aeddd4c094)
![image](https://github.com/user-attachments/assets/96644d63-8b22-4902-a888-f0fa4b412a77)
![image](https://github.com/user-attachments/assets/5dfeb1aa-7fa4-492d-93ae-3f289a7abe85)

Click "Save" button
![image](https://github.com/user-attachments/assets/fcfe69a8-2fff-41c9-b353-16f811affb69)
![image](https://github.com/user-attachments/assets/47964725-b73d-466f-b865-3cc2d39f657a)
![image](https://github.com/user-attachments/assets/cd1c6d54-417c-4dfc-92e6-066cb1a8e0ab)

## Lab - Let's create a credential in Ansible Tower to store the private key we generated

Navigate to Resources --> Credentials
![image](https://github.com/user-attachments/assets/59c2b03d-6c02-4fc0-b9df-a944ea598b6f)

Click "Add" button
![image](https://github.com/user-attachments/assets/6d9007c9-06d9-4d06-9e04-441b4ccf7420)
![image](https://github.com/user-attachments/assets/09414934-4e19-4fdb-9d26-c433c951e630)
![image](https://github.com/user-attachments/assets/dd1cae37-154a-4aaf-8727-62f91b06f80f)

Click "Save" button
![image](https://github.com/user-attachments/assets/d2d2503c-6d65-4f2a-9d4e-0e694e2b10a3)



## Lab - Let's create an inventory in Ansible Tower
