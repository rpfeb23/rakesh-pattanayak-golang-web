
RDS Postgres
    1. Dev/Test
    2. Multi AZ
    3. Slect t2.Micro [IMPORTANT]
    4. Custom VPC
    5. CREATE SUBNET GROUOP
        from Drop down select AZ and Choose DB Subnet Group
        do this once for 1b and 1c in our example
        10.0.22.0/24 and 10.0.21.0/24
    6. Choose DB Security group
-------------------------------------------
launch a Bastion Host in Custom VPC, Public Subnet 1b and  bastion-security group
    User UserData
        #!/bin/bash
        sudo yum update -y
------------------------------------------
Launch an Instance in Private 
    1. in Custome VPC, Private Subnet app 1b, appserver SG
    2. UserData 
#!/bin/bash
sudo yum update -y
cd /etc/systemd/system/
echo '[Unit]
    Description=Go Server
           
    [Service]
    ExecStart=/home/ec2-user/myAWSProject/AWSGOPOSGRESBINARY
    WorkingDirectory=/home/ec2-user/myAWSProject
    
    User=root
    Group=root
    Restart=always
           
    [Install]
    WantedBy=multi-user.target' > rakesh.service

---------------------------------------
ssh-add NCaliforniaKeyPair.pem 

SSH into Bastion Host : ssh -A -i NCaliforniaKeyPair.pem ec2-user@ec2-18-144-156-231.us-west-1.compute.amazonaws.com

    mkdir myAWSProject

    SSH to the appserver instance via Bastion (key forwarding) ==> ssh ec2-user@10.0.11.112
    1. verify inside cd /etc/systemd/system/ 'rakesh.service' is there or not
    2. cat rakesh.service
    3. mkdir myAWSProject
    4. Exit from Appserver
---------------------------------------
Goto GOLAND
    1. GOOS=linux GOARCH=amd64 go build -o AWSGOPOSGRESBINARY

Copy to Bastion
    2. scp -i ~/Documents/AWS/AWS-Solutions-Architect-Associate/AWS-Certified-Architect\ Asscoaite/NCaliforniaKeyPair.pem AWSGOPOSGRESBINARY ec2-user@ec2-18-144-156-231.us-west-1.compute.amazonaws.com:myAWSProject
  
    3. scp -i ~/Documents/AWS/AWS-Solutions-Architect-Associate/AWS-Certified-Architect\ Asscoaite/NCaliforniaKeyPair.pem -r templates ec2-user@ec2-18-144-156-231.us-west-1.compute.amazonaws.com:myAWSProject


---------------------------------------
SSH into Bastion Host
    Verify the folder cd myAWSProject 

    Ensure Both AWSGOPOSGRESBINARY and templates are present

Copy to Private Subnet 1b App server 
    scp -r myAWSProject ec2-user@10.0.11.112:

-----------------------------------------
SHH into Private App Instance ==> ssh ec2-user@10.0.11.112

In the SSH Intance Terminal
Verify the folder cd myAWSProject 
Ensure Both AWSGOPOSGRESBINARY and templates are present

Run
cd /etc/systemd/system/
sudo systemctl enable rakesh.service
sudo systemctl start rakesh.service
sudo systemctl status rakesh.service --> Check Status of your Service

-----------------------------------
Go to Loadbalancer Target Group. Register the Private EC2 Instance
Now Access the ALB DNS Endpoint
All should be working fine
---------------------------------
Now Take AMI from Appserver
---------------------------------
Verify AMI is working fine or not
   Launch a new instance from the Above AMI (It should be under myAMI)
   Select Custom VPC
   Try Subnet Private -app 1c
   User Data

        #!/bin/bash
        sudo yum update -y
        cd /etc/systemd/system/
        sudo systemctl enable rakesh.service
        sudo systemctl start rakesh.service
   Security Group : appserver security group

   Now go to Load Balancer Target Group and Add the newly created private EC2 Instance as registered target

   Validate the ALB DNS Endpoint, it should be switching back and forth between your Instances and all links should be working
--------------------------------------------
AUTOSCALING

Create a new Launch template
    Choose your AMI from dropdown
    Instance type t2.Micro
    Select Keypair if you want SSH in future
    Choose Security group as appserver-sg
    ADVANCE DETAIL Put below User Data
        #!/bin/bash
        sudo yum update -y
        cd /etc/systemd/system/
        sudo systemctl enable rakesh.service
        sudo systemctl start rakesh.service

Now Launch a New Autoscaling Group
    select launch template (our template GolangRESTapiPostgreSQL)
    Select VPC as Custome VPC
    Add Subnet Private app 1b and Private app 1c
    Group Size : Start with 2 Instances to see 1 EC2 in each Subnet
    Click Advanced
        Check Receive Traffic from Load Balancer
        Now in Targetgroup give our Target group

    Next in Scaling Policy 
        Scale between 1 an 4
        Click Scale the Auto Scaling group using step or simple scaling policies
        Give INCREASE Spec
        Give DECREASE Spec
        Add Tag [Name : FromASG] This will ease in identifying instances 

CHECK The Load Balancer TARGET GROUP to see if the ASG Instances are added or not (it may take few mins as instances are getting launched)

NOW got to ALB DNS Endpoint 



