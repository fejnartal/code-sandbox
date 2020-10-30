#!/bin/sh
# To learn more about concourse and fly go to: https://concoursetutorial.com/basics/task-hello-world/

fly -t tutorial login -c http://concourse:8080 -u test -p test
fly -t tutorial execute -c task_hello_world.yml