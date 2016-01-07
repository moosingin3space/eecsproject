#!/bin/sh
if [ "$TRAVIS_TAG" = "" ]; then
    RELEASE_NAME=snapshot
else
    RELEASE_NAME=$TRAVIS_TAG
fi

ghr --username moosingin3space --token $GITHUB_TOKEN --replace $RELEASE_NAME dist/
