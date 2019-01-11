#!/usr/bin/python
# coding=utf-8
###############################################################
# Author: juzhongguoji <juzhongguoji@hotmail.com>
# Date:   2018/12/16
###############################################################

import os
import sys
import getopt
from util import *

# Global defines
gSuccessCaseNum = 0
gFailedCaseNum = 0

# Main functions
def ToLowercaseTest():
    test_name = "ToLowercaseTest"
    Logger("INFO", "%s begin" % test_name) 
    def NormalCase():
        global gSuccessCaseNum, gFailedCaseNum
        case_result = True
        case_name = "NormalCase"
        #1 normal case
        input = "ExampleTestServer"
        output = "example_test_server"
        result = ToLowercase(input)
        if result != output:
            case_result = False
            Logger("ERROR", "%s.%s failed! Input[%s]-Output[%s]-Result[%s]" 
                % (test_name, case_name, input, output, result))
        #2 with number
        input = "ExampleTest222Server"
        output = "example_test222_server"
        result = ToLowercase(input)
        if result != output:
            case_result = False
            Logger("ERROR", "%s.%s failed! Input[%s]-Output[%s]-Result[%s]" 
                % (test_name, case_name, input, output, result))
        #3 with number
        input = "ExampleTestServer222"
        output = "example_test_server222"
        result = ToLowercase(input)
        if result != output:
            case_result = False
            Logger("ERROR", "%s.%s failed! Input[%s]-Output[%s]-Result[%s]" 
                % (test_name, case_name, input, output, result))
        #4 with number
        input = "Ex2am4444pl5eTestServer"
        output = "ex2am4444pl5e_test_server"
        result = ToLowercase(input)
        if result != output:
            case_result = False
            Logger("ERROR", "%s.%s failed! Input[%s]-Output[%s]-Result[%s]" 
                % (test_name, case_name, input, output, result))
        # Stat case result
        if case_result:
            gSuccessCaseNum += 1
        else:
            gFailedCaseNum += 1
    # Run case
    NormalCase()
    Logger("INFO", "%s finish" % test_name) 


def ToUppercaseTest():
    test_name = "ToUppercaseTest"
    Logger("INFO", "%s begin" % test_name)

    def NormalCase():
        global gSuccessCaseNum, gFailedCaseNum
        case_result = True
        case_name = "NormalCase"
        #1 normal case
        input = "ExampleTestServer"
        output = "ExampleTestServer"
        result = ToUppercase(input)
        if result != output:
            case_result = False
            Logger("ERROR", "%s.%s failed! Input[%s]-Output[%s]-Result[%s]"
                   % (test_name, case_name, input, output, result))
        #2 with number
        input = "ExampleTest222Server"
        output = "ExampleTest222Server"
        result = ToUppercase(input)
        if result != output:
            case_result = False
            Logger("ERROR", "%s.%s failed! Input[%s]-Output[%s]-Result[%s]"
                   % (test_name, case_name, input, output, result))
        #3 with number
        input = "example_test_server222"
        output = "ExampleTestServer222"
        result = ToUppercase(input)
        if result != output:
            case_result = False
            Logger("ERROR", "%s.%s failed! Input[%s]-Output[%s]-Result[%s]"
                   % (test_name, case_name, input, output, result))
        #4 with number
        input = "ex2am4444pl5e_test_server"
        output = "Ex2am4444pl5eTestServer"
        result = ToUppercase(input)
        if result != output:
            case_result = False
            Logger("ERROR", "%s.%s failed! Input[%s]-Output[%s]-Result[%s]"
                   % (test_name, case_name, input, output, result))
        # Stat case result
        if case_result:
            gSuccessCaseNum += 1
        else:
            gFailedCaseNum += 1
    # Run case
    NormalCase()
    Logger("INFO", "%s finish" % test_name)


if __name__ == "__main__":
    ToLowercaseTest()
    ToUppercaseTest()
    Logger("INFO", "Totally test %d cases, success %d, failed %d" % (gSuccessCaseNum+gFailedCaseNum, gSuccessCaseNum, gFailedCaseNum))
