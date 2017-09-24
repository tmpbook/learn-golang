#coding:utf8
import sys
print sys.getdefaultencoding()
a = "是"
print a.decode('ascii').encode()
print a.__str__()
