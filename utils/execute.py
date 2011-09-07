import os as _os

def execute( command, output = True, prefix = '' ):
	output = ''
	p = _os.popen( command )
	for line in p:
		output_line = prefix + ( '%s' % line ).strip() + '\n'
		if output:
			if output_line:
				print output_line.rstrip()
		else:
			output += output_line

	return ( not p.close(), output )
