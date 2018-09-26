#! /bin/bash

curl -X POST https://api.telegram.org/bot576734993:AAFhJQALvnwzXR8ZWlLGkRaZePlUcGGwMmA/sendMessage -d chat_id=-268977972 \
-d "text=${TRAVIS_BRANCH} (${TRAVIS_COMMIT_MESSAGE}) tests return code: ${TRAVIS_TEST_RESULT}"