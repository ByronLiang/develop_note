#!/usr/bin/env bash

set -e

# Install Some Basic Packages

apt install -y build-essential git autoconf supervisor software-properties-common

# Set My Timezone

# ln -sf /usr/share/zoneinfo/UTC /etc/localtime

# Install PHP Stuffs

add-apt-repository -y ppa:ondrej/php
apt update
apt install -y php7.2-cli php7.2-gd php-apcu php7.2-curl php7.2-imap php7.2-mysql php7.2-readline php7.2-mbstring php7.2-xml php7.2-zip php7.2-intl php7.2-bcmath php7.2-soap

# Add new user ganguo

useradd -m -U -s /bin/bash ganguo

# Install Composer

curl -sS https://getcomposer.org/installer | php
mv composer.phar /usr/local/bin/composer
# wget http://data.dev.ganguo.hk/static/composer
# mv composer /usr/local/bin/
chmod a+x /usr/local/bin/composer

# /usr/local/bin/composer config -g repo.packagist composer https://packagist.phpcomposer.com
su ganguo <<'EOF'
/usr/local/bin/composer config -g repo.packagist composer https://mirrors.aliyun.com/composer/
EOF

# Set Some PHP CLI Settings

sed -i "s/max_execution_time = .*/max_execution_time = 120/" /etc/php/7.2/cli/php.ini
sed -i "s/memory_limit = .*/memory_limit = 512M/" /etc/php/7.2/cli/php.ini
# sed -i "s/;date.timezone.*/date.timezone = UTC/" /etc/php/7.2/cli/php.ini

# Install Nginx & PHP-FPM

apt install -y nginx php7.2-fpm

rm /etc/nginx/sites-enabled/default
rm /etc/nginx/sites-available/default
systemctl restart nginx

sed -i "s/;cgi.fix_pathinfo=1/cgi.fix_pathinfo=0/" /etc/php/7.2/fpm/php.ini
sed -i "s/memory_limit = .*/memory_limit = 512M/" /etc/php/7.2/fpm/php.ini
sed -i "s/upload_max_filesize = .*/upload_max_filesize = 10M/" /etc/php/7.2/fpm/php.ini
sed -i "s/post_max_size = .*/post_max_size = 10M/" /etc/php/7.2/fpm/php.ini
# sed -i "s/;date.timezone.*/date.timezone = UTC/" /etc/php/7.2/fpm/php.ini

# Set The Nginx & PHP-FPM User

# sed -i "s/user www-data;/user ganguo;/" /etc/nginx/nginx.conf
# sed -i "s/# server_names_hash_bucket_size.*/server_names_hash_bucket_size 64;/" /etc/nginx/nginx.conf
sed -i "s/user = www-data/user = ganguo/" /etc/php/7.2/fpm/pool.d/www.conf
sed -i "s/group = www-data/group = ganguo/" /etc/php/7.2/fpm/pool.d/www.conf
sed -i "s/listen\.owner.*/listen.owner = ganguo/" /etc/php/7.2/fpm/pool.d/www.conf
sed -i "s/listen\.group.*/listen.group = ganguo/" /etc/php/7.2/fpm/pool.d/www.conf
sed -i "s/;listen\.mode.*/listen.mode = 0666/" /etc/php/7.2/fpm/pool.d/www.conf

systemctl restart nginx
systemctl restart php7.2-fpm

# Add User To WWW-Data

usermod -a -G www-data ganguo

# Add User to sudo

usermod -a -G sudo ganguo

# visudo
# ganguo    ALL=(ALL) NOPASSWD:ALL

# Install MySQL

debconf-set-selections <<< "mysql-community-server mysql-community-server/data-dir select ''"
debconf-set-selections <<< "mysql-community-server mysql-community-server/root-pass password secret"
debconf-set-selections <<< "mysql-community-server mysql-community-server/re-root-pass password secret"
apt install -y mysql-server

# Configure MySQL Password Lifetime
sed -i '1s/^/[mysqld]\n /' /etc/mysql/my.cnf
echo "default_password_lifetime = 0" >> /etc/mysql/my.cnf

# Configure MySQL Remote Access

# sed -i '/^bind-address/s/bind-address.*=.*/bind-address = 127.0.0.1/' /etc/mysql/my.cnf

mysql --user="root" --password="secret" -e "GRANT ALL ON *.* TO root@'127.0.0.1' IDENTIFIED BY 'secret' WITH GRANT OPTION;"

systemctl restart mysql

mysql --user="root" --password="secret" -e "CREATE USER 'homestead'@'127.0.0.1' IDENTIFIED BY 'secret';"
mysql --user="root" --password="secret" -e "GRANT ALL ON *.* TO 'homestead'@'127.0.0.1' IDENTIFIED BY 'secret' WITH GRANT OPTION;"
mysql --user="root" --password="secret" -e "GRANT ALL ON *.* TO 'homestead'@'localhost' IDENTIFIED BY 'secret' WITH GRANT OPTION;"
# mysql --user="root" --password="secret" -e "GRANT ALL ON *.* TO 'homestead'@'%' IDENTIFIED BY 'secret' WITH GRANT OPTION;"
mysql --user="root" --password="secret" -e "FLUSH PRIVILEGES;"
# mysql --user="root" --password="secret" -e "CREATE DATABASE homestead;"

systemctl restart mysql

# Add Timezone Support To MySQL

mysql_tzinfo_to_sql /usr/share/zoneinfo | mysql --user=root --password=secret mysql

# Enable Swap Memory

# dd if=/dev/zero of=/var/swap.1 bs=1M count=2048
# mkswap /var/swap.1
# chmod 0600 /var/swap.1
# swapon /var/swap.1

# Install Nodejs
curl -sL https://deb.nodesource.com/setup_8.x | bash -
apt install -y nodejs
npm config set registry https://registry.npm.taobao.org -g
npm config set disturl https://npm.taobao.org/dist -g
npm config set loglevel=http -g
npm config set sass_binary_site=https://npm.taobao.org/mirrors/node-sass -g

# Install Certbot
add-apt-repository ppa:certbot/certbot
apt update
apt install -y certbot
# certbot certonly --webroot -w /var/www/html -d example.com
# m h  dom mon dow   command
# 15 5 * * * certbot renew --post-hook "systemctl restart nginx"
# 30 5 * * * systemctl restart nginx

# Backup mysql
# 0 */12 * * * mysqldump -u homestead -psecret homestead | gzip -c > ~/backup/mysql/homestead/homestead.$(date +"\%Y\%m\%dT\%H").sql.tar.gz

apt -y upgrade
apt -y autoremove
