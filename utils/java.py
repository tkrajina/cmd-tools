import os
import sys

def get_jar_contents( jar_file ):
	result = []

	if jar_file[ -4 : ].lower() != '.jar':
		return []

	try:
		command = 'jar -tf "%s"' % jar_file
		jar_output = os.popen( command )
	except Exception, e:
		print 'Error loading %s (%s)' % ( jar_file, e )
		return []

	for line in jar_output:
		line = line.strip()
		if line[ -6 : ] == '.class':
			line = line[ 0 : -6 ]
			line = line.replace( '/', '.' )
			result.append( line )

	return result

def find_in_jar( jar_file, class_string ):
	print 'Searching %s' % jar_file
	entries = get_jar_contents( jar_file )
	for entry in entries:
		if entry.lower().find( class_string.lower() ) >= 0:
			print "Found %s in %s" % ( entry, jar_file )
			answer = raw_input( 'Continue ( [y]es / [N]o / next [j]ar ) ?' )
			answer = answer.lower()
			if answer == 'n':
				sys.exit( 0 )
			elif answer == 'j':
				return
			else:
				pass

