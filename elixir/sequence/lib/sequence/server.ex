# iex -S mix
# Sequence.Server.start_link(123)
# Sequence.Server.next_number
# Sequence.Server.next_number
# Sequence.Server.increment_number(100)
# Sequence.Server.next_number
defmodule Sequence.Server do
	use GenServer.Behaviour

	### external API ###
	def start_link(stash_pid) do
		:gen_server.start_link({:local, :sequence}, __MODULE__, stash_pid, [])
	end

	def next_number do
		:gen_server.call :sequence, :next_number
	end

	def increment_number(delta) do
		:gen_server.cast :sequence, {:increment_number, delta}
	end

	### genserver ###
	def init(stash_pid) do
		current_number = Sequence.Stash.get_value stash_pid
		{ :ok, {current_number, stash_pid } }
	end

	def handle_call(:next_number, _from, { current_number, stash_pid }) do
		{:reply, current_number, { current_number+1, stash_pid } }
	end

	#def handle_call({:set_number, new_number}, _from, _current_number) do
	#	{ :reply, new_number, new_number }
	#end
	
	# def handle_call({:factors, number}, _, _) do
	# 	{:reply, {:factors_of, number, factors(number)},[]}
	# end

	def handle_cast(:increment_number, delta, { current_number, stash_pid }) do
		{:noreply, { current_number + delta, stash_pid } }
	end

	def terminate(_reason, {current_number, stash_pid}) do
		Sequence.Stash.save_value stash_pid, current_number
	end

	def format_status(_reason, [ _pdict, state]) do
		[data: [{'State', "My current state is '#{inspect state}', and I'm happy"}]]
	end
	
end