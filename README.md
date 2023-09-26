Pre requisites:  
vscode  
go extension in VSCODE & install if there is any recomendation from VScode for go  


Step 1: git clone https://github.com/GoPlugin/Xdc_WebSocket_Subscription_Issue.git
Step 2: navigate to Xdc_WebSocket_Subscription_Issue  
Step 3: Replace your WSS URL at line 15 in test.go  
Step 4: execute command "go run test.go" in VSCode terminal and follow the logs  

Issue can be observed only when we have number Addresses and Topics greater than four.
