FROM nginx:alpine

# Remove default nginx config
RUN rm /etc/nginx/conf.d/default.conf

# Copy the template to the templates directory so envsubst can process it on startup
COPY default.conf.template /etc/nginx/templates/default.conf.template
