FROM nginx:stable-alpine 

RUN rm /etc/nginx/nginx.conf /etc/nginx/conf.d/default.conf

# Copy configuration
COPY ./nginx/nginx.conf /etc/nginx/

EXPOSE 8020

ENTRYPOINT ["nginx", "-g", "daemon off;"]