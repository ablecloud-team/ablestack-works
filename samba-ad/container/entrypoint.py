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
    if os.path.isfile("/samba4/installed"):
        runsh = sh.Command("/usr/local/bin/samba_run.sh")
        runsh(_bg=True)
    else:
        parser.print_help()

if parsed[0].action=="config":
    config_parsed = config_parser.parse_args(parsed[1])
    print("========config_parsed=========")
    arg = config_parsed
    pprint.pprint(arg)
    configsh = sh.Command("/usr/local/bin/samba_config.sh")
    output = configsh(arg.nameserver, arg.domain, arg.password, arg.hostname)
    print(output)
runsamba()