#!/usr/bin/python3
import sys
import argparse
import pprint
import sh
import os.path

print(sys.argv)

parser = argparse.ArgumentParser(description='Ablecloud AD initiator')

parser.add_argument("action", choices=["config", "run"], default="run", nargs='?', type=str)

config_parser = argparse.ArgumentParser(description='Ablecloud AD initiator')
# config_parser.add_argument("ip", type=str)
config_parser.add_argument("nameserver", type=str)
config_parser.add_argument("domain", type=str)
config_parser.add_argument("password", type=str, default="Ablecloud1!")
config_parser.add_argument("hostname", type=str, default="ad")

parsed=parser.parse_known_args()
print("========parsed=========")
pprint.pprint(parsed[0])
pprint.pprint(parsed[1])

def runsamba():
    if os.path.isfile("/usr/local/samba/etc/installed"):
        runsh = sh.Command("/usr/local/bin/samba_run.sh")
        print("========samba started=========")
        for output in runsh(_iter=True, _err_to_out=True):
            print(output,end='')
    else:
        parser.print_help()

if parsed[0].action=="config":
  if not os.path.isfile("/usr/local/samba/etc/installed"):
    config_parsed = config_parser.parse_args(parsed[1])
    print("========config_parsed=========")
    arg = config_parsed
    pprint.pprint(arg)
    configsh = sh.Command("/usr/local/bin/samba_config.sh")
    for output in configsh(arg.nameserver, arg.domain, arg.password, arg.hostname, _iter=True, _err_to_out=True):
        print(output,end='')
runsamba()