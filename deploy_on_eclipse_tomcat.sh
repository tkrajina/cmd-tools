#!/bin/bash

set -e

tomcat_instance=$1
app_name=$2
workspace_location=$3

if [ -z "$workspace_location" ]
then
    workspace_location=".."
    echo "No workspace location specified, using: $workspace_location"
fi

if [ -z "$tomcat_instance" ] || [ -z "$app_name" ]
then
    echo "Usage:"
    echo "deploy_on_eclipse_tomcat.sh tomcat_instance app_name [workspace_location]"
    exit 1
fi

if [ -d $tomcat_instance ]
then
    webapps_path=$tomcat_instance
else
    webapps_path=$workspace_location/.metadata/.plugins/org.eclipse.wst.server.core/$tomcat_instance/webapps/
fi

echo "Deploying to:$webapps_path"

if [ ! -d "$webapps_path" ]
then
    echo "Invalid path: $webapps_path"
    exit 1
fi

for app in $(ls $webapps_path)
do
    echo
    echo "Removing $webapps_path/$app ?"
    echo
    rm --interactive=once -r -v $webapps_path/$app
done

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
