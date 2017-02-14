setup_sshd() {
    #ssh-keygen -A
    #sudo /usr/sbin/sshd
    #service sshd start
    mkdir -p ${HOME}/.ssh
    ssh-keygen -t rsa -q -f ${HOME}/.ssh/orbit.key -N ""
    cp ${HOME}/.ssh/orbit.key.pub ${HOME}/.ssh/authorized_keys
    ssh-keyscan -t ecdsa-sha2-nistp256 localhost > ${HOME}/.ssh/known_hosts
}

init_orbit() {
    export ORBIT_KEY=/.ssh/orbit.key
    export ORBIT_HOME=/code/bintest/testFolder
    #export PATH=`pwd`/bintest/tools:$PATH
    #chmod -R u+x `pwd`/bintest/tools
    eval `ssh-agent -s`
    ssh-add ${HOME}${ORBIT_KEY}
}

setup_sshd
init_orbit
