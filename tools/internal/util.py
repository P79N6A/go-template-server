#!/usr/bin/python
# coding=utf-8
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
import platform

# Constants define

# Util functions
def GetSystemType():
    return platform.system()

def Logger(level, content):
    if level == "ERROR":
        print('\033[1;31;40m['+level+'] '+content+'\033[0m')
    elif level == "DEBUG":
        print('\033[1;32;40m['+level+'] '+content+'\033[0m')
    elif level == "INFO":
        print('\033[1;32;40m['+level+'] '+content+'\033[0m')

def Shell(cmd):
    try:
        Logger("DEBUG", cmd)
        retcode = call(cmd, shell=True)
        return retcode
    except OSError as e:
        Logger("ERROR", "Execution cmd failed: "+str(e))
        sys.exit(-1)

def ReplaceFile(filename, old_string, new_string):
    # Safely read the input filename using 'with'
    with io.open(filename, 'r', encoding="utf-8") as f:
        s = f.read()
        if old_string not in s:
            return
    # Safely write the changed content, if found in the file
    with io.open(filename, 'w', encoding="utf-8") as f:
        #Logger('DEBUG', 'Changing "{old_string}" to "{new_string}" in {filename}'.format(**locals()))
        s = s.replace(old_string, new_string)
        f.write(s)

def RenameFile(dirname, filename, old_string, new_string):
    if filename.find(old_string) != -1:
        new_filename  = filename.replace(old_string, new_string)
        Shell("mv "+dirname+"/"+filename+" "+dirname+"/"+new_filename)

def ToLowercase(camel_format):
    underline_format = ''
    upper_flag = True
    if isinstance(camel_format, str):
        for char in camel_format:
            if char.isupper():
                if not upper_flag:
                    underline_format += '_'
                underline_format += char.lower()
                upper_flag = True
            elif char.islower():
                upper_flag = False
                underline_format += char
            else:
                underline_format += char
    return underline_format

def ToUppercase(underline_format):
    camel_format = ''
    if isinstance(underline_format, str):
        char_list = underline_format.split('_')
        if len(char_list) == 1: # case: abcdef or AbcDef, if AbcDef, after capitalize becomes Abcdef
            return char_list[0][0].upper() + char_list[0][1:]
        for char in char_list:
            camel_format += char.capitalize()
    return camel_format


# Project function
kTemplateConfig = {
    "taf": {
        "template": "template",
        "servant_proto": "template/proto",
        "project_proto": "template/proto/project",
        "proto_name": "jce"
    },
    "tars": {
    }
}


def GetProjectDir(path):
    '''
    自下而上的反向遍历当前目录，直到找到project的目录结构为止，最多遍历5级目录
    '''
    max_travel_dir_num = 5
    project_dir_list = set(["proto", "config"])
    project_file_list = set(["README.md", "build.sh"])

    cur_travel_dir_num = 0
    for dir_name, sub_dir_list, file_list in os.walk(path, topdown=False):
        if max_travel_dir_num > cur_travel_dir_num:
            break
        if set(sub_dir_list) < project_dir_list:
            cur_travel_dir_num += 1
            continue
        if set(file_list) < project_file_list:
            cur_travel_dir_num += 1
            continue
        if not os.path.isfile(dir_name+"/proto/project.conf"):
            cur_travel_dir_num += 1
            continue
    return dir_name
    

def GetGitPath(group_name, project_name, use_git_key):
    if use_git_key:
        return "git@git.code.oa.com:"+group_name+"/"+project_name+".git"
    else:
        return "http://git.code.oa.com/"+group_name+"/"+project_name+".git"


def GetProjectConfig(project_name, project_type):
    config = '\
{ \n\
    "project_name": "%s", \n\
    "project_type": "%s", \n\
    "display_name": "%s", \n\
    "app_name": "%s" \n\
}' % (project_name, project_type, project_name, ToUppercase(project_name))
    return config


def LoadProjectConfig(file_path):
    with open(file_path, 'r') as content_file:
        return json.loads(content_file.read(), encoding="UTF-8")


def GetDefaultServantName(server_name):
    return server_name.replace("_server", "_servant")


def GetTemplateProjectPath(framework_type, tools_dir):
    return "%s/%s/%s_project" % (tools_dir, kTemplateConfig[framework_type]["template"], framework_type)


def GetTemplateProjectProtoPath(framework_type, tools_dir):
    return "%s/%s" % (tools_dir, kTemplateConfig[framework_type]["project_proto"])


def GetTemplateServantProtoPath(framework_type, tools_dir):
    return "%s/%s" % (tools_dir, kTemplateConfig[framework_type]["servant_proto"])


def GetTemplateServerPath(server_lang, server_type, framework_type, tools_dir):
    return "%s/%s/%s_%s_server/%s" % (tools_dir, kTemplateConfig[framework_type]["template"], framework_type, server_lang, server_type)

# Main functions
if __name__ == "__main__":
    print("util functions")
