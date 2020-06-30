# Blockchain using fabric
requirements\n
- install go, version 1.14+
- install gin
- clone https://github.com/maruyari/fabric-samples directory into your home directory
- clone this directory in a location of your choice
- open terminal in the directory of this folder and run : "go run main.go"
- open localhost:8080 in browser
- website should be hosted in the above url
- open new terminal 
- cd into fabric-samples/test-network directory 
- run the folowing command- "./run.sh"  to deploy the blockchain
- if you get permission denied error
  --run  "chmod +x ./run.sh", then run "./run.sh" to deploy the blockchain. might take a few minutes. once its successful, the website is fully functional.
