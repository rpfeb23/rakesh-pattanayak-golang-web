## STEPS to Build, Copy to AWS, Run on EC2 and Take AMI
### Open a Terminal to Build your Project or use your IDE terminal

- Build your Project
   `GOOS=linux GOARCH=amd64 go build -o AWSGOPOSGRESBINARY`
   

### Open a new Terminal for interacting with AWS
- Launch an Linux EC2 Instance Freetier default t2.micro will do. 
    - ![#f03c15](https://placehold.it/15/f03c15/000000?text=+) Make sure you edit your Default Security group to have  `HTTP at Port 80 addd`

-  SSH to the EC2 Instance
  
  `ssh -i ~/Documents/AWS/AWS-Solutions-Architect-Associate/AWS-Certified-Architect\ Asscoaite/myNorthvirginia1stkeypair.pem ec2-user@ec2-184-73-79-246.compute-1.amazonaws.com`
  
  - Create a Directory called `myAWSProject` can be any name
    -  `mkdir myAWSProject`
  
  - Securely Copy (SCP) your executable in local machine to AWS EC2 Instance
   
   `scp -i ~/Documents/AWS/AWS-Solutions-Architect-Associate/AWS-Certified-Architect\ Asscoaite/NCaliforniaKeyPair.pem AWSGOPOSGRESBINARY ec2-user@ec2-18-144-8-166.us-west-1.compute.amazonaws.com:myAWSProject`
   
   - Copy the Entire Template folder from local to remote with recursive copy on 
    
    `scp -i ~/Documents/AWS/AWS-Solutions-Architect-Associate/AWS-Certified-Architect\ Asscoaite/NCaliforniaKeyPair.pem -r templates ec2-user@ec2-18-144-8-166.us-west-1.compute.amazonaws.com:myAWSProject`
   
  - Run Below commands
     - `sudo yum update` 
     - `sudo chmod 700 AWSGOPOSGRESBINARY` This will aloow only you to have rwx access. verify by `ls -la`
    

## Now Make your session Persistent
- Log in to your EC2 Instance using SSH .. If you already SSHed in previous step do "CTRL+C" to come out
- Create a configuration file
    - `cd /etc/systemd/system/`
    - `sudo nano rakesh.service` 
    
    ```
           [Unit]
           Description=Go Server
           
           [Service]
           ExecStart=/home/ec2-user/myAWSProject/AWSGOPOSGRESBINARY
            WorkingDirectory=/home/ec2-user/myAWSProject
    
           User=root
           Group=root
           Restart=always
           
           [Install]
           WantedBy=multi-user.target
      ```
  
  - ![#f03c15](https://placehold.it/15/f03c15/000000?text=+) Modify <Username> and <exepath> above  In my example I have usernmae as `ec2-user` as I am using Linux2 AMI and my execpath is `AWSGOPOSGRESBINARY`
  - Exit out of Editor (CTRL+X, then Y)
  
- Now run below `systemctl` commands
    - `sudo systemctl enable rakesh.service `
    - `sudo systemctl start rakesh.service`
    - Check status of your service `sudo systemctl status rakesh.service`
    
    Now if you will hit the EC2 public DNS it will execute your binary. If you now Exit out of your terminal where you have SSHed, the Public DNS will still serve your code
    
   
   