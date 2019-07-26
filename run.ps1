# `dev` or `prod`
$mode=$args[0]

# `up` to start stack
# `build` to rebuild images and update dependencies
$action=$args[1]

$command = "docker-compose -f docker/docker-compose.yml -f docker/docker-compose." + $mode + ".yml " + $action

Invoke-Expression $command