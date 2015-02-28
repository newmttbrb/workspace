defmodule TCP.Client do
	def start!(host,port) do
		{:ok, socket} = :gen_tcp.connect(host,port,[{:active,true}, :binary])
		socket
	end

	def send!(socket,message), do: :gen_tcp.send(socket,message)
	
	def close!(socket), do: send(socket,"close")

	def recv!() do
		receive do
			{:tcp,_,msg} -> msg
		end
	end
end
