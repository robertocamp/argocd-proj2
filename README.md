# argocd-proj2
another demo project to work with argocd


## Docker


### basic setup

https://docs.docker.com/language/golang/build-images/

1. How can I list the images my local Docker:  `docker images`

### multistage-image build


1. **Build and Run for `arm64` locally:**

   ```sh
   docker buildx build --platform linux/arm64 -t argocd-proj2-arm64 --load .
   docker run --rm -p 8000:8000 argocd-proj2-arm64
   ```

2. **Build and Run for `amd64` locally:**

   ```sh
   docker buildx build --platform linux/amd64 -t argocd-proj2-amd64 --load .
   docker run --rm -p 8000:8000 argocd-proj2-amd64
   ```

3. **Tag the images for upload:**

   ```sh
   docker tag argocd-proj2-arm64 yourusername/argocd-proj2:dev-arm64
   docker tag argocd-proj2-amd64 yourusername/argocd-proj2:dev-amd64
   ```

4. **Push the tagged images to Docker Hub:**

   ```sh
   docker login
   docker push yourusername/argocd-proj2:dev-arm64
   docker push yourusername/argocd-proj2:dev-amd64
   ```

5. **Build and Push Multi-Platform Image to Docker Hub:**

   When you're ready to push the multi-platform image:

   ```sh
   docker buildx build --platform linux/amd64,linux/arm64 -t yourusername/argocd-proj2:latest --push .
   ```

Here is a summary of the complete workflow:

1. **Build and Run for `arm64` locally:**

   ```sh
   docker buildx build --platform linux/arm64 -t argocd-proj2-arm64 --load .
   docker run --rm -p 8000:8000 argocd-proj2-arm64
   ```
    - connect to localhost:8000 


2. **Build and Run for `amd64` locally:**

   ```sh
   docker buildx build --platform linux/amd64 -t argocd-proj2-amd64 --load .
   docker run --rm -p 8000:8000 argocd-proj2-amd64
   ```

3. **Tag the images for upload:**

   ```sh
   docker tag argocd-proj2-arm64 qslrdocker/argocd-proj2:dev-arm64
   ```

4. **Push the tagged images to Docker Hub:**

   ```sh
   docker login
   docker push qslrdocker/argocd-proj2:dev-arm64
   ```

5. **Build and Push Multi-Platform Image to Docker Hub:**

   ```sh
   docker buildx build --platform linux/amd64,linux/arm64 -t qslrdocker/argocd-proj2:v001 --push .
   ```

By following these steps, you ensure that your Docker image is built and tested for both `arm64` and `amd64` architectures, tagged for development, and eventually pushed to Docker Hub as a multi-platform image.


### docker container stop and local clean up


To stop a running Docker container gracefully and release the local port 8000, you can use the `docker stop` command followed by the container ID or name. Here’s how you can do it:

1. **Identify the running container:**

   If you don’t already know the container ID or name, list all running containers using:

   ```sh
   docker ps
   ```

   This command will output a list of running containers with their IDs and names.

2. **Stop the container:**

   Use the `docker stop` command followed by the container ID or name to stop the container gracefully. For example, if the container ID is `abc123`:

   ```sh
   docker stop abc123
   ```

   Or if you named your container, you can use the name instead of the ID:

   ```sh
   docker stop container_name
   ```

3. **Verify that the container is stopped:**

   Run `docker ps` again to ensure the container is no longer running.

### Example Workflow

Assuming you have a container running on port 8000:

1. **List running containers:**

   ```sh
   docker ps
   ```

   Output example:

   ```
   CONTAINER ID   IMAGE                    COMMAND                  CREATED         STATUS         PORTS                    NAMES
   abc123         argocd-proj2-arm64       "/docker-gs-ping"        10 seconds ago  Up 9 seconds   0.0.0.0:8000->8000/tcp   elegant_turing
   ```

2. **Stop the container:**

   ```sh
   docker stop abc123
   ```

   Or using the container name:

   ```sh
   docker stop elegant_turing
   ```

3. **Verify the container has stopped:**

   ```sh
   docker ps
   ```

   If the container has stopped successfully, it will no longer appear in the list of running containers.

By following these steps, you can gracefully stop the running Docker container and release the local port 8000 on your system.


### Docker Hub

