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

