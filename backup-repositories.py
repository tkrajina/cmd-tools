#!/usr/bin/python
# -*- coding: utf-8 -*-

import sys           as mod_sys
import os            as mod_os
import datetime      as mod_datetime
import shutil        as mod_shutil
import traceback     as mod_traceback

home = mod_os.environ[ 'HOME' ]
workspace_path = '{0}/workspace'.format( home )
backup_path = '{0}/backup'.format( home )
log_file = '{0}/backup.log'.format( home )

timestamp = mod_datetime.datetime.now().strftime( '%Y-%m-%d-%H-%M-%S' )

backup_file_name = 'git-backup-{0}.tar'.format( timestamp )

def execute( cmd ):
	p = mod_os.popen( cmd )
	for line in p:
		try:
			print line.strip()
		except Exception, e:
			# Pukne u cronu, nije problem...
			pass
	if p.close() != None:
		raise Exception( 'Error executing {0}'.format( cmd ) )

def backup():
	mod_os.chdir( workspace_path )
	execute( 'tar -cvf "{0}" */.git'.format( backup_file_name ) )
	mod_shutil.move( backup_file_name, '{0}/{1}'.format( backup_path, backup_file_name ) )
	print 'DONE'

try:
	backup()
except Exception, e:
	f = open( log_file, 'w' )
	f.write( mod_traceback.format_exc() )
	f.close()
	raise e
