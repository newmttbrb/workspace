credentials ssh-turret-75   :key-based       ipcdiag         /opt/ipc/security/keymgmt/sshkeys/diag_rsa_10.204.68.75

device      left-turret     :10.204.45.159   ssh-turret-75
device      right-turret    :10.204.44.28    ssh-turret-75

turret      left-turret     :enableAutomation
turret      left-turret     :forcePaint
turret      left-turret     :pause duration="500"
turret      left-turret     :disableAutomation
turret      left-turret     :returnQueuedMessages

turret      right-turret    :enableAutomation
turret      right-turret    :forcePaint
turret      right-turret    :pause duration="500"
turret      right-turret    :disableAutomation
turret      right-turret    :returnQueuedMessages

