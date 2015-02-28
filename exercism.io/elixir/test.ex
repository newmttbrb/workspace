defmodule Test do
  def concat(ll) do
    reduce(ll, [], fn element, acc ->
      reduce(element, acc, &([&1 | &2]))
    end) |> reverse
  end

 def reduce(l, acc, f) do
    case l do
      []      -> acc
      [h | t] -> reduce(t, f.(h, acc), f)
    end
  end

  def reverse(l) do
    reduce(l, [], &([&1 | &2]))
  end
end
