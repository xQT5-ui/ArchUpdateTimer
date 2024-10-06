# OVERWRITE

This is a small project for automatization execute some shell command through a emulatior of terminal.
I created it for autoupdating system for *Arch Linux* with using `yay` and `kgx`. But I suggest that you can use it for other actions because of the programm is using config file for managing programm.

## SETTINGS AND INSTALLS

Before to build a programm you need:

0. Install [go](https://go.dev/doc/install)
1. Download the code to your folder
2. Create **config.yaml** in *your_download_folder/app/config/* directive like this:

```yaml
command:
  exec_value: "yay -Syu --noconfirm && echo \"\nJobs for flatpak...\n\" && sudo flatpak repair && flatpak update -y && sudo flatpak remove --unused -y && echo \"\nJobs for flatpak...\n\" && sudo fwupdmgr update && sleep 3 && pkill kgx"
  # run command with auto-enter password
  cmd_shell_flag: true
  cmd_shell_exec: "sudo -S sh -c 'echo \"\nStart updating system...\"'"
  # kgx - emulator of terminal; -e - argument for execute command
  emulation_terminal: ["kgx", "-e"]
# your sudo password in base64
password: "MTIzNDVxd2VydHk=" #12345qwerty for example
# your base64 password hash in sha256
hash: "b554b3ad18c9605a700f350f2e3708af9ca03857afbefde4a429ec0d6b1a9965"
# days:
# - 0 = sunday
# - 3 = wednesday
# - 5 = friday
days:
  range: [0, 3, 5]
  range_flg: true
```

3. Run **install.sh** in current programm folder:

```bash
./install.sh
```

### For uninstalling programm run uninstall.sh in current programm folder

```bash
./uninstall.sh
```
