## STEPS to Build, Copy to AWS, Run on EC2
Make sure you code has no local reference to Port 8080, it will be 80 on the server

- Login to AWS Console and Launch an Linux AMI instance

- Launch an Linux EC2 Instance Freetier default t2.micro will do. 
    - ![#f03c15](https://placehold.it/15/f03c15/000000?text=+) Make sure you edit your Default Security group to have  `HTTP at Port 80 addd`

-  SSH to the EC2 Instance
  
  `ssh -i ~/Documents/AWS/AWS-Solutions-Architect-Associate/AWS-Certified-Architect\ Asscoaite/myNorthvirginia1stkeypair.pem ec2-user@ec2-184-73-79-246.compute-1.amazonaws.com`
  
  - Create a Directory called `Project2` can be any name
    -  `cd Project2`
    -  `mkdir Templates`
    
    
- In Another Teminal 
 Build your Project
   `GOOS=linux GOARCH=amd64 go build -o TemplateAndAWS`
 
- Securely Copy (SCP) your executable in local machine to AWS EC2 Instance
 
 `scp -i ~/Documents/AWS/AWS-Solutions-Architect-Associate/AWS-Certified-Architect\ Asscoaite/myNorthvirginia1stkeypair.pem TemplateAndAWS ec2-user@ec2-3-87-10-115.compute-1.amazonaws.com:Project2`
 
 - Copy the Entire Template folder from local to remote with recursive copy on 
 
 `scp -i ~/Documents/AWS/AWS-Solutions-Architect-Associate/AWS-Certified-Architect\ Asscoaite/myNorthvirginia1stkeypair.pem -r Templates ec2-user@ec2-3-87-10-115.compute-1.amazonaws.com:Project2`

  - Run Below commands
     - `sudo yum update` 
     - `sudo chmod 700 TemplateAndAWS` 
     
     - `sudo ./TemplateAndAWS` Execute Binary if you dont want to Persist. If you want to persist follow Persistent Section
     
### Go to Web browser and give your EC2 Instance Public IP 

You should see your content

## Now Make your session Persistent

- Create a configuration file
    - `cd /etc/systemd/system/`
    - `sudo nano <filename>.service` : Change <filename> to your defined file name
    and Edit the <filename>.service with below code. I have used `rakesh.service`
    
    ```
       [Unit]
       Description=Go Server
       
       [Service]
       ExecStart=/home/ec2-user/Project2/TemplateAndAWS
        WorkingDirectory=/home/ec2-user/Project2

       User=root
       Group=root
       Restart=always
       
       [Install]
       WantedBy=multi-user.target
  ```
  Above Working directory `Project2` is the folder I created in EC2 instance. In the ExecPath within `Project2` I have my build binary file called `TemplateAndAWS`
  
- Now run below `systemctl` commands
    - `sudo systemctl enable rakesh.service `
    - `sudo systemctl start rakesh.service`
    - Check status of your service `sudo systemctl status rakesh.service`
    
    Note: I ran into a problem where in my local the Templates Folder name was starting with Upper case T but my code has templates/* and it was working fine but in AWS it failed to run
   
   