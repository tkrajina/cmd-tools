#!/bin/bash

tomcat_instance=$1
app_name=$2

if [ -z "$tomcat_instance" ] || [ -z "$app_name" ]
then
    echo "Usage:"
    echo "deploy_on_eclipse_tomcat.sh tomcat_instance app_name"
    exit 1
fi

webapps_path=../.metadata/.plugins/org.eclipse.wst.server.core/$tomcat_instance/webapps/

if [ ! -d "$webapps_path" ]
then
    echo "Invalid path: $webapps_path"
    exit 1
fi

mvn clean package -DskipTests=true || exit 1

if [ -d "$webapps_path/$app_name" ]
then
    rm -v -Rf "$webapps_path/$app_name" || exit 1
fi

if [ -f "$webapps_path/$app_name.war" ]
then
    rm -v "$webapps_path/$app_name.war" || exit 1
fi

cp -v target/*war "$webapps_path/$app_name.war"

if [ ! -d target/generated-resources/eclipse ]
then
    mkdir -p target/generated-resources/eclipse
fi
