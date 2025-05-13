need_start_server_shell=(
  # rpc
  user-rpc-test.sh

  # api
)

# shellcheck disable=SC2048
for i in ${need_start_server_shell[*]} ; do
    chmod +x $i
    ./$i
done


docker ps

docker exec -it etcd etcdctl get --prefix ""