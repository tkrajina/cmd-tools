import sys as mod_sys

class OptParser:
	""" Similar to getopt """
	
	args = None
	comments = []

	def __init__( self, args = None ):
		if args:
			self.args = args
		else:
			self.args = mod_sys.argv[ 1 : ]

	def has_param( self, short_name = None, long_name = None, comment = None ):
		if comment:
			if short_name and long_name:
				arg = '-%s or --%s' % ( short_name, long_name )
			elif short_name:
				arg = '-%s' % short_name
			else:
				arg = '--%s' % long_name
			self.comments.append( '%s -> %s' % ( arg, comment ) )
		for arg in self.args:
			if short_name and arg == '-%s' % short_name:
				self.args.remove( '-%s' % short_name )
				return True
			if long_name and arg == '--%s' % long_name:
				self.args.remove( '--%s' % long_name )
				return True
		return None

	def get_param( self, short_name = None, long_name = None, default = None, comment = None ):
		if comment:
			if short_name and long_name:
				arg = '-%s <value> or --%s=<value>' % ( short_name, long_name )
			elif short_name:
				arg = '-%s <value>' % short_name
			else:
				arg = '--%s=<value>' % long_name
			self.comments.append( '%s -> %s' % ( arg, comment ) )
		for i in range( len( self.args ) ):
			arg = self.args[ i ]
			if short_name and arg == '-%s' % short_name:
				if i < len( self.args ) - 1:
					result = self.args[ i + 1 ]
					self.args.remove( '-%s' % short_name )
					self.args.remove( result )
					if not result:
						return default
					return result
				else:
					result = None
					if i < len( self.args ) - 1:
						self.args.remove( i )
					if not result:
						return default
					return result
			if long_name and arg.startswith( '--%s=' % long_name ): 
				result = arg[ len( '--%s=' % long_name ) : ]
				self.args.remove( arg )
				if not result:
					return default
				return result

		if default:
			return default
		
		return None

	def params_left( self ):
		result = []

		for arg in self.args:
			if arg.startswith( '-' ):
				result.append( args )

		return result

	def args_left( self ):
		result = []

		for arg in self.args:
			if not arg.startswith( '-' ):
				result.append( arg )

		return result

	def get_comments( self ):
		return self.comments

if __name__ == '__main__':
	args = [ '-a', '1', '-b', '--cccc=sdhajdhj', '--ddd', 'aaaaaads_sdajkdljk', '--xxxxxxxx' ]

	parser = OptParser( args )

	print parser.has_param( short_name = 'b' )
	print parser.get_param( short_name = 'a' )
	print parser.get_param( long_name = 'cccc' )
	print parser.get_param( long_name = 'xxx', default = 'vrijednost od x' )
	print parser.has_param( long_name = 'ddd' )

	print parser.args_left( remove_params = True )
