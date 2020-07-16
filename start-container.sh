#!/bin/sh

if [[ $CONFIG_VARS ]]; then
  JSON=`json_env ${CONFIG_FILE_PATH}/config.js`

  echo " ==> Writing ${CONFIG_FILE_PATH}/config.js with ${JSON}"

  echo "window.__env = ${JSON}" > ${CONFIG_FILE_PATH}/config.js
fi

exec "$@"
