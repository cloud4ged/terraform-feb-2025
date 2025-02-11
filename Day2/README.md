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
kubectl get secret awx-ubuntu-admin-password -o jsonpath="{.data.password}" -n ansible-awx | base64 --decode; echo
```

Expected output
![image](https://github.com/user-attachments/assets/3081631a-89f0-46cc-b84d-412c0ea41dd0)
![image](https://github.com/user-attachments/assets/6c8a2104-c97f-4f74-80cd-d9adb6443dc0)
