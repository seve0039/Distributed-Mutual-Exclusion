# Distributed-Mutual-Exclusion

1. 
You start the programming by going to the client folder in your terminal and running go run client.go -server=5400
you repeat this step in 2 other terminals but writing -server=5401 and -server=5402 at the end instead
This will give the clients the values for their serverports

2. 
To request the critical section you write "request-cs" in the termnial of the client that you want to acces the critical section
This should send a request through all other clients and grant you access to the critical section

3. 
The client will automatically exit the critical section
