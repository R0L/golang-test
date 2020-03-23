#!/usr/bin/env bash

load_config() {
  configFile=".config"
  if [ ! -f $configFile ]; then
    return 0
  fi
  printf "[1/6] load_config \`%s\`: \n" "$configFile"
  # config
  source "$configFile"
}

# shellcheck disable=SC2120
change_dir() {
  printf "[2/6] change dir \`%s\`: \n" "$1"
  if [ ! -d "$1" ]; then
    /bin/mkdir -pv "$1"
  fi
  cd "$1" || exit
}

exist_command() {
  if [ "$1" = '' ]; then
    return 0
  fi

  printf "[3/6] exists command \`%s\`: " "$1"
  if [ "$(command -v "$1")" != '' ]; then
    printf "\E[32m[EXIST]\E[0m \n"
    return 0
  else
    printf "\E[31m[MISS]\E[0m \n"
    printf "  - Please install it, via %s \n" "$2"
    break=1
    return 1
  fi
}

check_command() {
  break=0
  exist_command $1 "\`go get github.com/Shelnutt2/db2struct/cmd/db2struct\`"
  if [ $break = 1 ]; then
    exit
  fi
}

upper() {
  str=$1
  firstLetter=$(echo "${str:0:1}" | awk '{print toupper($0)}')
  otherLetter=${str:1}
  echo "$firstLetter$otherLetter"
}

lower() {
  str1=$1
  firstLetter1=$(echo "${str1:0:1}" | awk '{print tolower($0)}')
  otherLetter1=${str1:1}
  echo "$firstLetter1$otherLetter1"
}

toCamBak() {
  cambak=""
  list=${1//_/ }
  for item in ${list}
  do
       cambak=${cambak}$(upper "$item")
  done
  echo "$cambak"
}

gen() {
  printf "[4/6] gen command:\n"
  for table in "${TABLES[@]}"; do
    printf "  delte exists table file %s.go\n" "${table}"
    rm -rf "${SAVE_PATH}/${table}".go
    STRUCT=$(toCamBak "${table}")
    db2struct -H "${HOST}" -u "${USER}" -p "${PASSWORD}" \
      --mysql_port="${PORT}" -d "${DATABASE}" \
      --package "${GO_PACKAGE}" \
      --table "${table}" --struct "${STRUCT}" \
      --target="$(lower "${STRUCT}")".go \
      --gorm --json -v
    printf "  finish table file %s.go\n" "${table}"
  done
}

handle() {
  # 加载配置
  load_config

  # create path dir
  change_dir "$SAVE_PATH"

  # check command
  check_command "db2struct"

  # exce command
  gen
}

handle
printf "\E[32m[SUCCESS]\E[0m done!\n"
