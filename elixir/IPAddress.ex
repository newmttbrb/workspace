defmodule IPAddress do

  def to_tuple(str) when is_binary(str), do: str |> String.to_char_list! |> to_tuple
  def to_tuple(str) when is_list(str),   do: str |> :inet.parse_address  |> _to_tuple_parse_helper
  def to_tuple(str),                     do: str

  def to_string(tuple) when is_binary(tuple), do: tuple
  def to_string(tuple) when is_list(tuple),   do: tuple               |> iolist_to_binary
  def to_string(tuple),                       do: tuple |> :inet.ntoa |> iolist_to_binary

  def valid?(str), do: str |> to_tuple |> _not_nil

# private stuff
  defp _to_tuple_parse_helper({ :ok   , ip     }), do: ip
  defp _to_tuple_parse_helper({ :error, :einval}), do: nil

  defp _not_nil(x), do: not(nil?(x))
end
