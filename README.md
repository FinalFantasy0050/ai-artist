# AI Artist  
AI Artist is a platform where anyone can easily create images using AI and access a text-based chatbot for free.  
This project is designed to increase accessibility to AI technology and enable efficient AI services on various devices.

## 0. Project Goals  
1. Enable anyone to create AI images without complex technical skills.  
2. Provide free access to a text-based AI chatbot.  
3. Support AI services on resource-limited devices like smartphones.  
4. Facilitate creative idea generation through character and story creation tools.  

## 1. Target Users  
- Users aged 10 to 30  
- Students and those needing learning tools  

## 2. Supported Platforms  
- Smartphones (Android) and PCs (Windows, MacOS, Linux)  

## 3. Getting Started  
The Bash shell script for initialization is located in the `/bash` directory of the project.

### 3.1 Gateway Server  
```bash
cd backend/gateway
```  
Move to the directory where the Gateway Server source code is located.

```bash
openssl req -x509 -nodes -newkey rsa:2048 -keyout key.pem -out cert.pem -days 365
```  
Generate an SSL certificate.  
If you want to add a password to the certificate, remove the `-nodes` option.

```bash
# If necessary
cd ../../bash

bash ./start_gateway_server.sh
```  
Start the server.

### 3.2 Image Generator Server  
```bash
# If necessary
cd bash

bash ./start_image_generator_server.sh
```  
Start the server.

### 3.3 Triton Server  
```bash
# If necessary
cd bash

bash ./start_triton_server.sh
```  
Start the server.

### 3.4 Chatbot Server  
```bash
cd backend/chatbot
```  
Move to the directory where the Chatbot Server source code is located.

```bash
vi .env
```  
Create or edit the `.env` file using `vi` or `vim`.  
Add the following content:  
```plaintext
OPENAI_API_KEY=
```  
Enter your OpenAI API Key.

```bash
# If necessary
cd ../../bash

bash ./start_chatbot_server.sh
```  
Start the server.

---

*This project is a practice project developed for a club activity.*