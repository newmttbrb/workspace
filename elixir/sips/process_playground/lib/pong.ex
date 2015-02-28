defmodule Pong do
  def start() do
    await(0)
  end

  def await(count) do
    receive do
      {:ping, pid} -> 
        IO.puts "pong received another ping (#{count})"
        send(pid, {:pong, self})
    end
    await(count+1)
  end
end

