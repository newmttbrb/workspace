defmodule TCP.Server do
	def start(port) do
		{:ok, sock} = :gen_tcp.listen(port,[{:active,true}, :binary])
		IO.puts "socket open - ready to accept connection"
		{:ok, accept} = :gen_tcp.accept(sock)
		IO.puts "socket connected - waiting for data"
		serverReceiver(accept)
		IO.puts "socket closing - shutting down server"
		:gen_tcp.close(accept)
		:gen_tcp.close(sock)
		IO.puts "socket closed - server exit"
	end

	defp serverReceiver(socket) do
		receive do
			{:tcp,_,"close"} -> 
				IO.puts "socket closing due to client close request"
				
			{:tcp,_,msg} -> 
				IO.puts msg
			    :gen_tcp.send(socket,"{:ok,\"#{msg}\"}")
			    serverReceiver(socket)
		end
	end
end
