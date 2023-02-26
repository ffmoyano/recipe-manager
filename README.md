A simple recipe manager. The idea is to make a functional app with as few dependencies as possible.

there must be a .env file in the root folder (or whatever destination you define in the init method of the main file)  
with the following data (the values except the dbDriver are examples):

dbUrl=root:password@tcp(localhost:port)/example?parseTime=true  
port=:1234  
dbDriver=mysql