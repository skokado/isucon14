# ssh target host is depended on ~/.ssh/config
WORKSPACE="/home/skokado/workspace/isucon14"

copy_to_server_and_restart_service() {
  local host=$1

  echo "Deploy to $host"

  rsync -ahv --delete \
      --exclude .venv \
      --exclude .git \
      --exclude __pycache__ \
      --exclude *.pyc \
    $WORKSPACE/webapp/python \
    ${host}:/home/isucon/webapp/
  
  rsync -az \
    $WORKSPACE/webapp/sql \
    ${host}:/home/isucon/webapp/

  rsync -az \
    $WORKSPACE/env.sh \
    ${host}:/home/isucon/

  rsync -az \
    $WORKSPACE/initialize.sh \
    ${host}:/home/isucon/

  ssh $host "cd /home/isucon/webapp/python && /home/isucon/.local/bin/uv sync"
  ssh $host "sudo systemctl restart isuride-python.service"
  ssh $host "sudo systemctl restart isuride-matcher.service"

  echo "Deploy to $host OK."

}

copy_to_server_and_restart_service isucon1
