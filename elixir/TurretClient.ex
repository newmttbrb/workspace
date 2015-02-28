################################################
#Turret.Client.start_link(:turret,{10,204,45,168},8781)
#Turret.Client.send(:turret,"<?xml version=\"1.0\"?><commands><enableAutomation autoFormat=\"true\" cpus=\"0 1\"/><pressButton cpuid=\"0\" button=\"12\" duration=\"200\"/><disableAutomation/><returnQueuedMessages/></commands>")
#Turret.Client.send(:turret,"<?xml version=\"1.0\"?><commands><enableAutomation autoFormat=\"true\" cpus=\"0 1\"/><pressButton cpuid=\"0\" button=\"12\" duration=\"200\"/><pause duration=\"800\"/><pressButton cpuid=\"0\" button=\"21\" duration=\"200\"/><pause duration=\"800\"/><pressButton cpuid=\"0\" button=\"22\" duration=\"200\"/><pause duration=\"800\"/><pressButton cpuid=\"1\" button=\"37\" duration=\"200\"/><pause duration=\"800\"/><pressButton cpuid=\"1\" button=\"0\" duration=\"200\"/><pause duration=\"800\"/><pressButton cpuid=\"0\" button=\"28\" duration=\"200\"/><pause duration=\"300\"/><pressButton cpuid=\"0\" button=\"25\" duration=\"200\"/><pause duration=\"300\"/><pressButton cpuid=\"0\" button=\"25\" duration=\"200\"/><pause duration=\"300\"/><pressButton cpuid=\"0\" button=\"27\" duration=\"200\"/><pause duration=\"300\"/><pressButton cpuid=\"0\" button=\"36\" duration=\"200\"/><pause duration=\"300\"/><pause duration=\"2000\"/><pressButton cpuid=\"0\" button=\"11\" duration=\"200\"/><pause duration=\"800\"/><pressButton cpuid=\"1\" button=\"1\" duration=\"200\"/><pause duration=\"800\"/><pause duration=\"10000\"/><pressButton cpuid=\"0\" button=\"21\" duration=\"200\"/><pause duration=\"800\"/><disableAutomation/><returnQueuedMessages/></commands>")
################################################
defmodule Turret.Client do
	use GenServer.Behaviour

	#####
	# External API

	def start_link(name,host,port) do
		:gen_server.start_link({ :local, name }, __MODULE__, {_to_tuple(host),port}, [])
	end

	def send(name, message) do
   		:gen_server.call name, {:send, message}, :infinity
	end

	def send_automation(name, message) do
   		:gen_server.call name, {:send_automation, message}, :infinity
	end

	def receive(name) do
   		:gen_server.call name, {:receive, nil}, :infinity
	end

	#def increment_number(name,delta) do
   	#	:gen_server.cast :sequence, {:increment_number, delta}
  	#end

	#####
	# GenServer implementation
	def init({host,port}) do
		{ :ok, {host,port} }
	end

	def handle_call({:send, message}, _from, {host,port}) do
		tcp_send(tcp_connect(host,port), message)
		result = _tcp_receive()
		{ :reply, result, {host,port} }
	end

	def handle_call({:receive, nil }, _from, {host,port}) do
		result = _tcp_receive()
		{ :reply, result, {host,port} }
	end

	def handle_call({:send_automation, message}, _from, {host,port}) do
		tcp_send_automation(tcp_connect(host,port), message)
		result = _tcp_receive()
		{ :reply, result, {host,port} }
	end

	#####
	# TCP stuff
	defp tcp_connect(host,port) do 
		{ :ok, socket } = :gen_tcp.connect(host,port,[:binary, {:active, true}]) 
		socket
	end

	defp tcp_send_automation(socket,message), do: :ok = :gen_tcp.send(socket,"#{String.length(message)};#{message}")	
	defp tcp_send(socket,message), do: :ok = :gen_tcp.send(socket,message)	
	
	defp _tcp_receive() do
		receive do
			{:tcp_closed, _socket} -> []
			{:tcp_error, _socket, reason} -> [reason]
			{:tcp, _socket, data} -> [data | _tcp_receive()]
		end
	end

	defp _to_tuple str do
		addr = IPAddress.to_tuple(str) 
	 	case addr do
	 		nil -> {127,0,0,1}
	 		{_a1,_a2,_a3,_a4} -> addr
	 		true -> addr
	 	end
	 end
end
