#! /bin/bash

curl -X POST https://api.telegram.org/bot576734993:AAFhJQALvnwzXR8ZWlLGkRaZePlUcGGwMmA/sendMessage -d chat_id=324044527 \
-d "text=${TRAVIS_BRANCH} (${TRAVIS_COMMIT_MESSAGE}) tests result: ${TRAVIS_TEST_RESULT}"