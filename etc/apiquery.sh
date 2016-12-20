#!/usr/bin/env bash
# File : apiquery.sh
#
# Goal :
#       query todolist API
#
# History :
# 16/10/20 Creation (SFR)

############################
# debug                    #
############################
#set -x

############################
# declarations             #
############################
PROG_NAME=`basename $0`

# params
CREATE="false"
QUERY="false"
UPDATE="false"
DELETE="false"
VERBOSE=""

# check if running on osx or linux
IP=127.0.0.1
if [ ! -z "$DOCKER_MACHINE_NAME" ]; then
  IP=$(docker-machine ip $DOCKER_MACHINE_NAME)
fi

############################
# helper functions         #
############################

# verbose echo
vecho() {
  if [ "$VERBOSE" = "-v" ] ; then echo "$PROG_NAME: $*" ; fi
}

# help
usage() {
  echo "usage: $PROG_NAME [options] as follows :"
  echo "	[ -create : creates an entity        ]"
  echo "	[ -query : queries all entities      ]"
  echo "	[ -update : updates the first entity ]"
  echo "	[ -delete : deletes the first entity ]"
  exit 1
}

# parse options and parameters
param() {
      while [ $# -gt 0 ]
          do case $1 in
      		    -create) CREATE="true";;
      		    -query) QUERY="true";;
      		    -update) UPDATE="true";;
      		    -delete) DELETE="true";;
              -h) usage;;
      		    -v) VERBOSE="-v";;
          esac
	        shift
    done
}

# create
create() {
  vecho "Create task from file..."
  curl -s -X POST -H "Content-Type:application/json" -d @caroni.json ${IP}:8020/tasks | jq
}

# query
query() {
  vecho "Query all tasks..."
  curl -s ${IP}:8020/tasks | jq

  vecho "Retrieve first task by ID..."
  ID=$(curl -s ${IP}:8020/tasks | jq '.[0]' | jq -r '.id')

  vecho "Query one task by found ID ${ID}"
  curl -s ${IP}:8020/tasks/${ID} | jq
}

# update
update() {
  vecho "Update task ${ID} from file..."
  curl -s -X PUT -H "Content-Type:application/json" -d @clairin.json ${IP}:8020/tasks/${ID} | jq

  vecho "Query one task by found ID ${ID} after update"
  curl -s ${IP}:8020/tasks/${ID} | jq
}

# delete
delete() {
  vecho "Deleting task by ID ${ID}"
  curl -X DELETE -H "Content-Type:application/json" ${IP}:8020/tasks/${ID}
}

###########################
# main processing          #
############################

# args parsing
param $*

vecho "executing actions..."

# create
if [ "$CREATE" = "true" ] ; then
  create
fi;

# create
if [ "$QUERY" = "true" ] ; then
  query
fi;

# update
if [ "$UPDATE" = "true" ] ; then
  query
  update
fi;

# delete
if [ "$DELETE" = "true" ] ; then
  query
  delete
fi;

vecho "...done"
exit 0
