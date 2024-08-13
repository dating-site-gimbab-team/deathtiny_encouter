# Step 1: Build the Vue.js application
FROM node:lts-slim as build-stage

# Set working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json to install dependencies
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code
COPY . .

# Build the application
RUN npm run build

# Step 2: Serve the application using Nginx
FROM nginx:stable as production-stage

# Copy the built files from the previous stage to the nginx html directory
COPY --from=build-stage /app/dist /usr/share/nginx/html

# Copy the default nginx configuration file
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Expose port 80
EXPOSE 80

# Start Nginx
CMD ["nginx", "-g", "daemon off;"]