If you are using an IDP-based authentication mechanism via Google to log in to Docker Hub, you won't be able to use your Google credentials directly with the `docker login` command in the terminal. Instead, you can generate a Docker Hub Personal Access Token, which you can use to authenticate with the `docker login` command.

Here’s how you can do it:

1. **Generate a Personal Access Token:**

   - Log in to Docker Hub using the Docker Hub UI with your Google account.
   - Navigate to your account settings by clicking on your profile picture in the top-right corner and selecting "Account Settings."
   - In the Account Settings, go to the "Security" section.
   - Click on "New Access Token" under the "Access Tokens" section.
   - Provide a name for the token (e.g., "CLI Access Token") and generate the token.
   - Copy the generated token and save it securely. You will need it for the next step.

2. **Log in Using the Personal Access Token:**

   Use the `docker login` command in the terminal with your Docker Hub username and the Personal Access Token.

   ```sh
   docker login
   ```

   When prompted, enter your Docker Hub username and use the Personal Access Token as the password.

   ```sh
   Username: yourusername
   Password: youraccesstoken
   ```

3. **Push the Image to Docker Hub:**

   After successfully logging in, you can push your Docker image to Docker Hub.

   ```sh
   docker push yourusername/argocd-proj2:dev-arm64
   ```

#### Summary of Steps:

1. **Generate a Personal Access Token:**
   - Log in to Docker Hub using the UI.
   - Go to "Account Settings" > "Security" > "New Access Token."
   - Generate and save the token.

2. **Log in Using the Personal Access Token:**

   ```sh
   docker login
   ```

   Use your Docker Hub username and the generated Personal Access Token.

3. **Push the Image to Docker Hub:**

   ```sh
   docker push yourusername/argocd-proj2:dev-arm64
   ```

By following these steps, you can log in to Docker Hub using your Personal Access Token and push your Docker images successfully.


#### Docker Hub automation



Yes, you can store the Docker access token in your shell environment file to make the `docker login` process more seamless and less interactive. By setting up environment variables in your shell configuration file, you can use these variables in a script to log in to Docker Hub automatically.

Here's how you can do it:

1. **Store the Access Token in Your Shell Configuration File:**

   Open your shell configuration file (e.g., `~/.bashrc`, `~/.bash_profile`, `~/.zshrc`, depending on your shell) and add the following lines:

   ```sh
   export DOCKER_USERNAME=yourusername
   export DOCKER_ACCESS_TOKEN=youraccesstoken
   ```

   Replace `yourusername` with your Docker Hub username and `youraccesstoken` with your Docker Hub Personal Access Token.

2. **Apply the Changes:**

   Reload your shell configuration file to apply the changes:

   ```sh
   source ~/.bashrc
   ```

   or

   ```sh
   source ~/.zshrc
   ```

3. **Create a Script to Log In to Docker Hub:**

   Create a script that uses these environment variables to log in to Docker Hub. For example, create a file named `docker-login.sh`:

   ```sh
   #!/bin/sh

   echo $DOCKER_ACCESS_TOKEN | docker login -u $DOCKER_USERNAME --password-stdin
   ```

4. **Make the Script Executable:**

   Make the script executable:

   ```sh
   chmod +x docker-login.sh
   ```

5. **Run the Script to Log In:**

   Now you can run the script to log in to Docker Hub seamlessly:

   ```sh
   ./docker-login.sh
   ```

By following these steps, you can log in to Docker Hub automatically using a script that reads your stored credentials from environment variables. This makes the `docker login` sequence more seamless and less interactive.

### Summary of Steps:

1. **Add Credentials to Shell Configuration File:**

   ```sh
   export DOCKER_USERNAME=yourusername
   export DOCKER_ACCESS_TOKEN=youraccesstoken
   ```

2. **Apply the Changes:**

   ```sh
   source ~/.bashrc
   ```

3. **Create a Script for Docker Login:**

   ```sh
   #!/bin/sh

   echo $DOCKER_ACCESS_TOKEN | docker login -u $DOCKER_USERNAME --password-stdin
   ```

4. **Make the Script Executable:**

   ```sh
   chmod +x docker-login.sh
   ```

5. **Run the Script:**

   ```sh
   ./docker-login.sh
   ```

This approach allows you to log in to Docker Hub without having to manually enter your credentials each time.