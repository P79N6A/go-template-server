#!/usr/bin/env python
###############################################################
# Author: juzhongguoji <juzhongguoji@hotmail.com>
# Date:   2018/12/16
###############################################################

import getopt
import io
import os
import sys
from subprocess import call
import json
import datetime
from internal.util import *

# Global define
gToolsDir = os.path.dirname(os.path.realpath(__file__))
gCurDir = os.getcwd()
gHomeDir = os.path.expanduser("~")
gRtxName = ""
gCurDate = datetime.datetime.today().strftime('%Y/%m/%d')
gPrjDir = gHomeDir + "/go/src/online_consultant"

# Main functions
def Process():
    proto_dir = gPrjDir + "/proto"
    Logger('INFO', '############# Gen proto file #############')
    for dir_name, sub_dir_list, file_list in os.walk(proto_dir):
        os.chdir(dir_name)
        for file_name in file_list:
            name, ext = os.path.splitext(file_name)
            if ext == ".proto":
                Shell('protoc -I. -I.. -I' + gHomeDir + '/go/src --go_out=plugins=grpc:. '+file_name)
    Logger('INFO', 'Finish gen proto')


def Usage():
    print '''Usage: 
  %s --all

Parameters:
  all - Process all proto file

Example: %s --all
  Gen all proto file to pb
''' % (os.path.basename(__file__), os.path.basename(__file__))


if __name__ == "__main__":
    if len(sys.argv) <= 1:
        Usage()
        sys.exit(0)
    is_all = False
    lang = "go"
    try:
        options, args = getopt.getopt(sys.argv[1:], "ha", ['help', "all"])
        for name, value in options:
            if name in ('-h', '--help'):
                Usage()
            elif name in ('-a', '--all'):
                is_all = True
    except getopt.GetoptError, ex:
        Logger("ERROR", "Command line input error: "+str(ex))
        Usage()
        sys.exit(0)
    # Do main process
    Process()
