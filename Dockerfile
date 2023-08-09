# set the base image to Ubuntu 22.04 
FROM ubuntu:22.04 

# update the package list and install the MYSQL server
RUN apt-get update && \
    apt-get install -y mysql-server && \
    rm -rf /var/lib/apt/lists/* 


# Set the ownership of the MySQL data directory to the mysql user and group
RUN chown -R mysql:mysql /var/lib/mysql 

# Copy the MySQL configuration file into the container
COPY my.cnf /etc/mysql/my.cnf 

# define the arguments for the database name, user and password 
ARG DB 
ARG DB_USER 
ARG DB_PASSWORD 

# Create the database and user 
RUN /etc/init.d/mysql start && \
    mysql -u root -e "CREATE DATABASE $DB; CREATE USER '$DB_USER' IDENTIFIED BY '${DB_PASSWORD}'; GRANT ALL PRIVILEGES ON *.* TO '${DB_USER}'@'%'; GRANT GRANT OPTION ON *.* TO '${DB_USER}'@'%';" && \
    /etc/init.d/mysql stop

#Expose the mysql port 
EXPOSE 3306 

# Switch to the mysql user 
USER mysql 

# start the mysql server 
CMD [ "mysqld" ]